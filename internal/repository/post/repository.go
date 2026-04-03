package post

import (
	"context"

	"github.com/unblvvv/h-www-server/internal/model"
)

type Repository interface {
	CreatePost(ctx context.Context, post *model.APost) (string, error)
	DeletePost(ctx context.Context, id string) error
	GetPost(ctx context.Context, limit, offset int) ([]model.APost, error)
	UpdatePost(ctx context.Context, name, age, sex, description string, photo_urls []string, post_id string) error
}
