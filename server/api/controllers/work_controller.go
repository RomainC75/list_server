package controllers

import (
	"fmt"
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/services"
	"github.com/RomainC75/todo2/redis"
	redis_dto "github.com/RomainC75/todo2/redis/dto"
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
	fmt.Println("mlqksjdmlqkjfdmlqskjfmlqksdjf")
	var scanRequest requests.WorkRequest
	if err := c.ShouldBindJSON(&scanRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if scanRequest.PortMin > scanRequest.PortMax {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "port_min > port_max"})
		return
	}
	workRequest := redis_dto.ScanRequestMessage{
		Address: scanRequest.Address,
		PortMin: scanRequest.PortMin,
		PortMax: scanRequest.PortMax,
	}
	// Publish the message(context, message, queue Name)
	pub := redis.GetJobQueue()
	pub.PublishMessage(workRequest, "myqueue")
	fmt.Println("==> message sent : ")
	c.JSON(http.StatusAccepted, gin.H{"message": "got it"})
}
