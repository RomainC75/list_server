package controllers

import (
	"github.com/gin-gonic/gin"
)

type ItemCtrl struct {
	// listSrv services.ListSrv
	// itemSrv services.ItemSrv
}

func NewItemCtrl() *ItemCtrl {
	return &ItemCtrl{
		// listSrv: *services.NewListSrv(),
		// itemSrv: *services.NewItemSrv(),
	}
}

func (itemCtrl *ItemCtrl) HandleCreateItem(c *gin.Context) {
	// var newList requests.ItemToCreateRequest

	// if err := c.ShouldBind(&newList); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	// userId, exists := c.Get("user_id")

	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the userID"})
	// }

	// foundList, err := itemCtrl.itemSrv.CreateItemSrv()

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// }
	// c.JSON(http.StatusAccepted, gin.H{"created": createdList})
}
