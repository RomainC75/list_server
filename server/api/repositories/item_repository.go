package repositories

import "github.com/RomainC75/todo2/data/models"

type ItemRepositoryInterface interface {
	CreateItem(item models.Item) (models.Item, error)
}
