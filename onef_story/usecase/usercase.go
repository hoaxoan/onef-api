package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/hoaxoan/onef-api/onef_story"
	"github.com/hoaxoan/onef-api/onef_story/service"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"golang.org/x/crypto/bcrypt"
)

type storyUsecase struct {
	Repo         onef_story.Repository
}

func NewUsecase(repo onef_story.Repository) onef_story.Usecase {
	return &storyUsecase{
		Repo:         repo
	}
}

func (ucase *storyUsecase) Get(ctx context.Context, req *model.Story, res *model.StoryResponse) error {
	story, err := ucase.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.Story = story
	return nil
}

func (ucase *storyUsecase) GetAll(ctx context.Context, req *model.StoryRequest, res *model.StoryResponse) error {
	stories, err := ucase.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Stories = stories
	return nil
}

func (ucase *storyUsecase) Create(ctx context.Context, req *model.Story, res *model.StoryResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Story = req
	return nil
}

func (ucase *storyUsecase) Update(ctx context.Context, req *model.Story, res *model.StoryResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Story = req
	return nil
}