package repositories

import "github.com/RomainC75/todo2/data/models"

type ItemRepositoryInterface interface {
	CreateItem(item models.Item, preloadedList models.List) (models.Item, error)
}
