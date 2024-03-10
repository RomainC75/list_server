package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
)

type ListRepositoryInterface interface {
	CreateList(list db.List) (db.List, error)
	GetLists(userId uint) []db.List
	GetListById(listId uint) (db.List, error)
	// UpdateList(userId uint, list requests.UpdateListRequest) (db.List, error)
	DeleteList(listId uint) (db.List, error)
}
