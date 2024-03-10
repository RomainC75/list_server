package services

import (
	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/repositories"
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListSrv struct {
	listRepository repositories.ListRepositoryInterface
}

func NewListSrv() *ListSrv {
	return &ListSrv{
		// listRepository: repositories.NewListRepo(),
	}
}

func (listSrv *ListSrv) CreateListSrv(ctx *gin.Context, userId int32, list requests.CreateListReq) (db.List, error) {
	var newList db.CreateListParams

	newList.Name = list.Name
	newList.UserID = userId

	createdList, err := listSrv.listRepository.CreateList(ctx, newList)
	if err != nil {
		return db.List{}, err
	}
	return createdList, nil
}

// func (listSrv *ListSrv) GetListsByUserIdSrv(userId uint) []models.List {
// 	return listSrv.listRepository.GetLists(userId)
// }

// func (listSrv *ListSrv) GetListOwnedByUser(userId uint, listId uint) (db.List, error) {
// 	foundList, err := listSrv.listRepository.GetListById(listId)
// 	if err != nil {
// 		return models.List{}, err
// 	}

// 	if foundList.UserRefer != userId {
// 		return models.List{}, errors.New("not the owner of this list")
// 	}
// 	return models.List, nil
// 	return foundList, nil
// }

// func (listSrv *ListSrv) UpdateList(userId uint, list requests.UpdateListRequest) (models.List, error) {
// 	updatedList, err := listSrv.listRepository.UpdateList(userId, list)
// 	if err != nil {
// 		return models.List{}, err
// 	}
// 	return updatedList, nil
// }

// func (listSrv *ListSrv) DeleteList(userId uint, listId uint) (models.List, error) {
// 	foundList, err := listSrv.listRepository.GetListById(listId)
// 	if err != nil {
// 		return models.List{}, nil
// 	}

// 	if foundList.UserRefer != userId {
// 		return models.List{}, errors.New("not the owner of this list")
// 	}

// 	deleteList, err := listSrv.listRepository.DeleteList(listId)
// 	if err != nil {
// 		return models.List{}, err
// 	}
// 	return deleteList, nil
// }
