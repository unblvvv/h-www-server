package auth

import (
	"context"

	"github.com/unblvvv/h-www-server/internal/model"
	"github.com/unblvvv/h-www-server/internal/repository/post"
)

type AService struct {
	repo post.Repository
}

func New(repo post.Repository) *AService {
	return &AService{
		repo: repo,
	}
}

func (s *AService) CreateAPost(ctx context.Context, orgID string, name string, age string,
	sex model.ASex, description string, photoURL *string, status model.AStatus) (string, error) {
	newAnimal := &model.APost{
		OrganizationID: orgID,
		Name:           name,
		Age:            age,
		Sex:            sex,
		Description:    description,
		PhotoURL:       photoURL,
		Status:         status,
	}

	return s.repo.CreatePost(ctx, newAnimal)
}
