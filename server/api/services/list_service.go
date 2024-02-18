package services

import (
	"github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/api/repositories"
	"github.com/RomainC75/todo2/data/models"
)

type ListSrv struct {
	listRepository repositories.ListRepositoryInterface
}

func NewListSrv() *ListSrv {
	return &ListSrv{
		listRepository: repositories.NewListRepo(),
	}
}

func (listSrv *ListSrv) CreateListSrv(userId uint, list requests.CreateListRequest) (models.List, error) {
	var newList models.List

	newList.Name = list.Name
	newList.UserRefer = userId

	createdList, err := listSrv.listRepository.CreateList(newList)
	if err != nil {
		return models.List{}, err
	}
	return createdList, nil
}
