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

func (ucase *usecase) Get(ctx context.Context, req *model.CommunityRequest, res *model.CommunityResponse) error {
	communities, err := ucase.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Communities = communities
	return nil
}

func (ucase *usecase) Create(ctx context.Context, req *model.Community, res *model.CommunityResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Community = req
	return nil
}

func (ucase *usecase) Update(ctx context.Context, req *model.Community, res *model.CommunityResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Community = req
	return nil
}

func (ucase *usecase) Delete(ctx context.Context, req *model.Community, res *model.CommunityResponse) error {
	if err := ucase.Repo.Delete(req); err != nil {
		return err
	}
	res.Community = req
	return nil
}
