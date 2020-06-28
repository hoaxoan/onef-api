package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_common"
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type commonUsecase struct {
	Repo onef_common.Repository
}

func NewUsecase(repo onef_common.Repository) onef_common.Usecase {
	return &commonUsecase{
		Repo: repo,
	}
}

// Categories
func (ucase *commonUsecase) GetCategories(ctx context.Context, req *model.CategoryRequest, res *model.CategoryResponse) error {
	categories, err := ucase.Repo.GetCategories(req)
	if err != nil {
		return err
	}
	res.Categories = categories
	return nil
}

func (ucase *commonUsecase) CreateCategories(ctx context.Context, req []*model.Category, res *model.CategoryResponse) error {
	if err := ucase.Repo.CreateCategories(req); err != nil {
		return err
	}
	res.Categories = req
	return nil
}

func (ucase *commonUsecase) CreateCategory(ctx context.Context, req *model.Category, res *model.CategoryResponse) error {
	if err := ucase.Repo.CreateCategory(req); err != nil {
		return err
	}
	res.Category = req
	return nil
}

func (ucase *commonUsecase) UpdateCategory(ctx context.Context, req *model.Category, res *model.CategoryResponse) error {
	if err := ucase.Repo.UpdateCategory(req); err != nil {
		return err
	}
	res.Category = req
	return nil
}

func (ucase *commonUsecase) DeleteCategory(ctx context.Context, req *model.Category, res *model.CategoryResponse) error {
	if err := ucase.Repo.DeleteCategory(req); err != nil {
		return err
	}
	res.Category = req
	return nil
}

// Moods
func (ucase *commonUsecase) GetMoods(ctx context.Context, req *model.MoodRequest, res *model.MoodResponse) error {
	moods, err := ucase.Repo.GetMoods(req)
	if err != nil {
		return err
	}
	res.Moods = moods
	return nil
}

func (ucase *commonUsecase) CreateMoods(ctx context.Context, req []*model.Mood, res *model.MoodResponse) error {
	if err := ucase.Repo.CreateMoods(req); err != nil {
		return err
	}
	res.Moods = req
	return nil
}

func (ucase *commonUsecase) CreateMood(ctx context.Context, req *model.Mood, res *model.MoodResponse) error {
	if err := ucase.Repo.CreateMood(req); err != nil {
		return err
	}
	res.Mood = req
	return nil
}

func (ucase *commonUsecase) UpdateMood(ctx context.Context, req *model.Mood, res *model.MoodResponse) error {
	if err := ucase.Repo.UpdateMood(req); err != nil {
		return err
	}
	res.Mood = req
	return nil
}

func (ucase *commonUsecase) DeleteMood(ctx context.Context, req *model.Mood, res *model.MoodResponse) error {
	if err := ucase.Repo.DeleteMood(req); err != nil {
		return err
	}
	res.Mood = req
	return nil
}

// Color Ranges
func (ucase *commonUsecase) GetColorRanges(ctx context.Context, req *model.ColorRangeRequest, res *model.ColorRangeResponse) error {
	colorRanges, err := ucase.Repo.GetColorRanges(req)
	if err != nil {
		return err
	}
	res.ColorRanges = colorRanges
	return nil
}

func (ucase *commonUsecase) CreateColorRanges(ctx context.Context, req []*model.ColorRange, res *model.ColorRangeResponse) error {
	if err := ucase.Repo.CreateColorRanges(req); err != nil {
		return err
	}
	res.ColorRanges = req
	return nil
}

func (ucase *commonUsecase) CreateColorRange(ctx context.Context, req *model.ColorRange, res *model.ColorRangeResponse) error {
	if err := ucase.Repo.CreateColorRange(req); err != nil {
		return err
	}
	res.ColorRange = req
	return nil
}

func (ucase *commonUsecase) UpdateColorRange(ctx context.Context, req *model.ColorRange, res *model.ColorRangeResponse) error {
	if err := ucase.Repo.UpdateColorRange(req); err != nil {
		return err
	}
	res.ColorRange = req
	return nil
}

func (ucase *commonUsecase) DeleteColorRange(ctx context.Context, req *model.ColorRange, res *model.ColorRangeResponse) error {
	if err := ucase.Repo.DeleteColorRange(req); err != nil {
		return err
	}
	res.ColorRange = req
	return nil
}
