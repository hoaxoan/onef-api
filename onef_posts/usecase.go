package onef_categories

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.CategoryRequest, res *model.CategoryResponse) error
	Create(ctx context.Context, req *model.Category, res *model.CategoryResponse) error
	Update(ctx context.Context, req *model.Category, res *model.CategoryResponse) error
	Delete(ctx context.Context, req *model.Category, res *model.CategoryResponse) error
}
