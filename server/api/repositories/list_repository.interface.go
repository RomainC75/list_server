package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListRepositoryInterface interface {
	CreateList(ctx *gin.Context, arg db.CreateListParams) (db.List, error)
	GetLists(ctx *gin.Context, userId int32) ([]db.List, error)
	GetListByIdByOwner(ctx *gin.Context, listToGet db.GetlistParams) (db.List, error)
	UpdateList(ctx *gin.Context, listToGet db.GetListForUpdateParams, listToUpdate db.UpdateListParams) (db.List, error)
	DeleteList(ctx *gin.Context, listToDelete db.DeleteListParams) (db.List, error)
	// GetListById(listId uint) (db.List, error)
	// UpdateList(userId uint, list requests.UpdateListRequest) (db.List, error)
}
