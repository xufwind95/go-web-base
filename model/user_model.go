package model

import (
	"github.com/jinzhu/gorm"
	"github.com/xufwind95/go-web-base/pkg/database"
)

type UserTestModel struct {
	gorm.Model
	Username string `json:"username" gorm:"not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserTestModel) TableName() string {
	return "tb_users"
}

// 创建用户
func CreateUser(user *UserTestModel) error {
	db := database.DB
	return db.Create(&user).Error
}

// 根据条件查找一个用户: userModel, err := FindOneUser(&UserModel{Username: "username0"})
func FindOneUser(condition interface{}) (UserTestModel, error) {
	db := database.DB
	var user UserTestModel
	err := db.Where(condition).First(&user).Error
	return user, err
}
