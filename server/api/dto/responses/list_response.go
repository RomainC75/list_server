package responses

import (
	"time"

	db "github.com/RomainC75/todo2/db/sqlc"
)

type ListResponse struct {
	Id        int32      `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Name      string     `json:"name"`
}

func GetListResponseFromModel(list db.List) ListResponse {
	return ListResponse{
		Id:        list.ID,
		CreatedAt: &list.CreatedAt,
		UpdatedAt: &list.UpdatedAt,
		Name:      list.Name,
	}
}

func GetListResponseFromModelList(lists []db.List) []ListResponse {
	listsResponse := []ListResponse{}
	for _, lr := range lists {
		listsResponse = append(listsResponse, GetListResponseFromModel(lr))
	}
	return listsResponse
}
