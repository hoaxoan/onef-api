package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
)

const STATUS_DRAFT = "D"
const STATUS_PROCESSING = "PG"
const STATUS_PUBLISHED = "P"

type postUsecase struct {
	Repo onef_posts.PostRepository
}

func NewPostUsecase(repo onef_posts.PostRepository) onef_posts.PostUsecase {
	return &postUsecase{
		Repo: repo,
	}
}

func (ucase *postUsecase) Get(ctx context.Context, req *model.PostRequest, res *model.PostResponse) error {
	posts, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Posts = posts
	return nil
}

func (ucase *postUsecase) GetWithId(ctx context.Context, id int64, res *model.PostResponse) error {
	post, err := ucase.Repo.GetWithId(id)
	if err != nil {
		return err
	}
	res.Post = post
	return nil
}

func (ucase *postUsecase) CreatePost(ctx context.Context, req *model.CreatePostRequest, res *model.PostResponse) error {
	post := model.Post{UUId: req.UUId, Text: req.Text, Creator: req.Creator, CreatorId: &req.Creator.Id, Created: time.Now()}
	if req.IsDraft {
		post.Status = STATUS_DRAFT
	} else {
		post.Status = STATUS_PUBLISHED
	}

	if err := ucase.Repo.Create(&post); err != nil {
		return err
	}

	if req.CircleIds != nil {
		for _, id := range req.CircleIds {
			circleId, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				continue
			}

			if err := ucase.Repo.CreateCirclePost(circleId, post.Id); err != nil {
				return nil
			}
		}
	}
	res.Post = &post
	return nil
}

func (ucase *postUsecase) PublishPost(ctx context.Context, postUuid string, res *model.PostResponse) error {
	if err := ucase.Repo.UpdateStatus(postUuid, STATUS_PUBLISHED); err != nil {
		return err
	}
	return nil
}

func (ucase *postUsecase) GetPostWithUuid(ctx context.Context, postUuid string, res *model.PostResponse) error {
	post, err := ucase.Repo.GetPostWithUuid(postUuid)
	if err != nil {
		return err
	}
	res.Post = post
	return nil
}

func (ucase *postUsecase) Update(ctx context.Context, req *model.Post, res *model.PostResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Post = req
	return nil
}

func (ucase *postUsecase) Delete(ctx context.Context, req *model.Post, res *model.PostResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Post = req
	return nil
}
