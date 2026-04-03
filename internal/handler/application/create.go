package application

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	applicationservice "github.com/unblvvv/h-www-server/internal/service/application"
)

type CreateRequestDto struct {
	Body struct {
		AnimalID string `json:"animal_id" format:"uuid" doc:"animal id"`
		Name     string `json:"name" minLength:"2" maxLength:"100" doc:"username"`
		Email    string `json:"email" format:"email" doc:"email"`
		Phone    string `json:"phone" minLength:"10" doc:"phone number"`
		Message  string `json:"message" minLength:"10" maxLength:"1000" doc:"description"`
	}
}

type CreateResponseOutput struct {
	Body struct {
		Message string `json:"message" doc:"status of operation"`
	}
}

type CreateHandler struct {
	service *applicationservice.Service
}

func NewCreateHandler(service *applicationservice.Service) *CreateHandler {
	return &CreateHandler{service: service}
}

func (h *CreateHandler) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "public-application-create",
		Method:      "POST",
		Path:        "/v1/applications",
		Tags:        []string{"Public Applications"},
		Description: "Submit an application to adopt an animal",
	}
}

func (h *CreateHandler) Handle(ctx context.Context, input *CreateRequestDto) (*CreateResponseOutput, error) {
	err := h.service.CreateApplication(
		ctx,
		input.Body.AnimalID,
		input.Body.Name,
		input.Body.Email,
		input.Body.Phone,
		input.Body.Message,
	)

	if err != nil {
		return nil, huma.Error500InternalServerError("internal error", err)
	}

	return &CreateResponseOutput{
		Body: struct {
			Message string `json:"message" doc:"status of operation"`
		}{
			Message: "Your application has been submitted successfully",
		},
	}, nil
}

func (h *CreateHandler) Register(api huma.API) {
	huma.Register(api, h.GetMeta(), h.Handle)
}
