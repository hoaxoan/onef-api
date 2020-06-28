package onef_task

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	GetAll(ctx context.Context, req *model.TaskRequest, res *model.TaskResponse) error
	Get(ctx context.Context, id string, res *model.TaskResponse) error
	Create(ctx context.Context, req *model.Task, res *model.TaskResponse) error
	Update(ctx context.Context, req *model.Task, res *model.TaskResponse) error
}
