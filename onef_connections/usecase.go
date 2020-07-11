package onef_connections

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.ConnectionRequest, res *model.ConnectionResponse) error
	Create(ctx context.Context, req *model.Connection, res *model.ConnectionResponse) error
	Update(ctx context.Context, req *model.Connection, res *model.ConnectionResponse) error
	Delete(ctx context.Context, req *model.Connection, res *model.ConnectionResponse) error
}
