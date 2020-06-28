package onef_common

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	GetAll() ([]*model.Story, error)
	Get(id int) (*model.Story, error)
	Create(story *model.Story) error
	Update(story *model.Story) error
}
