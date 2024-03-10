package routes

import (
	"github.com/RomainC75/todo2/api/controllers"
	"github.com/RomainC75/todo2/api/middlewares"
	"github.com/gin-gonic/gin"
)

func ItemRoutes(router *gin.Engine) {
	itemController := controllers.NewItemCtrl()

	itemGroup := router.Group("/items")
	{
		itemGroup.POST("/:listId", middlewares.IsAuth(false), itemController.HandleCreateItem)
		// itemGroup.GET("/", middlewares.IsAuth(false), itemController.HandleGetListsFromUser)
		// itemGroup.GET("/:listId", middlewares.IsAuth(false), itemController.HandleGetList)
		// itemGroup.PUT("/:listId", middlewares.IsAuth(false), itemController.HandleUpdateList)
		// itemGroup.DELETE("/:listId", middlewares.IsAuth(false), itemController.HandleDeleteList)
	}
}
