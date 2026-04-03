package application

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	applicationservice "github.com/unblvvv/h-www-server/internal/service/application"
)

type UpdateRequestDto struct {
	ID   string `path:"id" format:"uuid" doc:"ticket id"`
	Body struct {
		Status string `json:"status" enum:"new,in_progress,resolved,rejected" doc:"new, in_progress, resolved, rejected"`
	}
}

type UpdateResponseOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

type UpdateHandler struct {
	service *applicationservice.Service
}

func NewUpdateHandler(service *applicationservice.Service) *UpdateHandler {
	return &UpdateHandler{service: service}
}

func (h *UpdateHandler) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "admin-applications-update",
		Method:      "PATCH",
		Path:        "/admin/applications/{id}",
		Tags:        []string{"Admin Inbox"},
		Description: "change application status",
		Security: []map[string][]string{
			{"admin_bearer": {}},
		},
	}
}

func (h *UpdateHandler) Handle(ctx context.Context, input *UpdateRequestDto) (*UpdateResponseOutput, error) {
	err := h.service.ResolveApplication(ctx, input.ID, input.Body.Status)
	if err != nil {
		return nil, huma.Error500InternalServerError("internal server error", err)
	}

	return &UpdateResponseOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "status updated successfully",
		},
	}, nil
}

func (h *UpdateHandler) Register(api huma.API) {
	huma.Register(api, h.GetMeta(), h.Handle)
}
