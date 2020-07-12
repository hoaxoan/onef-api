package onef_posts

import "github.com/hoaxoan/onef-api/onef_core/model"

type PostsRepository interface {
	GetTopPosts(req *model.TopPostRequest) ([]model.TopPost, error)
	GetTrendingPosts(req *model.TrendingPostRequest) ([]model.TrendingPost, error)
	GetTimelinePosts(req *model.PostRequest) ([]model.Post, error)
}
