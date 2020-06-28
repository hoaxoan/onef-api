package onef_auth

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	GetAll(ctx context.Context, req *model.UserRequest, res *model.UserResponse) error
	Get(ctx context.Context, id string, res *model.UserResponse) error
	Create(ctx context.Context, req *model.User, res *model.UserResponse) error
	Update(ctx context.Context, req *model.User, res *model.UserResponse) error
	Auth(ctx context.Context, req *model.User, res *model.Token) error
	ValidateToken(ctx context.Context, req *model.Token, res *model.Token) error
	GetToken(ctx context.Context, req *model.User, res *model.Token) error
}
