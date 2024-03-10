package repositories

import (
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

// func (listRepo *ListRepository) UpdateList(userId uint, list requests.UpdateListRequest) (models.List, error) {
// 	var foundList models.List
// 	// if result := listRepo.DB.Where("id = ?", list.Id).First(&foundList); result.RowsAffected == 0 {
// 	// 	return models.List{}, errors.New("not found")
// 	// }
// 	if foundList.UserRefer != userId {
// 		return models.List{}, errors.New("not authorized")
// 	}

// 	// result := listRepo.DB.Save(&foundList)
// 	// if result.Error != nil {
// 	//     return models.List{}, result.Error
// 	// }

// 	// if result := listRepo.DB.Model(&foundList).Update("name", list.Name); result.Error != nil {
// 	// 	return models.List{}, result.Error
// 	// }

// 	return foundList, nil
// }

// func (listRepo *ListRepository) DeleteList(listId uint) (models.List, error) {
// 	var deletedList models.List
// 	// if result := listRepo.DB.Delete(&deletedList, listId); result.RowsAffected == 0 {
// 	// 	return models.List{}, errors.New("error trying to deleted")
// 	// }

// 	return deletedList, nil
// }
