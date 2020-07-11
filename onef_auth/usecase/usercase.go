package usecase

import (
	"context"
	"log"

	"github.com/hoaxoan/onef-api/onef_auth"
	"github.com/hoaxoan/onef-api/onef_auth/service"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	Repo         onef_auth.Repository
	TokenService service.Authable
}

func NewUsecase(repo onef_auth.Repository, tokenService service.Authable) onef_auth.Usecase {
	return &userUsecase{
		Repo:         repo,
		TokenService: tokenService,
	}
}

func (ucase *userUsecase) Register(ctx context.Context, req *model.RegisterRequest, res *model.UserResponse) error {
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)

	user, err := ucase.Repo.CreateUser(req)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (ucase *userUsecase) Login(ctx context.Context, req *model.LoginRequest, res *model.Token) error {
	user, err := ucase.Repo.GetUserWithUserName(req.UserName)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(string(req.Password))); err != nil {
		return err
	}

	token, err := ucase.TokenService.Encode(user)
	if err != nil {
		return err
	}
	res.UserName = user.UserName
	res.Email = user.Email
	res.Token = token
	return nil
}

func (ucase *userUsecase) GetUserWithUserName(ctx context.Context, userName string, res *model.UserResponse) error {
	user, err := ucase.Repo.GetUserWithUserName(userName)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (ucase *userUsecase) GetUserWithEmail(ctx context.Context, email string, res *model.UserResponse) error {
	user, err := ucase.Repo.GetUserWithEmail(email)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (ucase *userUsecase) GetAll(ctx context.Context, req *model.UserRequest, res *model.UserResponse) error {
	users, err := ucase.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (ucase *userUsecase) Get(ctx context.Context, id int, res *model.UserResponse) error {
	user, err := ucase.Repo.Get(id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (ucase *userUsecase) Update(ctx context.Context, req *model.User, res *model.UserResponse) error {
	existUser, err := ucase.Repo.GetUserWithEmail(req.Email)
	if err != nil {
		return err
	}

	if req.Password != "" {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.Password = string(hashedPass)
	} else if existUser.Password != "" {
		req.Password = existUser.Password
	}

	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (ucase *userUsecase) ValidateToken(ctx context.Context, req *model.Token, res *model.Token) error {

	// Decode token
	claims, err := ucase.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println(claims)

	// if claims.User.Id == 0 {
	// 	return errors.New("invalid user")
	// }

	res.Valid = true

	return nil
}

func (ucase *userUsecase) GetToken(ctx context.Context, req *model.User, res *model.Token) error {
	token, err := ucase.TokenService.Encode(req)
	if err != nil {
		return err
	}
	res.UserName = req.UserName
	res.Email = req.Email
	res.Token = token
	return nil
}
