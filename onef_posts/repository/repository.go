package repository

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_posts.Repository {
	return &repository{db}
}

func (repo *repository) Get(req *model.PostRequest) ([]model.Post, error) {
	var posts []model.Post
	if dbc := repo.db.Find(&posts); dbc.Error != nil {
		return nil, dbc.Error
	}
	return posts, nil
}

func (repo *repository) GetWithId(id int64) (*model.Post, error) {
	var post model.Post
	if dbc := repo.db.Where("id = ?", id).First(&post); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &post, nil
}

func (repo *repository) Create(post *model.Post) error {
	if dbc := repo.db.Create(&post); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) Update(post *model.Post) error {
	err := repo.db.Where("id = ?", post.Id).Updates(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(post *model.Post) error {
	err := repo.db.Where("id = ?", post.Id).Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}
