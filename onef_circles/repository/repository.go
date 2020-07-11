package repository

import (
	"github.com/hoaxoan/onef-api/onef_circles"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_circles.Repository {
	return &repository{db}
}

// Categories
func (repo *repository) Get(req *model.CircleRequest) ([]model.Circle, error) {
	var circles []model.Circle
	if dbc := repo.db.Order("order asc").Find(&circles); dbc.Error != nil {
		return nil, dbc.Error
	}
	return circles, nil
}

func (repo *repository) Create(circle *model.Circle) error {
	if dbc := repo.db.Create(&circle); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) Update(circle *model.Circle) error {
	err := repo.db.Where("id = ?", circle.Id).Updates(&circle).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(circle *model.Circle) error {
	err := repo.db.Where("id = ?", circle.Id).Delete(&circle).Error
	if err != nil {
		return err
	}
	return nil
}
