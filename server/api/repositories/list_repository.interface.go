package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListRepositoryInterface interface {
	CreateList(ctx *gin.Context, arg db.CreateListParams) (db.List, error)
	GetLists(ctx *gin.Context, userId int32) ([]db.List, error)
	// GetListById(listId uint) (db.List, error)
	// UpdateList(userId uint, list requests.UpdateListRequest) (db.List, error)
	// DeleteList(listId uint) (db.List, error)
}
