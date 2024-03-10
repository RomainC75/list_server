package routes

import (
	"fmt"
	"net/http"

	"github.com/RomainC75/todo2/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(middlewares.CORSMiddleware())
	AuthRoutes(router)
	// ListRoutes(router)

	router.GET("/hello", func(c *gin.Context) {
		fmt.Println("received")
		c.JSON(http.StatusAccepted, gin.H{"message": "hello ! "})
	})

}
