package onef_circles

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.CircleRequest, res *model.CircleResponse) error
	Create(ctx context.Context, req *model.Circle, res *model.CircleResponse) error
	Update(ctx context.Context, req *model.Circle, res *model.CircleResponse) error
	Delete(ctx context.Context, req *model.Circle, res *model.CircleResponse) error
}
