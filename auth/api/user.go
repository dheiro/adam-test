package api

import (
	"github.com/dheiro/adam-test/auth/models"
	"github.com/dheiro/adam-test/auth/services"
	"github.com/gin-gonic/gin"
)

type AuthItf interface {
	Register(*gin.Context)
	// Login(user *models.User) error
	// GetUser(user *models.User) error
}

type AuthHandler struct {
	service services.AuthServiceItf
}

func NewAuthHandler(svc services.AuthServiceItf) AuthItf {
	return AuthHandler{service: svc}
}

func (s AuthHandler) Register(c *gin.Context) {
	var payload models.User

	if c.BindJSON(&payload) != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := s.service.Register(payload)
	if err != nil {
		c.JSON(500, gin.H{
			"message": string(err.Error()),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "User created",
	})
}
