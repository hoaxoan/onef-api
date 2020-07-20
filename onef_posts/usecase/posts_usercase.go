package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
)

type postsUsecase struct {
	Repo onef_posts.PostsRepository
}

func NewPostsUsecase(repo onef_posts.PostsRepository) onef_posts.PostsUsecase {
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

func (ucase *postsUsecase) GetPostComment(ctx context.Context, postUuid string, postCommentId int64, res *model.PostCommentResponse) error {
	postComment, err := ucase.Repo.GetPostComment(postUuid, postCommentId)
	if err != nil {
		return err
	}
	res.PostComment = postComment
	return nil
}

func (ucase *postsUsecase) GetCommentsForPostWithUuid(ctx context.Context, req *model.PostCommentRequest, res *model.PostCommentResponse) error {
	postComments, err := ucase.Repo.GetCommentsForPostWithUuid(req)
	if err != nil {
		return err
	}
	res.PostComments = postComments
	return nil
}

func (ucase *postsUsecase) GetRepliesForCommentWithIdForPostWithUuid(ctx context.Context, req *model.PostCommentRequest, res *model.PostCommentResponse) error {
	postComments, err := ucase.Repo.GetRepliesForCommentWithIdForPostWithUuid(req)
	if err != nil {
		return err
	}
	res.PostComments = postComments
	return nil
}
