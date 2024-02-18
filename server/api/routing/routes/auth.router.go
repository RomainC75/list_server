package routes

import (
	"github.com/RomainC75/todo2/api/controllers"
	"github.com/RomainC75/todo2/api/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authController := controllers.NewAuthCtrl()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authController.HandleSignup)
		authGroup.POST("/signin", authController.HandleLogin)
		authGroup.GET("/verify", middlewares.IsAuth(false), authController.HandleVerify)
	}
}
