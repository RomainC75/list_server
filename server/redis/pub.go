package redis

import (
	"context"
	"encoding/json"
	"log"
)

type MessagePublisher struct {
	redisClient Redis
}

var publisher *MessagePublisher

func CreateMessagePublisher(redisClient Redis) {
	publisher = &MessagePublisher{redisClient}
}

func GetPublisher() *MessagePublisher {
	return publisher
}

func (p *MessagePublisher) PublishMessages(ctx context.Context, message interface{}, queueName string) {
	serializedMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("==> [%s] Failed to serialize message: %v", queueName, err)
	}

	err = p.redisClient.RedisClient.Publish(ctx, queueName, serializedMessage).Err()
	if err != nil {
		log.Printf("==> [%s] Failed to publish message: %v", queueName, err)
	}
}
