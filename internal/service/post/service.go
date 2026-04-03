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

func (s *Service) CreateAPost(ctx context.Context, orgID, name string, age string,
	sex model.ASex, description string, photoURLs []string, status model.AStatus) (string, error) {
	newAnimal := &model.APost{
		OrganizationID: orgID,
		Name:           name,
		Age:            age,
		Sex:            sex,
		Description:    description,
		PhotoURLs:      photoURLs,
		Status:         status,
	}

	return s.repo.CreatePost(ctx, newAnimal)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.DeletePost(ctx, id)
}

func (s *Service) Update(ctx context.Context, p model.APost) error {
	return s.repo.UpdatePost(ctx, p.Name, p.Age, string(p.Sex), p.Description, p.PhotoURLs, p.ID)
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
