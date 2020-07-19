package onef_circles

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.CircleRequest, res *model.CircleResponse) error
	GetCircleWithId(ctx context.Context, circleId int64, res *model.CircleResponse) error
	Create(ctx context.Context, req *model.Circle, res *model.CircleResponse) error
	UpdateCircleWithId(ctx context.Context, req *model.UpdateCircleRequest, res *model.CircleResponse) error
	DeleteCircleWithId(ctx context.Context, circleId int64, res *model.CircleResponse) error
	CheckNameIsAvailable(ctx context.Context, req *model.Circle, res *model.CircleResponse) error
}
