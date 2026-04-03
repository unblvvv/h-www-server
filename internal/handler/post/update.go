package post

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
		PhotoURL    *string       `json:"photo_url,omitempty"`
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
		OperationID: "update-animal",
		Method:      "PUT",
		Path:        "animal/update/{id}",
		Tags:        []string{"Posts"},
		Security:    []map[string][]string{{"bearer": {}}},
	}
}

func (s *Update) Handler(ctx context.Context, input *UpdatePostRequestDto) (*UpdatePostOutput, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return nil, huma.Error401Unauthorized("unauthorized")
	}

	p := model.APost{
		ID:          input.ID,
		Name:        input.Body.Name,
		Age:         input.Body.Age,
		Sex:         input.Body.Sex,
		Description: input.Body.Description,
		PhotoURL:    input.Body.PhotoURL,
		Status:      input.Body.Status,
	}

	if err := s.service.Update(ctx, p, userID); err != nil {
		return nil, huma.Error403Forbidden("forbidden", err)
	}

	return &UpdatePostOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "Животное успешно добавлено в базу",
		},
	}, nil
}

func (s *Update) Register(api huma.API) {
	huma.Register(api, s.GetMeta(), s.Handler)
}
