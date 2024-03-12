package responses

import (
	"time"

	db "github.com/RomainC75/todo2/db/sqlc"
)

type ItemResponse struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetItemResponse(item db.Item) (itemResponse ItemResponse) {
	itemResponse.ID = item.ID
	itemResponse.Name = item.Name
	itemResponse.Description = item.Description.String
	itemResponse.Date = item.Date.Time
	itemResponse.CreatedAt = item.CreatedAt
	itemResponse.UpdatedAt = item.UpdatedAt
	return itemResponse
}

func GetItemsResponse(items []db.Item) []ItemResponse {
	itemsResponse := []ItemResponse{}
	for _, item := range items {
		itemsResponse = append(itemsResponse, GetItemResponse(item))
	}
	return itemsResponse
}
