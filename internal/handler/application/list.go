package application

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/model"
	applicationservice "github.com/unblvvv/h-www-server/internal/service/application"
)

type ListRequestDto struct {
	Status string `query:"status" doc:"new, in_progress, resolved, rejected"`
	Limit  int    `query:"limit" default:"20" minimum:"1" maximum:"100"`
	Offset int    `query:"offset" default:"0" minimum:"0"`
}

type ListResponseOutput struct {
	Body struct {
		Items []model.Application `json:"items"`
		Total int                 `json:"total" doc:"total number of items"`
	}
}

type ListHandler struct {
	service *applicationservice.Service
}

func NewListHandler(service *applicationservice.Service) *ListHandler {
	return &ListHandler{service: service}
}

func (h *ListHandler) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "admin-applications-list",
		Method:      "GET",
		Path:        "/admin/applications",
		Tags:        []string{"Admin Inbox"},
		Description: "get a list of applications",
		Security: []map[string][]string{
			{"admin_bearer": {}},
		},
	}
}

func (h *ListHandler) Handle(ctx context.Context, input *ListRequestDto) (*ListResponseOutput, error) {
	var statusPtr *string

	if input.Status != "" {
		statusPtr = &input.Status
	}

	items, total, err := h.service.GetInbox(ctx, statusPtr, input.Limit, input.Offset)
	if err != nil {
		return nil, huma.Error500InternalServerError("internal server error", err)
	}

	return &ListResponseOutput{
		Body: struct {
			Items []model.Application `json:"items"`
			Total int                 `json:"total" doc:"total number of items"`
		}{
			Items: items,
			Total: total,
		},
	}, nil
}

func (h *ListHandler) Register(api huma.API) {
	huma.Register(api, h.GetMeta(), h.Handle)
}
