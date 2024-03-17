package managers

import (
	"encoding/json"
	"fmt"
	"sync"

	SocketMessage "github.com/RomainC75/todo2/api/dto/requests"
	"github.com/google/uuid"
)

type JobList map[*Room]bool

type JobBasicInfos struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type Room struct {
	sync.RWMutex
	Id   uuid.UUID
	Name string
	JobBasicInfos
	Manager *Manager
	Clients ClientList
}

func NewJob(manager *Manager, uuid uuid.UUID) *Room {
	return &Room{
		Id:      uuid,
		Manager: manager,
	}
}

func (r *Room) AddClient(client *Client) {
	r.Clients[client] = true
}

func (r *Room) RemoveClient(client *Client) {
	r.Lock()
	defer r.Unlock()
	delete(r.Clients, client)
}

func (r *Room) BroadcastMessage(wsMessage SocketMessage.WebSocketMessage) {
	for client := range r.Clients {
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}
