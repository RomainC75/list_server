package controllers

import (
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
	// var newList requests.CreateListRequest

	// if err := c.ShouldBind(&newList); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	// userId, exists := c.Get("user_id")
	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the userID"})
	// }
	// createdList, err := listCtrl.listSrv.CreateListSrv(userId.(uint), newList)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// }
	// c.JSON(http.StatusAccepted, gin.H{"created": createdList})
}

func (listCtrl *ListCtrl) HandleGetListsFromUser(c *gin.Context) {
	// userId, _ := c.Get("user_id")

	// foundLists := listCtrl.listSrv.GetListsByUserIdSrv(userId.(uint))
	// c.JSON(http.StatusAccepted, gin.H{"lists": responses.GetListResponseFromModelList(foundLists)})
}

func (listCtrl *ListCtrl) HandleGetList(c *gin.Context) {
	// listId := c.Param("listId")
	// listIdInt, err := strconv.Atoi(listId)
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "list id not valid"})
	// 	return
	// }

	// userId, _ := c.Get("user_id")

	// foundList, err := listCtrl.listSrv.GetListOwnedByUser(userId.(uint), uint(listIdInt))
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusAccepted, gin.H{"list": responses.GetListResponseFromModel(foundList)})
}

func (listCtrl *ListCtrl) HandleUpdateList(c *gin.Context) {
	// listId, err := controller_utils.GetIdFromParam(c, "listId")
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	// userId, _ := c.Get("user_id")

	// var list requests.UpdateListRequest
	// if err := c.ShouldBind(&list); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	// list.Id = uint(listId)

	// updatedList, err := listCtrl.listSrv.UpdateList(userId.(uint), list)
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusAccepted, gin.H{"updated_list": responses.GetListResponseFromModel(updatedList)})
}

func (listCtrl *ListCtrl) HandleDeleteList(c *gin.Context) {
	// listId, err := controller_utils.GetIdFromParam(c, "listId")
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	// userId, _ := c.Get("user_id")
	// deletedList, err := listCtrl.listSrv.DeleteList(userId.(uint), uint(listId))
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusUnauthorized, gin.H{"deleted": responses.GetListResponseFromModel(deletedList)})

}
