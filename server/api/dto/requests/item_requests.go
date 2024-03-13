package requests

import "time"

type CreateItemRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type UpdateItemRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
