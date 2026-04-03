package post

import (
	"context"

	"github.com/unblvvv/h-www-server/internal/model"
)

type Repository interface {
	CreatePost(ctx context.Context, post *model.APost) (string, error)
}
