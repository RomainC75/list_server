package repositories

import (
	"database/sql"
	"errors"
	"fmt"
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

func (listRepo *ListRepository) GetListByIdByOwner(ctx *gin.Context, getListParams db.GetlistParams) (db.List, error) {
	foundList, err := (*listRepo.Store).Getlist(ctx, getListParams)
	if err == sql.ErrNoRows {
		fmt.Println("list not found : ", err.Error())
		return db.List{}, errors.New("no list found")
	}
	return foundList, err
}

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

func (listRepo *ListRepository) DeleteList(ctx *gin.Context, listToDelete db.DeleteListParams) (db.List, error) {
	// TODO : delete items too !!
	deletedList, err := (*listRepo.Store).DeleteList(ctx, listToDelete)
	if err == sql.ErrNoRows {
		err = errors.New("no list found")
	}
	return deletedList, err
}
