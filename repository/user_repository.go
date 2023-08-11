package repository

import "echo-todo-api/model"

// interfaceとはmethodの一覧
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
