package repositories

import (
	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/data/models"
)

type ListRepositoryInterface interface {
	CreateList(list models.List) (models.List, error)
	GetLists(userId uint) []models.List
	GetListById(listId uint) (models.List, error)
	UpdateList(userId uint, list requests.UpdateListRequest) (models.List, error)
	DeleteList(listId uint) (models.List, error)
}
