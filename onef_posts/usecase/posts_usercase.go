package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
)

type postsUsecase struct {
	Repo onef_posts.PostsRepository
}

func NewUsecase(repo onef_posts.PostsRepository) onef_posts.PostsUsecase {
	return &postsUsecase{
		Repo: repo,
	}
}

func (ucase *postsUsecase) GetTopPosts(ctx context.Context, req *model.TopPostRequest, res *model.TopPostResponse) error {
	topPosts, err := ucase.Repo.GetTopPosts(req)
	if err != nil {
		return err
	}
	res.TopPosts = topPosts
	return nil
}

func (ucase *postsUsecase) GetTrendingPosts(ctx context.Context, req *model.TrendingPostRequest, res *model.TrendingPostResponse) error {
	trendingPosts, err := ucase.Repo.GetTrendingPosts(req)
	if err != nil {
		return err
	}
	res.TrendingPosts = trendingPosts
	return nil
}

func (ucase *postsUsecase) GetTimelinePosts(ctx context.Context, req *model.PostRequest, res *model.PostResponse) error {
	posts, err := ucase.Repo.GetTimelinePosts(req)
	if err != nil {
		return err
	}
	res.Posts = posts
	return nil
}
