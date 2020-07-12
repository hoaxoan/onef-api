package onef_posts

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.PostRequest, res *model.PostResponse) error
	GetWithId(ctx context.Context, id int64, res *model.PostResponse) error
	Create(ctx context.Context, req *model.Post, res *model.PostResponse) error
	Update(ctx context.Context, req *model.Post, res *model.PostResponse) error
	Delete(ctx context.Context, req *model.Post, res *model.PostResponse) error
}
