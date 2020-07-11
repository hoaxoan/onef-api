package onef_circles

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	Get(req *model.CircleRequest) ([]model.Circle, error)
	Create(circle *model.Circle) error
	Update(circle *model.Circle) error
	Delete(circle *model.Circle) error
}
