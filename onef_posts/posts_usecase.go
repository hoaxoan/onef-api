package onef_posts

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type PostsUsecase interface {
	GetTopPosts(ctx context.Context, req *model.TopPostRequest, res *model.TopPostResponse) error
	GetTrendingPosts(ctx context.Context, req *model.TrendingPostRequest, res *model.TrendingPostResponse) error
	GetTimelinePosts(ctx context.Context, req *model.PostRequest, res *model.PostResponse) error
}
