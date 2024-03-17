package managers

import (
	"net/http"
	"sync"
	"time"

	"github.com/RomainC75/todo2/config"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ManagerInterface interface {
	// ServeWS(w gin.ResponseWriter, r *http.Request, userData UserData)
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
	// clients ClientList
	hubs []Hub
	sync.RWMutex
}
