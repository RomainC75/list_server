package bootstrap

import (
	Routing "github.com/RomainC75/todo2/api/routing"
	"github.com/RomainC75/todo2/config"
	"github.com/RomainC75/todo2/data/database"
)

func Serve() {
	config.Set()

	database.Connect()

	Routing.Init()

	Routing.RegisterRoutes()

	Routing.Serve()
}
