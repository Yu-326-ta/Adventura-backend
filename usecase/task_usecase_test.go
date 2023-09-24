package usecase

import (
	"echo-todo-api/model"
	mock_usecases "echo-todo-api/usecase/mock_usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_taskUsecase_GetAllTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_usecases.NewMockITaskusecase(ctrl)
	user := model.User{
		ID: 1,
		Email: "yuta@example.com",
		Password: "yuta326",
	}
	userId := user.ID
	tasks := []model.TaskResponse{}
	task1 := model.TaskResponse{ID: 1, Title: "programing"}
	task2 := model.TaskResponse{ID: 2, Title: "宿題"}
	task3 := model.TaskResponse{ID: 3, Title: "掃除"}
	tasks = append(tasks, task1, task2, task3)
	expected := tasks
	mockApi.EXPECT().GetAllTasks(userId).Return(tasks, nil)

	res, err := mockApi.GetAllTasks(userId)
	if err != nil {
		t.Errorf("error! %v", err)
	}
	for i, v := range tasks {
		if v != expected[i] {
			t.Errorf("GetAllTasks() = %v want %v", res, expected)
		}
	}
	
}

func Test_taskUsecase_GetTaskById(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_usecases.NewMockITaskusecase(ctrl)
	user := model.User{
		ID: 1,
		Email: "yuta@example.com",
		Password: "yuta326",
	}
	task := model.Task{
		ID: 1,
		Title: "programing",
		User: user,
		UserId: 1,
	}
	userId := user.ID
	taskId := task.ID
	expected := model.TaskResponse{ID: 1, Title: "programing"}
	mockApi.EXPECT().GetTaskById(userId, taskId).Return(model.TaskResponse{ID: 1, Title: "programing"}, nil)

	res, err := mockApi.GetTaskById(userId, taskId)
	if err != nil {
		t.Errorf("error! %v", err)
	}
	if res != expected {
		t.Errorf("GetTaskById() = %v want %v", res, expected)
	}
}

func Test_taskUsecase_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_usecases.NewMockITaskusecase(ctrl)
	user := model.User{
		ID: 1,
		Email: "yuta@example.com",
		Password: "yuta326",
	}
	task := model.Task{
		ID: 1,
		Title: "programing",
		User: user,
		UserId: 1,
	}
	expected := model.TaskResponse{ID: 1, Title: "programing"}
	mockApi.EXPECT().CreateTask(task).Return(model.TaskResponse{ID: 1, Title: "programing"}, nil)

	res, err := mockApi.CreateTask(task)
	if err != nil {
		t.Errorf("error! %v", err)
	}
	if res != expected {
		t.Errorf("CreateTask() = %v want %v", res, expected)
	}
}

func Test_taskUsecase_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_usecases.NewMockITaskusecase(ctrl)
	user := model.User{
		ID: 1,
		Email: "yuta@example.com",
		Password: "yuta326",
	}
	task := model.Task{
		ID: 1,
		Title: "programing",
		User: user,
		UserId: 1,
	}
	userId := user.ID
	taskId := task.ID
	expected := model.TaskResponse{ID: 1, Title: "new_programing"}
	mockApi.EXPECT().UpdateTask(task, userId, taskId).Return(model.TaskResponse{ID: 1, Title: "new_programing"}, nil)

	res, err := mockApi.UpdateTask(task, userId, taskId)
	if err != nil {
		t.Errorf("error! %v", err)
	}
	if res != expected {
		t.Errorf("UpdateTask() = %v want %v", res, expected)
	}
}

func Test_taskUsecase_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_usecases.NewMockITaskusecase(ctrl)
	user := model.User{
		ID: 1,
		Email: "yuta@example.com",
		Password: "yuta326",
	}
	task := model.Task{
		ID: 1,
		Title: "programing",
		User: user,
		UserId: 1,
	}
	userId := user.ID
	taskId := task.ID
	mockApi.EXPECT().DeleteTask(userId, taskId).Return(nil)

	err := mockApi.DeleteTask(userId, taskId)
	if err != nil {
		t.Errorf("error! %v", err)
	}
}
