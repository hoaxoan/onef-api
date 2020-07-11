package onef_connections

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	Get(req *model.ConnectionRequest) ([]model.Connection, error)
	Create(category *model.Connection) error
	Update(category *model.Connection) error
	Delete(category *model.Connection) error
}
