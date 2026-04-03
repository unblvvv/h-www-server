package middleware

import (
	"errors"
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
		isProtected := len(ctx.Operation().Security) > 0
		if !isProtected {
			next(ctx)
			return
		}

		authHeader := ctx.Header("Authorization")
		if authHeader == "" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid or expired token", err)
			return
		}

		if claims, ok := token.Claims.(*TokenClaims); ok {
			ctx = huma.WithValue(ctx, "userID", claims.UserID)
			next(ctx)
		} else {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token claims")
		}
	}
}
