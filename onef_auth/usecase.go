package onef_auth

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Register(ctx context.Context, req *model.RegisterRequest, res *model.UserResponse) error
	Login(ctx context.Context, req *model.LoginRequest, res *model.Token) error
	GetUserWithUserName(ctx context.Context, userName string, res *model.UserResponse) error
	GetUserWithEmail(ctx context.Context, email string, res *model.UserResponse) error

	GetAll(ctx context.Context, req *model.UserRequest, res *model.UserResponse) error
	Get(ctx context.Context, id int, res *model.UserResponse) error
	Update(ctx context.Context, req *model.User, res *model.UserResponse) error
	ValidateToken(ctx context.Context, req *model.Token, res *model.Token) error
	GetToken(ctx context.Context, req *model.User, res *model.Token) error
}
