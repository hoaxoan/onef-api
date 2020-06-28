package usecase

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_task"
)

type taskUsecase struct {
	Repo onef_task.Repository
}

func NewUsecase(repo onef_task.Repository) onef_task.Usecase {
	return &taskUsecase{
		Repo: repo,
	}
}

func (ucase *taskUsecase) GetAll(ctx context.Context, req *model.TaskRequest, res *model.TaskResponse) error {
	tasks, err := ucase.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Tasks = tasks
	return nil
}

func (ucase *taskUsecase) Get(ctx context.Context, id string, res *model.TaskResponse) error {
	task, err := ucase.Repo.Get(id)
	if err != nil {
		return err
	}
	res.Task = task
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
