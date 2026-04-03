package admin

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/model"
	postservice "github.com/unblvvv/h-www-server/internal/service/post"
)

type UpdatePostRequestDto struct {
	ID   string `path:"id" doc:"ID записи"`
	Body struct {
		Name        string        `json:"name"`
		Age         string        `json:"age"`
		Sex         model.ASex    `json:"sex"`
		Description string        `json:"description"`
		PhotoURLs    []string       `json:"photo_urls,omitempty"`
		Status      model.AStatus `json:"status"`
	}
}

type UpdatePostOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

type Update struct {
	service *postservice.Service
}

func NewUpdate(s *postservice.Service) *Update {
	return &Update{service: s}
}

func (s *Update) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "admin-update-animal",
		Method:      "PUT",
		Path:        "/admin/animal/update/{id}",
		Tags:        []string{"admin"},
		Security:    []map[string][]string{{"admin_bearer": {}}},
	}
}

func (s *Update) Handler(ctx context.Context, input *UpdatePostRequestDto) (*UpdatePostOutput, error) {
	p := model.APost{
		ID:          input.ID,
		Name:        input.Body.Name,
		Age:         input.Body.Age,
		Sex:         input.Body.Sex,
		Description: input.Body.Description,
		PhotoURLs:    input.Body.PhotoURLs,
		Status:      input.Body.Status,
	}

	if err := s.service.Update(ctx, p); err != nil {
		return nil, huma.Error500InternalServerError("internal server error", err)
	}

	return &UpdatePostOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "Post updated successfully",
		},
	}, nil
}

func (s *Update) Register(api huma.API) {
	huma.Register(api, s.GetMeta(), s.Handler)
}
