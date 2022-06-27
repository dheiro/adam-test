package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/dheiro/adam-test/auth/api"
	"github.com/dheiro/adam-test/auth/services"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth Service API",
		})
	})

	authService := services.NewUserService(db)
	authController := api.NewAuthHandler(authService)

	r.POST("register", authController.Register)

	return r
}
