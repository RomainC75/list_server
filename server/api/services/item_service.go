package services

import (
	"database/sql"
	"errors"
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

func (itemSrv *ItemSrv) GetAvailableItems(ctx *gin.Context) ([]db.Item, error) {
	return itemSrv.itemRepo.GetEveryItems(ctx)
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
		UserCreatorID: userId,
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

func (itemSrv *ItemSrv) GetItemsByListSrv(ctx *gin.Context, listId int32) ([]db.Item, error) {
	return itemSrv.itemRepo.GetItems(ctx, listId)
}

func (itemSrv *ItemSrv) UpdateItem(ctx *gin.Context, itemId int32, userId int32, itemRequest requests.UpdateItemRequest) (db.Item, error) {
	// TODO : check if item exists
	arg := db.UpdateItemParams{
		ID:          itemId,
		Name:        sql.NullString{String: itemRequest.Name, Valid: itemRequest.Name != ""},
		Description: sql.NullString{String: itemRequest.Description, Valid: itemRequest.Description != ""},
		Date:        sql.NullTime{Time: itemRequest.Date, Valid: !itemRequest.Date.IsZero()},
	}
	return itemSrv.itemRepo.UpdateItem(ctx, arg)
}

func (itemSrv *ItemSrv) DeleteItem(ctx *gin.Context, itemId int32, itemCreatorId int32) (db.Item, error) {
	arg := db.DeleteItemParams{
		ID:            itemId,
		UserCreatorID: itemCreatorId,
	}
	return itemSrv.itemRepo.DeleteItem(ctx, arg)
}

func (itemSrv *ItemSrv) AddItemToList(ctx *gin.Context, userId int32, listId int32, itemId int32) error {
	_, err := itemSrv.listRepo.GetListByIdByOwner(ctx, db.GetlistParams{
		ID:     listId,
		UserID: userId,
	})
	if err != nil {
		return errors.New("list not found/owned")
	}

	_, err = itemSrv.itemRepo.LinkItemToList(ctx, db.LinkItemToListParams{
		ListID: listId,
		ItemID: itemId,
	})
	if err != nil {
		return errors.New("item does not exist")
	}
	return nil
}
