package responses

import (
	"time"

	"github.com/RomainC75/todo2/data/models"
)

type ListResponse struct {
	Id        uint       `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Name      string
}

func GetListResponseFromModel(list models.List) ListResponse {
	return ListResponse{
		Id:        list.ID,
		CreatedAt: &list.CreatedAt,
		UpdatedAt: &list.UpdatedAt,
		Name:      list.Name,
	}
}

func GetListResponseFromModelList(lists []models.List) []ListResponse {
	listsResponse := []ListResponse{}
	for _, lr := range lists {
		listsResponse = append(listsResponse, GetListResponseFromModel(lr))
	}
	return listsResponse
}
