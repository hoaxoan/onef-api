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

func (ucase *usecase) GetCircleWithId(ctx context.Context, circleId int64, res *model.CircleResponse) error {
	circle, err := ucase.Repo.GetCircleWithId(circleId)
	if err != nil {
		return err
	}
	res.Circle = circle
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Circle, res *model.CircleResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Circle = req
	return nil
}

func (ucase *usecase) UpdateCircleWithId(ctx context.Context, req *model.UpdateCircleRequest, res *model.CircleResponse) error {
	circle, err := ucase.Repo.GetCircleWithId(req.Id)
	if err != nil {
		return err
	}

	circle.Name = req.Name
	circle.Color = req.Color

	if err := ucase.Repo.Update(circle); err != nil {
		return err
	}
	res.Circle = circle
	return nil
}

func (ucase *usecase) DeleteCircleWithId(ctx context.Context, circleId int64, res *model.CircleResponse) error {
	circle := model.Circle{}
	circle.Id = circleId
	if err := ucase.Repo.Delete(&circle); err != nil {
		return err
	}
	res.Circle = &circle
	return nil
}

func (ucase *usecase) CheckNameIsAvailable(ctx context.Context, req *model.Circle, res *model.CircleResponse) error {
	circle, err := ucase.Repo.CheckNameIsAvailable(req)
	if err != nil {
		return err
	}
	res.Circle = circle
	return nil
}
