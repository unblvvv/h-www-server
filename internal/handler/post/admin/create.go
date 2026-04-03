package admin

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/unblvvv/h-www-server/internal/model"
	"github.com/unblvvv/h-www-server/internal/repository/post"
	postservice "github.com/unblvvv/h-www-server/internal/service/post"
)

type CreateAnimalRequestDto struct {
	Body struct {
		Name        string        `json:"name" minLength:"1" maxLength:"100" doc:"animal name"`
		Age         string        `json:"age" minLength:"1" maxLength:"50" doc:"age"`
		Sex         model.ASex    `json:"sex" enum:"male,female,unknown" doc:"sex"`
		Description string        `json:"description" minLength:"1" doc:"description"`
		PhotoURLs    []string       `json:"photo_urls,omitempty" format:"uri" doc:"photo url (временно)"`
		Status      model.AStatus `json:"status" enum:"available,adopted,treatment" default:"available" doc:"status (available,adopted,treatment)"`
	}
}

type CreatePostOutput struct {
	Body struct {
		ID      string `json:"id" doc:"post id"`
		Message string `json:"message"`
	}
}

type Post struct {
	service *postservice.Service
	repo    post.Repository
}

func NewPost(service *postservice.Service, repo post.Repository) *Post {
	return &Post{service: service, repo: repo}
}

func (s *Post) GetMeta() huma.Operation {
	return huma.Operation{
		OperationID: "animal-posts",
		Method:      "POST",
		Path:        "/admin/animal/create",
		Tags:        []string{"admin"},
		Description: "Create a new animal post",
		Security: []map[string][]string{
			{"admin_bearer": {}},
		},
	}
}

func (s *Post) Handler(ctx context.Context, input *CreateAnimalRequestDto) (*CreatePostOutput, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return nil, huma.Error401Unauthorized("unauthorized")
	}

	id, err := s.service.CreateAPost(
		ctx,
		userID,
		input.Body.Name,
		input.Body.Age,
		input.Body.Sex,
		input.Body.Description,
		input.Body.PhotoURLs,
		input.Body.Status,
	)

	if err != nil {
		return nil, huma.Error500InternalServerError("internal server error", err)
	}

	return &CreatePostOutput{
		Body: struct {
			ID      string `json:"id" doc:"post id"`
			Message string `json:"message"`
		}{
			ID:      id,
			Message: "Post created successfully",
		},
	}, nil
}

func (s *Post) Register(api huma.API) {
	huma.Register(api, s.GetMeta(), s.Handler)
}
