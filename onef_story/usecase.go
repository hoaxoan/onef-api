package onef_story

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	GetAll(ctx context.Context, req *model.StoryRequest, res *model.StoryResponse) error
	Get(ctx context.Context, id string, res *model.StoryResponse) error
	Create(ctx context.Context, req *model.Story, res *model.StoryResponse) error
	Update(ctx context.Context, req *model.Story, res *model.StoryResponse) error
	Delete(ctx context.Context, req *model.Story, res *model.StoryResponse) error
}
