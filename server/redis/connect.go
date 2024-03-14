package redis

import (
	"errors"

	"github.com/go-redis/redis"
)

type Redis struct {
	RedisClient redis.Client
}

func NewRedis(env Env) Redis {

	var client = redis.NewClient(&redis.Options{
		// Container name + port since we are using docker
		Addr:     "redis:6379",
		Password: env.RedisPassword,
	})

	if client == nil {
		errors.New("Cannot run redis")
	}

	return Redis{
		RedisClient: *client,
	}
}
