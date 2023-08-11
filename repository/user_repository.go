package repository

import (
	"echo-todo-api/model"

	"gorm.io/gorm"
)

// interfaceとはmethodの一覧
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRipository struct {
	db *gorm.DB
}

func NewUserRipository(db *gorm.DB) IUserRepository {
	return &userRipository{db}
}

func (ur *userRipository) GetUserByEmail(user *model.User, email string) error {
	// 検索したユーザーが存在する場合、ポインタの指し示す先の値を検索したユーザーに変更
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRipository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
