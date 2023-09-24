package repository

import (
	"echo-todo-api/model"
	"echo-todo-api/repository/mock_repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_userRipository_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_repository.NewMockIUserRepository(ctrl)
	user := &model.User{ID: 1, Email: "yuta@example.com", Password: "yuta326"}
	mockApi.EXPECT().GetUserByEmail(user, user.Email).Return(nil)

	err := mockApi.GetUserByEmail(user, user.Email)
	if err != nil {
		t.Errorf("error! %v", err)
	}
}

func Test_userRipository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_repository.NewMockIUserRepository(ctrl)
	user := &model.User{ID: 1, Email: "yuta@example.com", Password: "yuta326"}
	mockApi.EXPECT().CreateUser(user).Return(nil)

	err := mockApi.CreateUser(user)
	if err != nil {
		t.Errorf("error! %v", err)
	}
}
