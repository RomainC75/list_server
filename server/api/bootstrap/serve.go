package bootstrap

import (
	"github.com/RomainC75/todo2/config"
	"github.com/RomainC75/todo2/data/database"
	// Routing "github.com/RomainC75/todo2/api/routing"
)

func Serve() {
	config.Set()

	database.Connect()

	// Routing.Init()

	// Routing.RegisterRoutes()

	// Routing.Serve()
}
