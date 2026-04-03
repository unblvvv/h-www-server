package auth

import (
	"context"

	"github.com/unblvvv/h-www-server/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}
