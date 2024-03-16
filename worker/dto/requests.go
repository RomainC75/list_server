package redis_dto

import (
	"time"

	"github.com/google/uuid"
)

type ScanRequestMessage struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Address   string    `json:"address"`
	PortMin   int       `json:"port_min"`
	PortMax   int       `json:"port_max"`
}
