package onef_common

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	GetCategories(req *model.CategoryRequest) ([]*model.Category, error)
	CreateCategories(categories []*model.Category) error
	CreateCategory(category *model.Category) error
	UpdateCategory(category *model.Category) error
	DeleteCategory(category *model.Category) error

	GetMoods(req *model.MoodRequest) ([]*model.Mood, error)
	CreateMoods(moods []*model.Mood) error
	CreateMood(mood *model.Mood) error
	UpdateMood(mood *model.Mood) error
	DeleteMood(mood *model.Mood) error

	GetColorRanges(req *model.ColorRangeRequest) ([]*model.ColorRange, error)
	CreateColorRanges(colorRanges []*model.ColorRange) error
	CreateColorRange(colorRange *model.ColorRange) error
	UpdateColorRange(colorRange *model.ColorRange) error
	DeleteColorRange(colorRange *model.ColorRange) error
}
