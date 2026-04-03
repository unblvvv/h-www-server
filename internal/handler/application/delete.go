package application

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/service/application"
)

type DeleteRequestDto struct {
	ID string `path:"id" doc:"ID заявки" format:"uuid"`
}

type DeleteOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

type DeleteHandler struct {
	service *application.Service
}

func NewDeleteHandler(service *application.Service) *DeleteHandler {
	return &DeleteHandler{service: service}
}

func (h *DeleteHandler) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "delete-application",
		Method:      http.MethodDelete,
		Path:        "/admin/applications/{id}",
		Tags:        []string{"Admin Applications"},
		Description: "delete application by id",
		Security: []map[string][]string{
			{"admin_bearer": {}},
		},
	}
}

func (h *DeleteHandler) Handle(ctx context.Context, input *DeleteRequestDto) (*DeleteOutput, error) {
	if err := h.service.Delete(ctx, input.ID); err != nil {
		return nil, huma.Error500InternalServerError("internal server error", err)
	}

	return &DeleteOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "Application deleted successfully",
		},
	}, nil
}

func (h *DeleteHandler) Register(api huma.API) {
	huma.Register(api, h.GetMeta(), h.Handle)
}
