package controllers

import (
	"net/http"
	"strconv"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/dto/responses"
	"github.com/RomainC75/todo2/api/services"
	"github.com/gin-gonic/gin"
)

type ListCtrl struct {
	listSrv services.ListSrv
}

func NewListCtrl() *ListCtrl {
	return &ListCtrl{
		listSrv: *services.NewListSrv(),
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
	createdList, err := listCtrl.listSrv.CreateListSrv(userId.(uint), newList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusAccepted, gin.H{"created": createdList})
}

func (listCtrl *ListCtrl) HandleGetListsFromUser(c *gin.Context) {
	userId, _ := c.Get("user_id")

	foundLists := listCtrl.listSrv.GetListsByUserIdSrv(userId.(uint))
	c.JSON(http.StatusAccepted, gin.H{"lists": responses.GetListResponseFromModelList(foundLists)})
}

func (listCtrl *ListCtrl) HandleGetList(c *gin.Context) {
	listId := c.Param("listId")
	listIdInt, err := strconv.Atoi(listId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "list id not valid"})
		return
	}

	userId, _ := c.Get("user_id")

	foundList, err := listCtrl.listSrv.GetListOwnedByUser(userId.(uint), uint(listIdInt))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"list": responses.GetListResponseFromModel(foundList)})
}
