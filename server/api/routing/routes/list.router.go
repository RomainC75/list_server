package routes

import (
	"github.com/RomainC75/todo2/api/controllers"
	"github.com/RomainC75/todo2/api/middlewares"
	"github.com/gin-gonic/gin"
)

func ListRoutes(router *gin.Engine) {
	listController := controllers.NewListCtrl()

	authGroup := router.Group("/list")
	{
		authGroup.POST("/", middlewares.IsAuth(false), listController.HandleCreateList)
	}
}
