package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unblvvv/h-www-server/internal/config"
	"github.com/unblvvv/h-www-server/internal/handler"
	applicationhandler "github.com/unblvvv/h-www-server/internal/handler/application"
	authhandler "github.com/unblvvv/h-www-server/internal/handler/auth"
	posthandler "github.com/unblvvv/h-www-server/internal/handler/post"
	"github.com/unblvvv/h-www-server/internal/handler/post/admin"
	"github.com/unblvvv/h-www-server/internal/middleware"
	"github.com/unblvvv/h-www-server/internal/repository"
	"github.com/unblvvv/h-www-server/internal/repository/application"
	"github.com/unblvvv/h-www-server/internal/repository/auth"
	"github.com/unblvvv/h-www-server/internal/repository/post"
	applicationservice "github.com/unblvvv/h-www-server/internal/service/application"
	authservice "github.com/unblvvv/h-www-server/internal/service/auth"
	postservice "github.com/unblvvv/h-www-server/internal/service/post"
	"github.com/unblvvv/h-www-server/internal/service/storage"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			fx.Annotate(
				auth.NewFx,
				fx.As(new(auth.Repository)),
			),
			fx.Annotate(
				post.NewFx,
				fx.As(new(post.Repository)),
			),
			fx.Annotate(
				application.NewFx,
				fx.As(new(application.Repository)),
			),
		),
		fx.Provide(
			config.Load,
			repository.NewDB,

			authservice.New,
			postservice.New,
			storage.New,
			applicationservice.New,

			NewHumaAPI,

			func() *gin.Engine {
				r := gin.Default()

				r.Use(cors.New(cors.Config{
					AllowOrigins:     []string{"*"},
					AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
					AllowHeaders:     []string{"*"},
					ExposeHeaders:    []string{"Content-Length"},
					AllowCredentials: true,
				}))

				return r
			},
		),

		authhandler.FxModule,
		posthandler.FxModule,
		admin.FxModule,
		applicationhandler.FxModule,

		fx.Invoke(
			startServer,
			handler.RegisterRoutes,
		),
	)
}

func startServer(lc fx.Lifecycle, r *gin.Engine, cfg *config.Config) {
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Fatalf("Server error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}

func NewHumaAPI(r *gin.Engine, cfg *config.Config) huma.API {
	humaConfig := huma.DefaultConfig("hwww", "1.0.0")

	humaConfig.Formats["multipart/form-data"] = huma.Format{
		Unmarshal: func(data []byte, v any) error {
			return nil
		},
	}

	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
		"admin_bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	api := humagin.New(r, humaConfig)

	api.UseMiddleware(middleware.AuthMiddleware(api, cfg))

	return api
}
