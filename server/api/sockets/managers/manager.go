package managers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	SocketMessage "github.com/RomainC75/todo2/api/dto/requests"
	"github.com/RomainC75/todo2/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ManagerInterface interface {
	ServeWS(w gin.ResponseWriter, r *http.Request, userData UserData)
}

type Hub struct {
	uuid      uuid.UUID
	createdAt time.Time
}

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			cfg := config.Get()
			frontUrl := cfg.Front.Host
			return origin == frontUrl
		},
	}
)

type Manager struct {
	clients ClientList
	jobs    JobList
	hubs    []Hub
	sync.RWMutex
}

func New() *Manager {
	manager := Manager{
		jobs:    make(JobList),
		clients: make(ClientList),
	}
	return &manager
}

func (m *Manager) ServeWS(w gin.ResponseWriter, r *http.Request, userData UserData) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// add to client list
	client := NewClient(conn, m, userData)

	m.AddClient(client)
	go client.writeMessages()
	m.CreateJob(client)
	// test ...
	client.ResponseToClient(SocketMessage.WebSocketMessage{
		Type: "connection",
		Content: map[string]string{
			"message": "connected to server",
		},
	})
}

func (m *Manager) AddClient(client *Client) {
	// do not modify at the same time, when 2 people are trying to connect at the same time.
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) RemoveClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

func (m *Manager) CreateJob(client *Client) *Job {
	m.Lock()
	defer m.Unlock()

	//create Job

	newJob := NewJob(m, client)
	// add it to the job
	m.jobs[newJob] = true
	// add the job to the client
	client.Job = newJob

	return newJob
}

func (m *Manager) BroadcastMessage(mType string, content map[string]string) {
	wsMessage := SocketMessage.WebSocketMessage{
		Type:    mType,
		Content: content,
	}
	fmt.Println("++++++broadcasting message", m.clients)
	for client := range m.clients {
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}
