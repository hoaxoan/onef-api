package onef_posts

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type PostRepository interface {
	Get(req *model.PostRequest) ([]model.Post, error)
	GetWithId(id int64) (*model.Post, error)
	GetPostWithUuid(postUuuid string) (*model.Post, error)
	UpdateStatus(postUuuid string, status string) error
	Create(post *model.Post) error
	Update(post *model.Post) error
	Delete(post *model.Post) error
	CreateCirclePost(circleId int64, postId int64) error
}
