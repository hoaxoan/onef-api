package onef_devices

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.DeviceRequest, res *model.DeviceResponse) error
	Create(ctx context.Context, req *model.Device, res *model.DeviceResponse) error
	Update(ctx context.Context, req *model.Device, res *model.DeviceResponse) error
	Delete(ctx context.Context, req *model.Device, res *model.DeviceResponse) error

	GetDeviceWithUuid(ctx context.Context, deviceUuid string, res *model.DeviceResponse) error
}
