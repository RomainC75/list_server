package controllers

import (
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/services"
	"github.com/gin-gonic/gin"
)

type ListCtrl struct {
	userSrv services.ListSrv
}

func NewListCtrl() *ListCtrl {
	return &ListCtrl{
		userSrv: *services.NewListSrv(),
	}
}

func (listCtrl *ListCtrl) HandleCreateList(c *gin.Context) {
	var newList requests.CreateListRequest

	if err := c.ShouldBind(&newList); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the userID"})
	}
	createdList, err := listCtrl.userSrv.CreateListSrv(userId.(uint), newList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusAccepted, gin.H{"created": createdList})
}
