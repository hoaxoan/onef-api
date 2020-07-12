package onef_devices

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Repository interface {
	Get(req *model.DeviceRequest) ([]model.Device, error)
	Create(category *model.Device) error
	Update(category *model.Device) error
	Delete(category *model.Device) error

	GetDeviceWithUuid(deviceUuid string) (*model.Device, error)
}
