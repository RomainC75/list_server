package repositories

import (
	"errors"

	"github.com/RomainC75/todo2/data/database"
	"github.com/RomainC75/todo2/data/models"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepo() *ItemRepository {
	return &ItemRepository{
		DB: database.GetConnection(),
	}
}

func (itemRepo *ItemRepository) CreateItem(item models.Item, preloadedList models.List) (models.Item, error) {

	itemRepo.DB.Preload("Item").First(&preloadedList, 1)
	preloadedList.Items = append(preloadedList.Items, &item)

	if result := itemRepo.DB.Save(&preloadedList); result.RowsAffected == 0 {
		return models.Item{}, errors.New("ekrror trying to create")
	}

	// var newItem models.Item
	// if result := itemRepo.DB.Create(&item).Scan(&newItem); result.RowsAffected == 0 {
	// 	return models.Item{}, errors.New("error trying to create a new user :-(")
	// }

	return item, nil
}
