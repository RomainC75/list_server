package controllers

import (
	"fmt"
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/services"
	redis_server_handler "github.com/RomainC75/todo2/redis"
	redis_dto "github.com/RomainC75/todo2/redis/dto"
	"github.com/RomainC75/todo2/utils/controller_utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	uuid := uuid.New()
	workRequest := redis_dto.ScanRequestMessage{
		Id:      uuid,
		Address: scanRequest.Address,
		PortMin: scanRequest.PortMin,
		PortMax: scanRequest.PortMax,
	}
	// Publish the message(context, message, queue Name)
	pub := redis_server_handler.GetJobQueue()
	pub.PublishMessage(workRequest, "myqueue")
	fmt.Println("==> message sent : ")
	c.JSON(http.StatusAccepted, gin.H{"message": "work sent", "uuid": uuid.String()})
}

func (workCtrl *ListCtrl) HandleGetSocket(c *gin.Context) {
	socketId, err := controller_utils.GetUUIDFromParam(c, "socketId")
	//int32
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("socketId: ", socketId)
}
