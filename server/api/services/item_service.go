package services

import (
	"database/sql"
	"fmt"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/repositories"
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/RomainC75/todo2/utils"
	"github.com/gin-gonic/gin"
)

type ItemSrv struct {
	itemRepo repositories.ItemRepositoryInterface
	listRepo repositories.ListRepositoryInterface
}

func NewItemSrv() *ItemSrv {
	return &ItemSrv{
		itemRepo: repositories.NewItemRepo(),
		listRepo: repositories.NewListRepo(),
	}
}

func (itemSrv *ItemSrv) CreateItemSrv(ctx *gin.Context, userId int32, listId int32, item requests.CreateItemRequest) (db.Item, error) {

	foundList, err := itemSrv.listRepo.GetListByIdByOwner(ctx, db.GetlistParams{
		ID:     listId,
		UserID: userId,
	})
	if err != nil {
		return db.Item{}, err
	}
	fmt.Println("foundList : ", foundList)
	utils.PrettyDisplay("new item to create : ", item)

	itemToCreate := db.CreateItemParams{
		Name: item.Name,
		Description: sql.NullString{
			String: item.Description,
			Valid:  true,
		},
	}
	if !item.Date.IsZero() {
		itemToCreate.Date = sql.NullTime{
			Time:  item.Date,
			Valid: true,
		}
	}

	createdItem, err := itemSrv.itemRepo.CreateItem(ctx, itemToCreate, listId)
	if err != nil {
		return db.Item{}, err
	}

	return createdItem, nil
}

// func (listSrv *ListSrv) GetListsByUserIdSrv(userId uint) []db.List {
// 	return listSrv.listRepository.GetLists(userId)
// }

// func (listSrv *ListSrv) GetListOwnedByUser(userId uint, listId uint) (db.List, error) {
// 	foundList, err := listSrv.listRepository.GetListById(listId)
// 	if err != nil {
// 		return db.List{}, err
// 	}

// 	if foundList.UserRefer != userId {
// 		return db.List{}, errors.New("not the owner of this list")
// 	}

// 	return foundList, nil
// }
