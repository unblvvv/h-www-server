package post

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"github.com/unblvvv/h-www-server/internal/service/storage"
)

type UploadRequestDto struct {
	RawBody multipart.Form `contentType:"multipart/form-data"`
}

type UploadOutput struct {
	Body struct {
		URLs []string `json:"urls" doc:"links to uploaded files"`
	}
}

type Handler struct {
	storage *storage.R2Service
}

func NewUpload(storage *storage.R2Service) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "upload-image",
		Method:      "POST",
		Path:        "/v1/upload",
		Tags:        []string{"Media"},
		Description: "Upload an image file and get its URL in response",
		Security: []map[string][]string{
			{"admin_bearer": {}},
		},
	}
}

func (h *Handler) Handle(ctx context.Context, input *UploadRequestDto) (*UploadOutput, error) {
	files := input.RawBody.File["file"]
	if len(files) == 0 {
		return nil, huma.Error400BadRequest("file is required")
	}

	if len(files) > 10 {
		return nil, huma.Error400BadRequest("maximum 10 photos allowed")
	}

	var urls []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, huma.Error400BadRequest("bad request", err)
		}

		ext := filepath.Ext(fileHeader.Filename)
		filename := "photos/" + uuid.New().String() + ext

		contentType := fileHeader.Header.Get("Content-Type")
		if contentType == "" {
			contentType = "image/jpeg"
		}

		fileURL, err := h.storage.UploadFile(ctx, file, filename, contentType)
		file.Close()
		if err != nil {
			return nil, huma.Error500InternalServerError("cloudflare save error", err)
		}

		urls = append(urls, fileURL)
	}

	return &UploadOutput{
		Body: struct {
			URLs []string `json:"urls" doc:"links to uploaded files"`
		}{URLs: urls},
	}, nil
}

func (h *Handler) Register(api huma.API) {
	huma.Register(api, h.GetMeta(), h.Handle)
}
