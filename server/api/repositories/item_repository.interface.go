package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
)

type ItemRepositoryInterface interface {
	CreateItem(item db.Item, preloadedList db.List) (db.Item, error)
}
