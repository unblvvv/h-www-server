package auth

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/model"
	"github.com/unblvvv/h-www-server/internal/repository/auth"
	authservice "github.com/unblvvv/h-www-server/internal/service/auth"
)

type RegisterRequestDto struct {
	Body struct {
		Email    string `json:"email" format:"email" doc:"email"`
		Username string `json:"username" minLength:"3" maxLength:"50" doc:"username"`
		Password string `json:"password" minLength:"6" doc:"password"`
	}
}

type RegisterResponseOutput struct {
	Body struct {
		Token string `json:"token" doc:"jwt token"`
	}
}

type Register struct {
	service *authservice.AuthService
	repo    auth.Repository
}

func NewRegister(service *authservice.AuthService, repo auth.Repository) *Register {
	return &Register{service: service, repo: repo}
}

func (s *Register) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "auth-register",
		Method:      "POST",
		Path:        "/v1/auth/register",
		Tags:        []string{"Auth"},
		Description: "Register with email, username and password",
	}
}

func (s *Register) Handler(ctx context.Context, input *RegisterRequestDto) (*RegisterResponseOutput, error) {
	existingUser, _ := s.repo.GetUserByEmail(ctx, input.Body.Email)
	if existingUser != nil {
		return nil, huma.Error409Conflict("User with this email already exists")
	}

	user := model.User{
		Email:    input.Body.Email,
		Username: input.Body.Username,
		Password: input.Body.Password,
	}

	id, err := s.service.CreateUser(ctx, user)
	if err != nil {
		return nil, huma.Error500InternalServerError("Failed to create user", err)
	}

	return &RegisterResponseOutput{
		Body: struct {
			Token string `json:"token" doc:"jwt token"`
		}{
			Token: id,
		},
	}, nil
}

func (s *Register) Register(api huma.API) {
	huma.Register(api, s.GetMeta(), s.Handler)
}
