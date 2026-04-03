package post

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/model"
	"github.com/unblvvv/h-www-server/internal/repository/post"
	postservice "github.com/unblvvv/h-www-server/internal/service/post"
)

type ListAnimalRequestDto struct {
	Limit  int `query:"limit" default:"20" minimum:"1" maximum:"100" doc:"items per page"`
	Offset int `query:"offset" default:"0" minimum:"0" doc:"items to skip"`
}

type ListPostOutput struct {
	Body struct {
		Items []model.APost `json:"items"`
	}
}

type ListPost struct {
	service *postservice.Service
	repo    post.Repository
}

func NewListPost(service *postservice.Service, repo post.Repository) *ListPost {
	return &ListPost{service: service, repo: repo}
}

func (s *ListPost) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "list-animal-posts",
		Method:      "GET",
		Path:        "/animal",
		Tags:        []string{"Posts"},
		Description: "Get a list of animal posts",
	}
}

func (s *ListPost) Handler(ctx context.Context, input *ListAnimalRequestDto) (*ListPostOutput, error) {
	items, err := s.service.GetAll(ctx, input.Limit, input.Offset)
	if err != nil {
		return nil, huma.Error500InternalServerError("InternalServerError", err)
	}

	if items == nil {
		items = make([]model.APost, 0)
	}

	return &ListPostOutput{
		Body: struct {
			Items []model.APost `json:"items"`
		}{
			Items: items,
		},
	}, nil
}

func (s *ListPost) Register(api huma.API) {
	huma.Register(api, s.GetMeta(), s.Handler)
}
