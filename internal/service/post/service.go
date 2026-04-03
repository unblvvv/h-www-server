package auth

import (
	"context"

	"github.com/unblvvv/h-www-server/internal/model"
	"github.com/unblvvv/h-www-server/internal/repository/post"
)

type Service struct {
	repo post.Repository
}

func New(repo post.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateAPost(ctx context.Context, orgID string, name string, age string,
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

func (s *Service) Update(ctx context.Context, post model.APost, userId string) error {
	return s.repo.UpdatePost(ctx, post.Name, post.Age, string(post.Sex), post.Description, post.PhotoURL, post.ID, userId)
}

func (s *Service) Delete(ctx context.Context, id, userID string) error {
	return s.repo.DeletePost(ctx, id, userID)
}

func (s *Service) GetAll(ctx context.Context, limit, offset int) ([]model.APost, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.GetPost(ctx, limit, offset)
}
