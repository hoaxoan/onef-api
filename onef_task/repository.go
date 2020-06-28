package onef_task

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	GetAll() ([]*model.Task, error)
	Get(id string) (*model.Task, error)
	Create(task *model.Task) error
	Update(task *model.Task) error
}
