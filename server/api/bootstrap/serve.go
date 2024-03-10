package bootstrap

import (
	Routing "github.com/RomainC75/todo2/api/routing"
	"github.com/RomainC75/todo2/config"
	db "github.com/RomainC75/todo2/db/sqlc"
)

func Serve() {
	config.Set()

	db.Connect()

	Routing.Init()

	Routing.RegisterRoutes()

	Routing.Serve()
}
