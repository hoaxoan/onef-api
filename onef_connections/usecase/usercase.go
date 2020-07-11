package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_connections"
	"github.com/hoaxoan/onef-api/onef_core/model"
)

type usecase struct {
	Repo onef_connections.Repository
}

func NewUsecase(repo onef_connections.Repository) onef_connections.Usecase {
	return &usecase{
		Repo: repo,
	}
}

func (ucase *usecase) Get(ctx context.Context, req *model.ConnectionRequest, res *model.ConnectionResponse) error {
	connections, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Connections = connections
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Connection, res *model.ConnectionResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Connection = req
	return nil
}

func (ucase *usecase) Update(ctx context.Context, req *model.Connection, res *model.ConnectionResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Connection = req
	return nil
}

func (ucase *usecase) Delete(ctx context.Context, req *model.Connection, res *model.ConnectionResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Connection = req
	return nil
}
