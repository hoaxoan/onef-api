package onef_posts

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type PostsUsecase interface {
	GetTopPosts(ctx context.Context, req *model.TopPostRequest, res *model.TopPostResponse) error
	GetTrendingPosts(ctx context.Context, req *model.TrendingPostRequest, res *model.TrendingPostResponse) error
	GetTimelinePosts(ctx context.Context, req *model.PostRequest, res *model.PostResponse) error
	GetPostComment(ctx context.Context, postUuid string, postCommentId int64, res *model.PostCommentResponse) error
	GetCommentsForPostWithUuid(ctx context.Context, req *model.PostCommentRequest, res *model.PostCommentResponse) error
	GetRepliesForCommentWithIdForPostWithUuid(ctx context.Context, req *model.PostCommentRequest, res *model.PostCommentResponse) error
}
