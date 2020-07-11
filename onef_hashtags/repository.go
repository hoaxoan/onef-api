package onef_hashtags

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	Get(req *model.HashtagRequest) ([]model.Hashtag, error)
	Create(category *model.Hashtag) error
	Update(category *model.Hashtag) error
	Delete(category *model.Hashtag) error
}
