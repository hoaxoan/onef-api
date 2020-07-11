package repository

import (
	"github.com/hoaxoan/onef-api/onef_connections"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_connections.Repository {
	return &repository{db}
}

// Categories
func (repo *repository) Get(req *model.ConnectionRequest) ([]model.Connection, error) {
	var connections []model.Connection
	if dbc := repo.db.Find(&connections); dbc.Error != nil {
		return nil, dbc.Error
	}
	return connections, nil
}

func (repo *repository) Create(connection *model.Connection) error {
	if dbc := repo.db.Create(&connection); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) Update(connection *model.Connection) error {
	err := repo.db.Where("id = ?", connection.Id).Updates(&connection).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(connection *model.Connection) error {
	err := repo.db.Where("id = ?", connection.Id).Delete(&connection).Error
	if err != nil {
		return err
	}
	return nil
}
