package middleware

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/unblvvv/h-www-server/internal/config"
)

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware(api huma.API, cfg *config.Config) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		op := ctx.Operation()

		if op == nil || len(op.Security) == 0 {
			next(ctx)
			return
		}

		requiresAdmin := false

		for _, sec := range op.Security {
			if _, ok := sec["admin_bearer"]; ok {
				requiresAdmin = true
				break
			}
		}

		authHeader := ctx.Header("Authorization")
		if authHeader == "" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Authorization header is missing")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid authorization format")
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		role, _ := claims["role"].(string)
		if requiresAdmin && role != "admin" {
			huma.WriteErr(api, ctx, http.StatusForbidden, "Only admins can access this resource")
			return
		}

		userID, _ := claims["user_id"].(string)
		ctx = huma.WithValue(ctx, "userID", userID)

		next(ctx)
	}
}
