package application

import (
	"context"

	"github.com/unblvvv/h-www-server/internal/model"
	"github.com/unblvvv/h-www-server/internal/repository/application"
)

type Service struct {
	repo application.Repository
}

func New(repo application.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateApplication(ctx context.Context, userID, animalID, name, email, phone, message string) error {
	return s.repo.Create(ctx, &model.Application{
		UserID:   userID,
		AnimalID: animalID,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Message:  message,
	})
}

func (s *Service) GetInbox(ctx context.Context, status *string, limit, offset int) ([]model.Application, int, error) {
	return s.repo.GetList(ctx, status, limit, offset)
}

func (s *Service) ResolveApplication(ctx context.Context, id string, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
