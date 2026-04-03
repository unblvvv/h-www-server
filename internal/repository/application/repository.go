package application

import (
	"context"

	"github.com/unblvvv/h-www-server/internal/model"
)

type Repository interface {
	Create(ctx context.Context, app *model.Application) error
	GetList(ctx context.Context, status *string, limit, offset int) ([]model.Application, int, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	Delete(ctx context.Context, id string) error
}
