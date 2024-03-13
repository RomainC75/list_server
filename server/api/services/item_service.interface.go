package services

import (
	"github.com/RomainC75/todo2/api/dto/requests"
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ItemSrvInterface interface {
	CreateItemSrv(ctx *gin.Context, userId uint, listId int32, item requests.CreateItemRequest) (db.Item, error)
	GetItemsByListSrv(ctx *gin.Context, listId int32) ([]db.Item, error)
	UpdateItem(ctx *gin.Context, itemId int32, itemRequest requests.UpdateItemRequest) (db.Item, error)
}
