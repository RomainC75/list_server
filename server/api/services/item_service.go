package services

import (
	"github.com/RomainC75/todo2/api/repositories"
	db "github.com/RomainC75/todo2/db/sqlc"
)

type ItemSrv struct {
	itemRepository repositories.ItemRepositoryInterface
	listRepository repositories.ListRepositoryInterface
}

func NewItemSrv() *ItemSrv {
	return &ItemSrv{
		// itemRepository: repositories.NewItemRepo(),
	}
}

func (itemSrv *ItemSrv) CreateItemSrv(userId uint, listId uint, item db.Item) (db.Item, error) {
	// foundList, err := itemSrv.listRepository.GetListById(listId)
	// if err != nil {
	// 	return db.Item{}, err
	// }

	// if foundList.UserRefer != userId {
	// 	return db.Item{}, errors.New("unauthorized to create  an item to this list")
	// }

	// itemToCreate := db.Item{
	// 	Name:        item.Name,
	// 	Description: item.Description,
	// 	Date:        *item.Date,
	// }

	// createdItem, err := itemSrv.itemRepository.CreateItem(itemToCreate, foundList)
	// if err != nil {
	// 	return db.Item{}, err
	// }

	return db.Item{}, nil
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
