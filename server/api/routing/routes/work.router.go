package routes

import (
	"github.com/RomainC75/todo2/api/controllers"
	"github.com/RomainC75/todo2/api/middlewares"
	"github.com/gin-gonic/gin"
)

func WorkRoutes(router *gin.Engine) {
	workController := controllers.NewWorkCtrl()

	workGroup := router.Group("/work")
	{
		workGroup.POST("/", middlewares.IsAuth(false), workController.HandleCreateWork)
	}
}
