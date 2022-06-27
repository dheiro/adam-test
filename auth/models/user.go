package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"not null;unique" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Name     string `json:"name" gorm:"not null" binding:"required"`
}

func (User) TableName() string {
	return "users"
}
