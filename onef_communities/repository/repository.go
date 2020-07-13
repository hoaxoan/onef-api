package repository

import (
	"github.com/hoaxoan/onef-api/onef_communities"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_communities.Repository {
	return &repository{db}
}

func (repo *repository) Get(req *model.CommunityRequest) ([]model.Community, error) {
	var communities []model.Community
	if dbc := repo.db.Find(&communities); dbc.Error != nil {
		return nil, dbc.Error
	}
	return communities, nil
}

func (repo *repository) Create(community *model.Community) error {
	if dbc := repo.db.Create(&community); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) Update(community *model.Community) error {
	err := repo.db.Where("id = ?", community.Id).Updates(&community).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(community *model.Community) error {
	err := repo.db.Where("id = ?", community.Id).Delete(&community).Error
	if err != nil {
		return err
	}
	return nil
}
