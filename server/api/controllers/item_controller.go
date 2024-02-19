package controllers

import (
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/services"
	"github.com/gin-gonic/gin"
)

type ItemCtrl struct {
	listSrv services.ListSrv
}

func NewItemCtrl() *ItemCtrl {
	return &ItemCtrl{
		listSrv: *services.NewListSrv(),
	}
}

func (itemCtrl *ItemCtrl) HandleCreateItem(c *gin.Context) {
	var newList requests.CreateListRequest

	if err := c.ShouldBind(&newList); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the userID"})
	}
	createdList, err := itemCtrl.listSrv.CreateListSrv(userId.(uint), newList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusAccepted, gin.H{"created": createdList})
}
