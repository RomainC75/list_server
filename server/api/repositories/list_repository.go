package repositories

import (
	"errors"
	"strings"
	"time"

	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListRepository struct {
	Store *db.Store
}

func NewListRepo() *ListRepository {
	return &ListRepository{
		Store: db.GetConnection(),
	}
}

func (listRepo *ListRepository) CreateList(ctx *gin.Context, arg db.CreateListParams) (db.List, error) {
	arg.CreatedAt = time.Now()
	arg.UpdatedAt = arg.CreatedAt
	newList, err := (*listRepo.Store).CreateList(ctx, arg)
	if err != nil {
		return db.List{}, err
	}
	return newList, nil
}

func (listRepo *ListRepository) GetLists(ctx *gin.Context, userId int32) ([]db.List, error) {
	foundLists, err := (*listRepo.Store).Getlists(ctx, userId)
	return foundLists, err
}

// func (listRepo *ListRepository) GetListById(listId uint) (models.List, error) {
// 	var foundList models.List
// 	fmt.Println("listID : ", listId)
// 	// if result := listRepo.DB.Where("id = ? ", listId).First(&foundList); result.RowsAffected == 0 {
// 	// 	return models.List{}, errors.New("not found")
// 	// }
// 	return foundList, nil
// }

func (listRepo *ListRepository) UpdateList(ctx *gin.Context, listToGet db.GetListForUpdateParams, listToUpdate db.UpdateListParams) (db.List, error) {
	_, err := (*listRepo.Store).GetListForUpdate(ctx, listToGet)
	if err != nil {
		return db.List{}, err
	}

	listToUpdate.UpdatedAt = time.Now()
	updatedList, err := (*listRepo.Store).UpdateList(ctx, listToUpdate)

	if err != nil && strings.Index(err.Error(), "duplicate key value violates unique constraint") > 1 {
		err = errors.New("list name already used")
	}

	return updatedList, err
}

// func (listRepo *ListRepository) DeleteList(listId uint) (models.List, error) {
// 	var deletedList models.List
// 	// if result := listRepo.DB.Delete(&deletedList, listId); result.RowsAffected == 0 {
// 	// 	return models.List{}, errors.New("error trying to deleted")
// 	// }

// 	return deletedList, nil
// }
