package onef_communities

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	Get(req *model.CommunityRequest) ([]model.Community, error)
	Create(category *model.Community) error
	Update(category *model.Community) error
	Delete(category *model.Community) error
}
