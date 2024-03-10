package repositories

import (
	"fmt"
	"time"

	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ItemRepository struct {
	Store *db.Store
}

func NewItemRepo() *ItemRepository {
	return &ItemRepository{
		Store: db.GetConnection(),
	}
}

func (itemRepo *ItemRepository) CreateItem(ctx *gin.Context, itemToCreate db.CreateItemParams, listId int32) (db.Item, error) {
	itemToCreate.CreatedAt = time.Now()
	itemToCreate.UpdatedAt = itemToCreate.CreatedAt
	createdItem, err := (*itemRepo.Store).CreateItem(ctx, itemToCreate)
	fmt.Println("==> createdItem : ", createdItem, err)
	if err != nil {
		return db.Item{}, err
	}

	var linkItemToListParam db.LinkItemToListParams
	linkItemToListParam.ItemID = createdItem.ID
	linkItemToListParam.ListID = listId
	_, err = (*itemRepo.Store).LinkItemToList(ctx, linkItemToListParam)
	if err != nil {
		fmt.Println("=> is error creating new item ? ", err.Error())
	}
	// if err != nil {
	// TODO: DELETE ITEM ?
	// TODO : transactions Tx??
	// }
	return createdItem, nil
}
