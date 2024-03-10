package repositories

import (
	db "github.com/RomainC75/todo2/db/sqlc"
)

type ListRepository struct {
	DB *db.Store
}

func NewListRepo() *ListRepository {
	return &ListRepository{
		DB: db.GetConnection(),
	}
}

// func (listRepo *ListRepository) CreateList(list models.List) (models.List, error) {
// 	var newList models.List
// 	// if result := listRepo.DB.Create(&list).Scan(&newList); result.RowsAffected == 0 {
// 	// 	return models.List{}, errors.New("error trying to create a new user :-(")
// 	// }
// 	return newList, nil
// }

// func (listRepo *ListRepository) GetLists(userId uint) []models.List {
// 	var foundLists []models.List
// 	// listRepo.DB.Where("user_refer = ?", userId).Find(&foundLists)
// 	return foundLists
// }

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
