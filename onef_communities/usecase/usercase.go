package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_communities"
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type usecase struct {
	Repo onef_communities.Repository
}

func NewUsecase(repo onef_communities.Repository) onef_communities.Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (ucase *usecase) Get(ctx context.Context, req *model.DeviceRequest, res *model.DeviceResponse) error {
	devices, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Devices = devices
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Device, res *model.DeviceResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Device = req
	return nil
}

func (ucase *usecase) Update(ctx context.Context, req *model.Device, res *model.DeviceResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Device = req
	return nil
}

func (ucase *usecase) Delete(ctx context.Context, req *model.Device, res *model.DeviceResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Device = req
	return nil
}

func (ucase *usecase) GetDeviceWithUuid(ctx context.Context, deviceUuid string, res *model.DeviceResponse) error {
	device, err := ucase.Repo.GetDeviceWithUuid(deviceUuid)
	if err != nil {
		return err
	}
	res.Device = device
	return nil
}
