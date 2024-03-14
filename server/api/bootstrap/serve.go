package bootstrap

import (
	Routing "github.com/RomainC75/todo2/api/routing"
	"github.com/RomainC75/todo2/config"
	db "github.com/RomainC75/todo2/db/sqlc"
	"github.com/RomainC75/todo2/redis"
)

func Serve() {
	config.Set()

	db.Connect()

	redis.ConnectRedis()
	redis.CreateMessagePublisher(redis.Get())

	Routing.Init()

	Routing.RegisterRoutes()

	Routing.Serve()
}
