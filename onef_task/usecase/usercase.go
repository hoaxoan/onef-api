package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/hoaxoan/onef-api/onef_task"
	"github.com/hoaxoan/onef-api/onef_task/service"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"golang.org/x/crypto/bcrypt"
)

type taskUsecase struct {
	Repo         onef_task.Repository
}

func NewUsecase(repo onef_task.Repository) onef_task.Usecase {
	return &taskUsecase{
		Repo:         repo
	}
}

func (ucase *taskUsecase) Get(ctx context.Context, req *model.Task, res *model.TaskResponse) error {
	task, err := ucase.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.Task = task
	return nil
}

func (ucase *taskUsecase) GetAll(ctx context.Context, req *model.TaskRequest, res *model.TaskResponse) error {
	tasks, err := ucase.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Tasks = tasks
	return nil
}

func (ucase *taskUsecase) Create(ctx context.Context, req *model.Task, res *model.TaskResponse) error {
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.Task = req
	return nil
}

func (ucase *taskUsecase) Update(ctx context.Context, req *model.Task, res *model.TaskResponse) error {
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.Task = req
	return nil
}