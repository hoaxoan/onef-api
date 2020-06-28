package onef_story

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	GetAll() ([]*model.Story, error)
	Get(id string) (*model.Story, error)
	Create(story *model.Story) error
	Update(story *model.Story) error
	Delete(story *model.Story) error
}
