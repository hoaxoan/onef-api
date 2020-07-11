package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_categories"
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type usecase struct {
	Repo onef_categories.Repository
}

func NewUsecase(repo onef_categories.Repository) onef_categories.Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (ucase *usecase) Get(ctx context.Context, req *model.CategoryRequest, res *model.CategoryResponse) error {
	categories, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Categories = categories
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Category, res *model.CategoryResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Category = req
	return nil
}

func (ucase *usecase) Update(ctx context.Context, req *model.Category, res *model.CategoryResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Category = req
	return nil
}

func (ucase *usecase) Delete(ctx context.Context, req *model.Category, res *model.CategoryResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Category = req
	return nil
}
