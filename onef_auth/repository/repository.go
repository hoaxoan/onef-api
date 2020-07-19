package repository

import (
	"errors"

	"github.com/hoaxoan/onef-api/onef_auth"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_auth.Repository {
	return &userRepository{db}
}

func (repo *userRepository) CreateUser(req *model.RegisterRequest) (*model.User, error) {
	// Validator
	if !req.IsOfLegalAge {
		return nil, errors.New("You must confirm you are over 16 years old to make an account.")
	}

	if !req.IsOfLegalAge {
		return nil, errors.New("You must accept the guidelines to make an account.")
	}

	if repo.IsUserNameToken(req.UserName) {
		return nil, errors.New("The username is already taken.")
	}

	if repo.IsEmailToken(req.Email) {
		return nil, errors.New("The email is already taken.")
	}

	// Create User
	user := model.User{UserName: req.UserName, Email: req.Email, Password: req.Password, AreGuidelinesAccepted: req.AreGuidelinesAccepted}
	err := repo.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Create UserProfile
	userProfile := model.UserProfile{User: &user, Name: req.Name, Avatar: req.Avatar, IsOfLegalAge: req.IsOfLegalAge, UserId: user.Id}
	if err := repo.db.Create(&userProfile).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	if dbc := repo.db.Find(&users); dbc.Error != nil {
		return nil, dbc.Error
	}
	return users, nil
}

func (repo *userRepository) GetWithId(id int64) (*model.User, error) {
	var user *model.User
	if dbc := repo.db.Where("id = ?", id).Take(&user); dbc.Error != nil {
		return nil, dbc.Error
	}
	return user, nil
}

func (repo *userRepository) Update(user *model.User) error {
	if dbc := repo.db.Where("id = ?", user.Id).Updates(&user); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *userRepository) GetUserWithUserName(userName string) (*model.User, error) {
	var user model.User
	if dbc := repo.db.Where("username = ?", userName).First(&user); dbc.Error != nil {
		return nil, dbc.Error
	}

	var profile model.UserProfile
	if dbc := repo.db.Where("user_id = ?", user.Id).First(&profile); dbc.Error != nil {
		return nil, dbc.Error
	}

	user.Profile = &profile

	var circle model.Circle
	if dbc := repo.db.Where("creator_id = ?", user.Id).First(&circle); dbc.Error == nil {
		user.ConnectionsCircleId = circle.Id
	}

	return &user, nil
}

func (repo *userRepository) GetUserWithEmail(email string) (*model.User, error) {
	var user model.User
	if dbc := repo.db.Where("email = ?", email).First(&user).Related(&user.Profile); dbc.Error != nil {
		return nil, dbc.Error
	}

	// if dbc := repo.db.Where("user_id = ?", user.Id).First(&user.Profile); dbc.Error != nil {
	// 	return nil, dbc.Error
	// }

	return &user, nil
}

func (repo *userRepository) IsUserNameToken(userName string) bool {
	var user model.User
	if dbc := repo.db.Where("username = ?", userName).First(&user); dbc.Error != nil {
		return false
	}

	if user.UserName == userName {
		return true
	}

	return false
}

func (repo *userRepository) IsEmailToken(email string) bool {
	var user model.User
	if dbc := repo.db.Where("email = ?", email).First(&user); dbc.Error != nil {
		return false
	}

	if user.Email == email {
		return true
	}

	return false
}
