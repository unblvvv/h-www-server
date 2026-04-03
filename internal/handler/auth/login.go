package auth

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/repository/auth"
	authservice "github.com/unblvvv/h-www-server/internal/service/auth"
)

type LoginRequestDto struct {
	Body struct {
		Email    string `json:"email" format:"email" doc:"email"`
		Password string `json:"password" minLength:"6" doc:"password"`
	}
}

type LoginResponseOutput struct {
	Body struct {
		Token    string `json:"token" doc:"jwt token"`
		Username string `json:"username" doc:"username"`
		Email    string `json:"email" doc:"email"`
	}
}

type Login struct {
	service *authservice.AuthService
	repo    auth.Repository
}

func NewLogin(service *authservice.AuthService, repo auth.Repository) *Login {
	return &Login{service: service, repo: repo}
}

func (s *Login) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "auth-login",
		Method:      "POST",
		Path:        "/v1/auth/login",
		Tags:        []string{"Auth"},
		Description: "Login with email and password",
	}
}

func (s *Login) Handler(ctx context.Context, input *LoginRequestDto) (*LoginResponseOutput, error) {
	token, user, err := s.service.GenerateToken(ctx, input.Body.Email, input.Body.Password)
	if err != nil {
		return nil, huma.Error401Unauthorized("Invalid email or password", err)
	}

	return &LoginResponseOutput{
		Body: struct {
			Token    string `json:"token" doc:"jwt token"`
			Username string `json:"username" doc:"username"`
			Email    string `json:"email" doc:"email"`
		}{
			Token:    token,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *Login) Register(api huma.API) {
	huma.Register(api, s.GetMeta(), s.Handler)
}
