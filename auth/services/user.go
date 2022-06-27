package services

import (
	"github.com/dheiro/adam-test/auth/models"
	"gorm.io/gorm"
)

type AuthServiceItf interface {
	Register(data models.User) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{db: db}
}

func (u userService) Register(payload models.User) error {
	var (
		result models.User
	)

	result.Email = payload.Email
	result.Password = payload.Password
	result.Name = payload.Name
	// result.CreatedAt = time.Now()
	// result.UpdatedAt = time.Now()

	return u.db.Create(&result).Error
}
