package services

import (
	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/data/models"
)

type ListSrvInterface interface {
	CreateListSrv(userId uint, list requests.CreateListRequest) (models.List, error)
	GetListsByUserIdSrv(userId uint) []models.List
	GetListOwnedByUser(userId uint, listId uint) (models.List, error)
	UpdateList(userId uint, list requests.UpdateListRequest) (models.List, error)
}
