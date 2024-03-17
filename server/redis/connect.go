package redis

import (
	"fmt"
	"log"
	"os"

	"github.com/RomainC75/todo2/config"
	"github.com/gomodule/redigo/redis"
)

// type RedisInfos struct {
// 	redisClient *redis.Pool
// }

var redisClient *redis.Pool

func ConnectRedis() {
	config := config.Get()
	address := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	fmt.Println("=> address: ", address)

	redisClient = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address)
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			fmt.Println("=> err inside: ", err)
			return conn, err
		},
	}
	// Check the connection
	ping(redisClient.Get())
}

func ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Printf("ERROR: fail ping redis conn: %s", err.Error())
		os.Exit(1)
	}
}

// func set(conn redis.Conn, key string, val string) error {
// 	// get conn and put back when exit from method
// 	// defer conn.Close()

// 	_, err := conn.Do("SET", key, val)
// 	if err != nil {
// 		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
// 		return err
// 	}

// 	return nil
// }

func Get() *redis.Pool {
	return redisClient
}
