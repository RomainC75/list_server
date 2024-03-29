package repositories

import (
	"time"

	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ItemRepository struct {
	Store *db.Store
}

func NewItemRepo() *ItemRepository {
	// lksjdf
	return &ItemRepository{
		Store: db.GetConnection(),
	}
}

func (itemRepo *ItemRepository) CreateItem(ctx *gin.Context, itemToCreate db.CreateItemParams, listId int32) (db.Item, error) {
	var createdItem db.Item
	err := (*itemRepo.Store).ExecTx(ctx, func(q *db.Queries) error {
		var err error

		itemToCreate.CreatedAt = time.Now()
		itemToCreate.UpdatedAt = itemToCreate.CreatedAt
		createdItem, err = q.CreateItem(ctx, itemToCreate)
		if err != nil {
			return err
		}

		var linkItemToListParam db.LinkItemToListParams
		linkItemToListParam.ItemID = createdItem.ID
		linkItemToListParam.ListID = listId
		_, err = q.LinkItemToList(ctx, linkItemToListParam)
		if err != nil {
			return err
		}
		return nil
	})
	return createdItem, err
}

func (itemRepo *ItemRepository) GetItems(ctx *gin.Context, listId int32) ([]db.Item, error) {
	foundItems, err := (*itemRepo.Store).GetItemsByListName(ctx, listId)
	return foundItems, err
}

func (itemRepo *ItemRepository) UpdateItem(ctx *gin.Context, arg db.UpdateItemParams) (db.Item, error) {
	arg.UpdatedAt = time.Now()
	return (*itemRepo.Store).UpdateItem(ctx, arg)
}
