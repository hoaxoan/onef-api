package onef_auth

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	GetAll() ([]*model.User, error)
	GetWithId(id int64) (*model.User, error)
	CreateUser(req *model.RegisterRequest) (*model.User, error)
	Update(user *model.User) error

	GetUserWithEmail(email string) (*model.User, error)
	GetUserWithUserName(userName string) (*model.User, error)
	IsUserNameToken(userName string) bool
	IsEmailToken(email string) bool
}
