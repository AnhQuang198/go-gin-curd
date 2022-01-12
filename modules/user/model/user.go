package model

import (
	"errors"
	"strings"
)

type User struct {
	Name     string `json:"name" gorm:"column:name"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

type UserUpdate struct {
	Name     *string `json:"name" gorm:"column:name"`
	Username *string `json:"username" gorm:"column:username"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserCreate struct {
	Name     string `json:"name" gorm:"column:name"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (user *UserCreate) Validate() error {
	user.Name = strings.TrimSpace(user.Name)

	if len(user.Name) == 0 {
		return errors.New("name can't be blank")
	}

	return nil
}
