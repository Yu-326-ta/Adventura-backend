// Code generated by MockGen. DO NOT EDIT.
// Source: repository/task_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "echo-todo-api/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITaskRepository is a mock of ITaskRepository interface.
type MockITaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITaskRepositoryMockRecorder
}

// MockITaskRepositoryMockRecorder is the mock recorder for MockITaskRepository.
type MockITaskRepositoryMockRecorder struct {
	mock *MockITaskRepository
}

// NewMockITaskRepository creates a new mock instance.
func NewMockITaskRepository(ctrl *gomock.Controller) *MockITaskRepository {
	mock := &MockITaskRepository{ctrl: ctrl}
	mock.recorder = &MockITaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITaskRepository) EXPECT() *MockITaskRepositoryMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockITaskRepository) CreateTask(task *model.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockITaskRepositoryMockRecorder) CreateTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockITaskRepository)(nil).CreateTask), task)
}

// DeleteTask mocks base method.
func (m *MockITaskRepository) DeleteTask(userId, taskId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", userId, taskId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockITaskRepositoryMockRecorder) DeleteTask(userId, taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockITaskRepository)(nil).DeleteTask), userId, taskId)
}

// GetAllTasks mocks base method.
func (m *MockITaskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks", tasks, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockITaskRepositoryMockRecorder) GetAllTasks(tasks, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockITaskRepository)(nil).GetAllTasks), tasks, userId)
}

// GetTaskById mocks base method.
func (m *MockITaskRepository) GetTaskById(task *model.Task, userId, taskId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskById", task, userId, taskId)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTaskById indicates an expected call of GetTaskById.
func (mr *MockITaskRepositoryMockRecorder) GetTaskById(task, userId, taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskById", reflect.TypeOf((*MockITaskRepository)(nil).GetTaskById), task, userId, taskId)
}

// UpdateTask mocks base method.
func (m *MockITaskRepository) UpdateTask(task *model.Task, userId, taskId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", task, userId, taskId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockITaskRepositoryMockRecorder) UpdateTask(task, userId, taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockITaskRepository)(nil).UpdateTask), task, userId, taskId)
}
