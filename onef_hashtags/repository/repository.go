package repository

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_hashtags"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_hashtags.Repository {
	return &repository{db}
}

// Categories
func (repo *repository) Get(req *model.HashtagRequest) ([]model.Hashtag, error) {
	var hashtags []model.Hashtag
	if dbc := repo.db.Find(&hashtags); dbc.Error != nil {
		return nil, dbc.Error
	}
	return hashtags, nil
}

func (repo *repository) Create(hashtag *model.Hashtag) error {
	if dbc := repo.db.Create(&hashtag); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) Update(hashtag *model.Hashtag) error {
	err := repo.db.Where("id = ?", hashtag.Id).Updates(&hashtag).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(hashtag *model.Hashtag) error {
	err := repo.db.Where("id = ?", hashtag.Id).Delete(&hashtag).Error
	if err != nil {
		return err
	}
	return nil
}
