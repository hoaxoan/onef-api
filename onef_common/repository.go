package onef_categories

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	Get(req *model.CategoryRequest) ([]model.Category, error)
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(category *model.Category) error
}
