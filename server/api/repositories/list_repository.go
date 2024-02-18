package repositories

import (
	"errors"
	"fmt"

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

func (listRepo *ListRepository) GetLists(userId uint) []models.List {
	var foundLists []models.List
	listRepo.DB.Where("user_refer = ?", userId).Find(&foundLists)
	return foundLists
}

func (listRepo *ListRepository) GetListById(listId uint) (models.List, error) {
	var foundList models.List
	fmt.Println("listID : ", listId)
	if result := listRepo.DB.Where("id = ? ", listId).First(&foundList); result.RowsAffected == 0 {
		return models.List{}, errors.New("not found")
	}
	return foundList, nil
}
