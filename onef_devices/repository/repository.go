package repository

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_devices"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) onef_devices.Repository {
	return &repository{db}
}

func (repo *repository) Get(req *model.DeviceRequest) ([]model.Device, error) {
	var devices []model.Device
	if dbc := repo.db.Find(&devices); dbc.Error != nil {
		return nil, dbc.Error
	}
	return devices, nil
}

func (repo *repository) Create(device *model.Device) error {
	if dbc := repo.db.Create(&device); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) Update(device *model.Device) error {
	err := repo.db.Where("id = ?", device.Id).Updates(&device).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(device *model.Device) error {
	err := repo.db.Where("id = ?", device.Id).Delete(&device).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetDeviceWithUuid(deviceUuid string) (*model.Device, error) {
	var device model.Device
	if dbc := repo.db.Where("uuid = ?", deviceUuid).First(device); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &device, nil
}
