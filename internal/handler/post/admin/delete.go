package admin

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/repository/post"
	postservice "github.com/unblvvv/h-www-server/internal/service/post"
)

type DeleteAnimalRequestDto struct {
	ID string `path:"id" format:"uuid" doc:"post id"`
}

type DeletePostOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

type DeletePost struct {
	service *postservice.Service
	repo    post.Repository
}

func NewDeletePost(service *postservice.Service, repo post.Repository) *DeletePost {
	return &DeletePost{service: service, repo: repo}
}

func (s *DeletePost) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "admin-delete-animal",
		Method:      "DELETE",
		Path:        "/admin/animal/delete/{id}",
		Tags:        []string{"admin"},
		Description: "Delete an animal post",
		Security: []map[string][]string{
			{"admin_bearer": {}},
		},
	}
}

func (s *DeletePost) Handler(ctx context.Context, input *DeleteAnimalRequestDto) (*DeletePostOutput, error) {
	if err := s.service.Delete(ctx, input.ID); err != nil {
		return nil, huma.Error500InternalServerError("internal server error", err)
	}

	return &DeletePostOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "Post deleted successfully",
		},
	}, nil
}

func (s *DeletePost) Register(api huma.API) {
	huma.Register(api, s.GetMeta(), s.Handler)
}
