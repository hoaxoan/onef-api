package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_hashtags"
)

type usecase struct {
	Repo onef_hashtags.Repository
}

func NewUsecase(repo onef_hashtags.Repository) onef_hashtags.Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (ucase *usecase) Get(ctx context.Context, req *model.HashtagRequest, res *model.HashtagResponse) error {
	hashtags, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Hashtags = hashtags
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Hashtag, res *model.HashtagResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Hashtag = req
	return nil
}

func (ucase *usecase) Update(ctx context.Context, req *model.Hashtag, res *model.HashtagResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Hashtag = req
	return nil
}

func (ucase *usecase) Delete(ctx context.Context, req *model.Hashtag, res *model.HashtagResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Hashtag = req
	return nil
}
