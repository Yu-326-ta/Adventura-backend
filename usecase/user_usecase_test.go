package usecase

import (
	"echo-todo-api/model"
	mock_usecases "echo-todo-api/usecase/mock_usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_userUsecase_SignUp(t *testing.T) {
	// モックを呼び出すための Controller を生成
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	// モックの生成
    mockApi := mock_usecases.NewMockIUserUsecase(ctrl)
	user := model.User{
		ID: 1,
		Email: "yuta@example.com",
		Password: "yuta326",
	}
	expected := model.UserResponse{ID: 1, Email: "yuta@example.com"}
	// テストに呼ばれるべきメソッドと引数・戻り値を指定
	mockApi.EXPECT().SignUp(user).Return(model.UserResponse{ID: 1, Email: "yuta@example.com"}, nil)

	res, err := mockApi.SignUp(user)
	if err != nil {
		t.Errorf("error! %v", err)
	}
	if res != expected {
		t.Errorf("SignUp() = %v want %v", res, expected)
	}
}

func Test_userUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockApi := mock_usecases.NewMockIUserUsecase(ctrl)
	user := model.User{ID: 1, Email: "yuta@example.com", Password: "yuta326"}
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU1MDQxOTcsInVzZXJfaWQiOjF9.EjeP9ivCGRDOy4RVI6g711kjiDCooz6rv1eCfE1bp70"
	mockApi.EXPECT().Login(user).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU1MDQxOTcsInVzZXJfaWQiOjF9.EjeP9ivCGRDOy4RVI6g711kjiDCooz6rv1eCfE1bp70", nil)

	res, err := mockApi.Login(user)
	if err != nil {
		t.Errorf("error! %v", err)
	}
	if res != expected {
		t.Errorf("Login() = %v want %v", res, expected)
	}
}
