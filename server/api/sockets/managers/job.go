package managers

import (
	"encoding/json"
	"fmt"
	"sync"

	SocketMessage "github.com/RomainC75/todo2/api/dto/requests"
	"github.com/google/uuid"
)

type JobList map[*Job]bool

type JobBasicInfos struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type Job struct {
	sync.RWMutex
	Id   uuid.UUID
	Name string
	JobBasicInfos
	Manager *Manager
	Clients ClientList
}

func NewJob(manager *Manager, client *Client) *Job {
	newJob := &Job{
		Id:      client.userData.RequestId,
		Manager: manager,
		Clients: make(ClientList),
	}
	newJob.AddClient(client)
	return newJob
}

func (r *Job) AddClient(client *Client) {
	r.Clients[client] = true
}

func (r *Job) RemoveClient(client *Client) {
	r.Lock()
	defer r.Unlock()
	delete(r.Clients, client)
}

func (r *Job) BroadcastMessage(wsMessage SocketMessage.WebSocketMessage) {
	for client := range r.Clients {
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}
