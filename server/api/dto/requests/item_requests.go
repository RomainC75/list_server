package requests

import "time"

type CreateItemRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
