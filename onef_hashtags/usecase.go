package onef_hashtags

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.HashtagRequest, res *model.HashtagResponse) error
	Create(ctx context.Context, req *model.Hashtag, res *model.HashtagResponse) error
	Update(ctx context.Context, req *model.Hashtag, res *model.HashtagResponse) error
	Delete(ctx context.Context, req *model.Hashtag, res *model.HashtagResponse) error
}
