package repositories

import "github.com/RomainC75/todo2/data/models"

type ListRepositoryInterface interface {
	CreateList(list models.List) (models.List, error)
}
