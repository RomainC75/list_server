package redis

import (
	"log"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type JobQueue struct {
	enqueuer *work.Enqueuer
}

var jobQueue *JobQueue

type JobQueueInterface interface {
	PublishMessage(message interface{}, queueName string)
}

func CreateMessagePublisher(redisClient *redis.Pool) {
	jobQueue = &JobQueue{
		enqueuer: work.NewEnqueuer("tcp_scan", redisClient),
	}
}

func (jq *JobQueue) PublishMessage(message interface{}, queueName string) {
	// w, err := (*jq).enqueuer.Enqueue("requests", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})
	_, err := (*jq).enqueuer.Enqueue("requests", work.Q{"message": message})
	if err != nil {
		log.Fatal(err)
	}
}

func GetJobQueue() *JobQueue {
	return jobQueue
}
