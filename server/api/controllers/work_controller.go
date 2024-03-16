package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/services"
	"github.com/RomainC75/todo2/redis"
	"github.com/gin-gonic/gin"
)

type WorkCtrl struct {
	listSrv services.ListSrv
}

func NewWorkCtrl() *ListCtrl {
	return &ListCtrl{
		listSrv: *services.NewListSrv(),
	}
}

func (workCtrl *ListCtrl) HandleCreateWork(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var scanRequest requests.WorkRequest
	if err := c.ShouldBindJSON(&scanRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if scanRequest.PortMin > scanRequest.PortMax {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "port_min > port_max"})
		return
	}

	// Publish the message(context, message, queue Name)
	pub := redis.GetPublisher()
	pub.PublishMessages(ctx, scanRequest, "tcpScan")
	fmt.Println("==> message sent : ")
	c.JSON(http.StatusAccepted, gin.H{"message": "got it"})
}
