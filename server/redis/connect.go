package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RomainC75/todo2/config"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	RedisClient redis.Client
}

var redisClient Redis

func ConnectRedis() {
	config := config.Get()
	var client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		// Password: config.Redis.Password,
		DB: 0,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		// Sleep for 3 seconds and wait for Redis to initialize
		time.Sleep(3 * time.Second)
		err := client.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}
	}

	if client == nil {
		log.Fatal("Cannot run redis")
	}
	redisClient = Redis{
		RedisClient: *client,
	}
}

func Get() Redis {
	return redisClient
}
