package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
)

type usecase struct {
	Repo onef_posts.Repository
}

func NewUsecase(repo onef_posts.Repository) onef_posts.Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (ucase *usecase) Get(ctx context.Context, req *model.PostRequest, res *model.PostResponse) error {
	posts, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Posts = posts
	return nil
}

func (ucase *usecase) GetWithId(ctx context.Context, id int64, res *model.PostResponse) error {
	post, err := ucase.Repo.GetWithId(id)
	if err != nil {
		return err
	}
	res.Post = post
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Post, res *model.PostResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Post = req
	return nil
}

func (ucase *usecase) Update(ctx context.Context, req *model.Post, res *model.PostResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Post = req
	return nil
}

func (ucase *usecase) Delete(ctx context.Context, req *model.Post, res *model.PostResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Post = req
	return nil
}
