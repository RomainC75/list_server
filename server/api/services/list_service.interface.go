package services

import (
	"github.com/RomainC75/todo2/api/dto/requests"
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListSrvInterface interface {
	CreateListSrv(ctx *gin.Context, userId int32, list requests.CreateListReq) (db.List, error)
	GetListsByUserIdSrv(userId uint) []db.List
	GetListOwnedByUser(userId uint, listId uint) (db.List, error)
	UpdateList(userId uint, list db.List) (db.List, error)
}
