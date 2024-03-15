package controllers

import (
	"context"
	"fmt"
	"net/http"

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

type Message struct {
	Message string `json:"message"`
}

func (workCtrl *ListCtrl) HandleCreateWork(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Publish the message(context, message, queue Name)
	pub := redis.GetPublisher()
	pub.PublishMessages(ctx, Message{"mklqjsdfm"}, "myqueue")
	fmt.Println("==> message sent : ")
	c.JSON(http.StatusAccepted, gin.H{"message": "got it"})
}
