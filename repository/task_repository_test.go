package repository

import (
	"echo-todo-api/model"
	"echo-todo-api/repository/mock_repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_taskRepository_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_repository.NewMockITaskRepository(ctrl)
	task := &model.Task{ID: 1, Title: "programig"}
	mockApi.EXPECT().CreateTask(task).Return(nil)

	err := mockApi.CreateTask(task)
	if err != nil {
		t.Errorf("error! %v", err)
	}
}

func Test_taskRepository_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_repository.NewMockITaskRepository(ctrl)
	user := model.User{ID: 1, Email: "yuta@example.com", Password: "yuta326"}
	task := &model.Task{ID: 1, Title: "programig"}
	mockApi.EXPECT().UpdateTask(task, user.ID, task.ID).Return(nil)

	err := mockApi.UpdateTask(task, user.ID, task.ID)
	if err != nil {
		t.Errorf("error! %v", err)
	}
}

func Test_taskRepository_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_repository.NewMockITaskRepository(ctrl)
	user := model.User{ID: 1, Email: "yuta@example.com", Password: "yuta326"}
	task := &model.Task{ID: 1, Title: "programig"}
	mockApi.EXPECT().DeleteTask(user.ID, task.ID).Return(nil)

	err := mockApi.DeleteTask(user.ID, task.ID)
	if err != nil {
		t.Errorf("error! %v", err)
	}
}
