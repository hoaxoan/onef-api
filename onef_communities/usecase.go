package onef_communities

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.CommunityRequest, res *model.CommunityResponse) error
	Create(ctx context.Context, req *model.Community, res *model.CommunityResponse) error
	Update(ctx context.Context, req *model.Community, res *model.CommunityResponse) error
	Delete(ctx context.Context, req *model.Community, res *model.CommunityResponse) error
}
