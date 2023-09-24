// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/task_usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	model "echo-todo-api/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITaskusecase is a mock of ITaskusecase interface.
type MockITaskusecase struct {
	ctrl     *gomock.Controller
	recorder *MockITaskusecaseMockRecorder
}

// MockITaskusecaseMockRecorder is the mock recorder for MockITaskusecase.
type MockITaskusecaseMockRecorder struct {
	mock *MockITaskusecase
}

// NewMockITaskusecase creates a new mock instance.
func NewMockITaskusecase(ctrl *gomock.Controller) *MockITaskusecase {
	mock := &MockITaskusecase{ctrl: ctrl}
	mock.recorder = &MockITaskusecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITaskusecase) EXPECT() *MockITaskusecaseMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockITaskusecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", task)
	ret0, _ := ret[0].(model.TaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockITaskusecaseMockRecorder) CreateTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockITaskusecase)(nil).CreateTask), task)
}

// DeleteTask mocks base method.
func (m *MockITaskusecase) DeleteTask(userId, taskId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", userId, taskId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockITaskusecaseMockRecorder) DeleteTask(userId, taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockITaskusecase)(nil).DeleteTask), userId, taskId)
}

// GetAllTasks mocks base method.
func (m *MockITaskusecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks", userId)
	ret0, _ := ret[0].([]model.TaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockITaskusecaseMockRecorder) GetAllTasks(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockITaskusecase)(nil).GetAllTasks), userId)
}

// GetTaskById mocks base method.
func (m *MockITaskusecase) GetTaskById(userId, taskId uint) (model.TaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskById", userId, taskId)
	ret0, _ := ret[0].(model.TaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskById indicates an expected call of GetTaskById.
func (mr *MockITaskusecaseMockRecorder) GetTaskById(userId, taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskById", reflect.TypeOf((*MockITaskusecase)(nil).GetTaskById), userId, taskId)
}

// UpdateTask mocks base method.
func (m *MockITaskusecase) UpdateTask(task model.Task, userId, taskId uint) (model.TaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", task, userId, taskId)
	ret0, _ := ret[0].(model.TaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockITaskusecaseMockRecorder) UpdateTask(task, userId, taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockITaskusecase)(nil).UpdateTask), task, userId, taskId)
}