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

func (repo *repository) Get(req *model.CircleRequest) ([]model.Circle, error) {
	var circles []model.Circle
	if dbc := repo.db.Find(&circles); dbc.Error != nil {
		return nil, dbc.Error
	}

	circleList := make([]model.Circle, 0)
	for _, circle := range circles {
		var creator model.User
		if dbc := repo.db.Where("id = ?", circle.CreatorId).First(&creator); dbc.Error == nil {
			circle.Creator = &creator
		}
		circleList = append(circleList, circle)
	}

	return circleList, nil

}

func (repo *repository) GetCircleWithId(circleId int64) (*model.Circle, error) {
	var circle model.Circle
	if dbc := repo.db.Where("id = ?", circleId).First(&circle); dbc.Error != nil {
		return nil, dbc.Error
	}

	var creator model.User
	if dbc := repo.db.Where("id = ?", circle.CreatorId).First(&creator); dbc.Error == nil {
		circle.Creator = &creator
	}

	return &circle, nil
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

func (repo *repository) CheckNameIsAvailable(req *model.Circle) (*model.Circle, error) {
	var circle model.Circle
	if dbc := repo.db.Where("name = ? AND creator_id = ?", req.Name, req.CreatorId).First(&circle); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &circle, nil
}
