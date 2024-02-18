package routes

import (
	"github.com/RomainC75/todo2/api/controllers"
	"github.com/RomainC75/todo2/api/middlewares"
	"github.com/gin-gonic/gin"
)

func ListRoutes(router *gin.Engine) {
	listController := controllers.NewListCtrl()

	listGroup := router.Group("/list")
	{
		listGroup.POST("/", middlewares.IsAuth(false), listController.HandleCreateList)
		listGroup.GET("/", middlewares.IsAuth(false), listController.HandleGetListsFromUser)
		listGroup.GET("/:listId", middlewares.IsAuth(false), listController.HandleGetList)
		listGroup.PUT("/:listId", middlewares.IsAuth(false), listController.HandleUpdateList)
		listGroup.DELETE("/:listId", middlewares.IsAuth(false), listController.HandleDeleteList)
	}
}
