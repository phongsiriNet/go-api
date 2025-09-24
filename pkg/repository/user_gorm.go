package repository

import (
	"fmt"
	"go-api/pkg/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUser {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(user *model.User) (*model.User, error) {
	result := u.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(result.Statement.SQL.String())
	return user, nil
}

func (u *userRepository) GetUser(username string) (*model.User, error) {
	user := model.User{}
	result := u.db.Where("username =?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(result.Statement.SQL.String())
	return &user, nil
}
