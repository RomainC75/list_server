package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
)

type ItemRepository struct {
	DB *db.Store
}

func NewItemRepo() *ItemRepository {
	return &ItemRepository{
		DB: db.GetConnection(),
	}
}

// func (itemRepo *ItemRepository) CreateItem(item models.Item, preloadedList models.List) (models.Item, error) {

// 	// itemRepo.DB.Preload("Item").First(&preloadedList, 1)
// 	// preloadedList.Items = append(preloadedList.Items, &item)

// 	// if result := itemRepo.DB.Save(&preloadedList); result.RowsAffected == 0 {
// 	// 	return models.Item{}, errors.New("ekrror trying to create")
// 	// }

// 	// var newItem models.Item
// 	// if result := itemRepo.DB.Create(&item).Scan(&newItem); result.RowsAffected == 0 {
// 	// 	return models.Item{}, errors.New("error trying to create a new user :-(")
// 	// }

// 	return item, nil
// }
