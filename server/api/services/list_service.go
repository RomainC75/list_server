package services

import (
	"fmt"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/repositories"
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListSrv struct {
	listRepo repositories.ListRepositoryInterface
}

func NewListSrv() *ListSrv {
	return &ListSrv{
		listRepo: repositories.NewListRepo(),
	}
}

func (listSrv *ListSrv) CreateListSrv(ctx *gin.Context, userId int32, list requests.CreateListReq) (db.List, error) {
	foundLists, _ := listSrv.listRepo.GetLists(ctx, userId)
	for _, foundList := range foundLists {
		if foundList.Name == list.Name {
			return db.List{}, fmt.Errorf("duplicate name : %s", list.Name)
		}
	}
	var newList db.CreateListParams

	newList.Name = list.Name
	newList.UserID = userId
	createdList, err := listSrv.listRepo.CreateList(ctx, newList)
	if err != nil {
		return db.List{}, err
	}
	return createdList, nil
}

func (listSrv *ListSrv) GetListsByUserIdSrv(ctx *gin.Context, userId int32) ([]db.List, error) {
	return listSrv.listRepo.GetLists(ctx, userId)
}

func (listSrv *ListSrv) UpdateListNameSrv(ctx *gin.Context, userId int32, listId int32, list requests.UpdateListReq) (db.List, error) {
	var listToGet db.GetListForUpdateParams
	listToGet.ID = listId
	listToGet.UserID = userId

	var listToUpdate db.UpdateListParams
	listToUpdate.ID = listId
	listToUpdate.Name = list.Name

	return listSrv.listRepo.UpdateList(ctx, listToGet, listToUpdate)
}

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
