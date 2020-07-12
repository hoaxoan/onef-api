package onef_posts

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type PostUsecase interface {
	Get(ctx context.Context, req *model.PostRequest, res *model.PostResponse) error
	GetWithId(ctx context.Context, id int64, res *model.PostResponse) error
	CreatePost(ctx context.Context, req *model.CreatePostRequest, res *model.PostResponse) error
	PublishPost(ctx context.Context, postUuid string, res *model.PostResponse) error
	GetPostWithUuid(ctx context.Context, postUuid string, res *model.PostResponse) error
	Update(ctx context.Context, req *model.Post, res *model.PostResponse) error
	Delete(ctx context.Context, req *model.Post, res *model.PostResponse) error
}
