package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ItemRepositoryInterface interface {
	CreateItem(ctx *gin.Context, itemToCreate db.CreateItemParams, listId int32) (db.Item, error)
}
