package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "pass",
		DB: 0,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		time.Sleep(3 * time.Second)
		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}
	}
	ctx := context.Background()
	topic := redisClient.Subscribe(ctx, "myqueue")
	defer topic.Close()

	channel := topic.Channel()
	for msg := range channel {
		fmt.Println("=> payload : ", msg.Payload)
		u := &Message{}
		err := u.UnmarshalBinary([]byte(msg.Payload))
		if err != nil {
			panic(err)
		}

		fmt.Println(u)
	}
}

type Message struct {
	Message string
}

func (u *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Message) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, u); err != nil {
		return err
	}
	return nil
}

func (u *Message) String() string {
	return "You received : " + u.Message
}
