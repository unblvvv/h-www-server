package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"github.com/unblvvv/h-www-server/internal/config"
	"github.com/unblvvv/h-www-server/internal/handler"
	authhandler "github.com/unblvvv/h-www-server/internal/handler/auth"
	"github.com/unblvvv/h-www-server/internal/repository"
	"github.com/unblvvv/h-www-server/internal/repository/auth"
	authservice "github.com/unblvvv/h-www-server/internal/service/auth"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			fx.Annotate(
				auth.NewFx,
				fx.As(new(auth.Repository)),
			),
		),
		fx.Provide(
			gin.New,

			config.Load,
			repository.NewDB,

			authservice.New,

			NewHumaAPI,
		),
		authhandler.FxModule,
		fx.Invoke(
			startServer,
			handler.RegisterRoutes,
		),
	)
}

func startServer(lc fx.Lifecycle, r *gin.Engine) {
	srv := &http.Server{
		Addr:    ":8080",
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
	humaConfig := huma.DefaultConfig("BiteWay API", "1.0.0")

	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	api := humagin.New(r, humaConfig)

	//api.UseMiddleware(middleware.AuthMiddleware(api, cfg))

	return api
}
