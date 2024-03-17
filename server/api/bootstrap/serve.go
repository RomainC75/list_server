package bootstrap

import (
	"fmt"

	Routing "github.com/RomainC75/todo2/api/routing"
	"github.com/RomainC75/todo2/config"
	db "github.com/RomainC75/todo2/db/sqlc"
	redis_server_handler "github.com/RomainC75/todo2/redis"
)

func Serve() {
	config.Set()

	db.Connect()

	redis_server_handler.ConnectRedis()
	redis_server_handler.CreateMessagePublisher(redis_server_handler.Get())
	fmt.Println("pub done")
	redis_server_handler.GoSubscribe(redis_server_handler.Get())
	fmt.Println("sub done")

	Routing.Init()

	Routing.RegisterRoutes()

	Routing.Serve()
}
