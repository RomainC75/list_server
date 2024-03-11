package controllers

import (
	"net/http"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/services"
	"github.com/RomainC75/todo2/utils/controller_utils"
	"github.com/gin-gonic/gin"
)

type ItemCtrl struct {
	itemSrv services.ItemSrv
}

func NewItemCtrl() *ItemCtrl {
	return &ItemCtrl{
		itemSrv: *services.NewItemSrv(),
	}
}

func (itemCtrl *ItemCtrl) HandleCreateItem(c *gin.Context) {
	var newItem requests.CreateItemRequest

	if err := c.ShouldBind(&newItem); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the userID"})
		return
	}

	listId, err := controller_utils.GetIdFromParam(c, "listId")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the list id from the url"})
		return
	}

	createdItem, err := itemCtrl.itemSrv.CreateItemSrv(c, userId.(int32), listId, newItem)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"created": createdItem})
}

func (itemCtrl *ItemCtrl) HandleGetItemsByListId(c *gin.Context) {
	listId, err := controller_utils.GetIdFromParam(c, "listId")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the list id from the url"})
		return
	}
	itemsFound, err := itemCtrl.itemSrv.GetItemsByListSrv(c, listId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"found_items": itemsFound})
}
