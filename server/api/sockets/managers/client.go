package managers

import (
	"encoding/json"
	"log"
	"sync"

	SocketMessage "github.com/RomainC75/todo2/api/dto/requests"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type UserData struct {
	UserId    uuid.UUID
	UserEmail string
	RequestId uuid.UUID
}

type Client struct {
	sync.RWMutex
	userData     UserData
	connection   *websocket.Conn
	manager      *Manager
	egress       chan []byte
	Room         *Room
	PlayerNumber int
	CommandIn    chan int
	LastCommand  int
}

func NewClient(conn *websocket.Conn, manager *Manager, userData UserData) *Client {
	return &Client{
		userData:   userData,
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte),
	}
}

// func (c *Client) WriteLastCommand(){
// 	c.L
// }

func (c *Client) writeMessages() {
	defer func() {
		c.manager.RemoveClient(c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed:", err)
				}
				break
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("failed to send message: %v", err)
			}
			log.Println("message sent")
		}
	}
}

func (c *Client) ResponseToClient(message SocketMessage.WebSocketMessage) {
	m, _ := json.Marshal(message)
	c.egress <- m
}
