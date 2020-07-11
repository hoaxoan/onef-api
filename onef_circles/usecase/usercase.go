package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_circles"
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type usecase struct {
	Repo onef_circles.Repository
}

func NewUsecase(repo onef_circles.Repository) onef_circles.Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (ucase *usecase) Get(ctx context.Context, req *model.CircleRequest, res *model.CircleResponse) error {
	circles, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Circles = circles
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Circle, res *model.CircleResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Circle = req
	return nil
}

func (ucase *usecase) Update(ctx context.Context, req *model.Circle, res *model.CircleResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Circle = req
	return nil
}

func (ucase *usecase) Delete(ctx context.Context, req *model.Circle, res *model.CircleResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Circle = req
	return nil
}
