package redis_dto

import (
	"time"

	"github.com/google/uuid"
)

type ScanResponseMessage struct {
	Id        uuid.UUID `json:"id"`
	RequestID uuid.UUID `json:"request_id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
}
