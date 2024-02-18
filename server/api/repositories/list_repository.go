package repositories

import (
	"errors"

	"github.com/RomainC75/todo2/data/database"
	"github.com/RomainC75/todo2/data/models"
	"gorm.io/gorm"
)

type ListRepository struct {
	DB *gorm.DB
}

func NewListRepo() *ListRepository {
	return &ListRepository{
		DB: database.GetConnection(),
	}
}

func (listRepo *ListRepository) CreateList(list models.List) (models.List, error) {
	var newList models.List
	if result := listRepo.DB.Create(&list).Scan(&newList); result.RowsAffected == 0 {
		return models.List{}, errors.New("error trying to create a new user :-(")
	}
	return newList, nil
}
