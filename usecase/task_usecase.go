package usecase

import (
	"echo-todo-api/model"
	"echo-todo-api/repository"
)

type ITaskusecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdataTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskRepository(tr repository.ITaskRepository) ITaskusecase {
	return &taskUsecase{tr}
}