package controllers

import (
	"fmt"
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/services"
	Manager "github.com/RomainC75/todo2/api/sockets/managers"
	redis_server_handler "github.com/RomainC75/todo2/redis"
	redis_dto "github.com/RomainC75/todo2/redis/dto"
	"github.com/RomainC75/todo2/utils/controller_utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkCtrl struct {
	listSrv services.ListSrv
	manager Manager.ManagerInterface
}

func NewWorkCtrl() *WorkCtrl {
	manager := Manager.New()
	redis_server_handler.GoSubscribe(redis_server_handler.Get(), manager)
	return &WorkCtrl{
		listSrv: *services.NewListSrv(),
		manager: manager,
	}
}

func (workCtrl *WorkCtrl) HandleCreateWork(c *gin.Context) {
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
	c.JSON(http.StatusAccepted, gin.H{"message": "work sent", "uuid": uuid.String()})
}

func (workCtrl *WorkCtrl) HandleGetSocket(c *gin.Context) {
	socketId, err := controller_utils.GetUUIDFromParam(c, "socketId")
	//int32
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	userUuid := uuid.New()
	userData := Manager.UserData{
		UserId:    userUuid,
		RequestId: socketId,
	}

	workCtrl.manager.ServeWS(c.Writer, c.Request, userData)

	fmt.Println("socketId: ", socketId)
}
