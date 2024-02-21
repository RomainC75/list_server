package services

import (
	"errors"

	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/repositories"
	"github.com/RomainC75/todo2/data/models"
)

type ItemSrv struct {
	itemRepository repositories.ItemRepositoryInterface
	listRepository repositories.ListRepositoryInterface
}

func NewItemSrv() *ItemSrv {
	return &ItemSrv{
		itemRepository: repositories.NewItemRepo(),
	}
}

func (itemSrv *ItemSrv) CreateItemSrv(userId uint, listId uint, item requests.ItemToCreateRequest) (models.Item, error) {
	foundList, err := itemSrv.listRepository.GetListById(listId)
	if err != nil {
		return models.Item{}, err
	}

	if foundList.UserRefer != userId {
		return models.Item{}, errors.New("unauthorized to create  an item to this list")
	}

	itemToCreate := models.Item{
		Name:        item.Name,
		Description: item.Description,
		Date:        *item.Date,
	}

	createdItem, err := itemSrv.itemRepository.CreateItem(itemToCreate, foundList)
	if err != nil {
		return models.Item{}, err
	}

	return createdItem, nil
}

// func (listSrv *ListSrv) GetListsByUserIdSrv(userId uint) []models.List {
// 	return listSrv.listRepository.GetLists(userId)
// }

// func (listSrv *ListSrv) GetListOwnedByUser(userId uint, listId uint) (models.List, error) {
// 	foundList, err := listSrv.listRepository.GetListById(listId)
// 	if err != nil {
// 		return models.List{}, err
// 	}

// 	if foundList.UserRefer != userId {
// 		return models.List{}, errors.New("not the owner of this list")
// 	}

// 	return foundList, nil
// }
