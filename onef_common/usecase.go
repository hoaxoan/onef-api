package onef_common

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	GetCategories(ctx context.Context, req *model.CategoryRequest, res *model.CategoryResponse) error
	CreateCategories(ctx context.Context, req []*model.Category, res *model.CategoryResponse) error
	CreateCategory(ctx context.Context, req *model.Category, res *model.CategoryResponse) error
	UpdateCategory(ctx context.Context, req *model.Category, res *model.CategoryResponse) error
	DeleteCategory(ctx context.Context, req *model.Category, res *model.CategoryResponse) error

	GetMoods(ctx context.Context, req *model.MoodRequest, res *model.MoodResponse) error
	CreateMoods(ctx context.Context, req []*model.Mood, res *model.MoodResponse) error
	CreateMood(ctx context.Context, req *model.Mood, res *model.MoodResponse) error
	UpdateMood(ctx context.Context, req *model.Mood, res *model.MoodResponse) error
	DeleteMood(ctx context.Context, req *model.Mood, res *model.MoodResponse) error

	GetColorRanges(ctx context.Context, req *model.ColorRangeRequest, res *model.ColorRangeResponse) error
	CreateColorRanges(ctx context.Context, req []*model.ColorRange, res *model.ColorRangeResponse) error
	CreateColorRange(ctx context.Context, req *model.ColorRange, res *model.ColorRangeResponse) error
	UpdateColorRange(ctx context.Context, req *model.ColorRange, res *model.ColorRangeResponse) error
	DeleteColorRange(ctx context.Context, req *model.ColorRange, res *model.ColorRangeResponse) error
}
