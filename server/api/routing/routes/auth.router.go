package routes

import (
	"github.com/RomainC75/todo2/api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authController := controllers.NewAuthCtrl()

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/signup", authController.HandleSignup)
	}
}
