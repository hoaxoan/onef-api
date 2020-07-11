package repository

import (
	"github.com/hoaxoan/onef-api/onef_categories"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_categories.Repository {
	return &repository{db}
}

func (repo *repository) Get(req *model.CategoryRequest) ([]model.Category, error) {
	var categories []model.Category
	if dbc := repo.db.Order("order asc").Find(&categories); dbc.Error != nil {
		return nil, dbc.Error
	}
	return categories, nil
}

func (repo *repository) Create(category *model.Category) error {
	if dbc := repo.db.Create(&category); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) Update(category *model.Category) error {
	err := repo.db.Where("id = ?", category.Id).Updates(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(category *model.Category) error {
	err := repo.db.Where("id = ?", category.Id).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
