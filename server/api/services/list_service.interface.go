package services

import (
	db "github.com/RomainC75/todo2/db/sqlc"
)

type ListSrvInterface interface {
	CreateListSrv(userId uint, list db.List) (db.List, error)
	GetListsByUserIdSrv(userId uint) []db.List
	GetListOwnedByUser(userId uint, listId uint) (db.List, error)
	UpdateList(userId uint, list db.List) (db.List, error)
}
