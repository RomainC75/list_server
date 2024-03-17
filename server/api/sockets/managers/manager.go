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
		jobs: make(JobList),
	}
	return &manager
}

func (m *Manager) ServeWS(w gin.ResponseWriter, r *http.Request, userData UserData) {
	log.Println("new Connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// add to client list
	client := NewClient(conn, m, userData)

	m.AddClient(client)
	go client.writeMessages()
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

func (m *Manager) BroadcastMessage(mType string, content map[string]string) {
	wsMessage := SocketMessage.WebSocketMessage{
		Type:    mType,
		Content: content,
	}
	for client := range m.clients {
		fmt.Println("send.....")
		b, _ := json.Marshal(wsMessage)
		client.egress <- b
	}
}
