package redis_server_handler

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	redis_dto "github.com/RomainC75/todo2/redis/dto"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
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

func (jq *JobQueue) PublishMessage(m redis_dto.ScanRequestMessage, queueName string) {
	// w, err := (*jq).enqueuer.Enqueue("requests", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})

	m.Id = uuid.New()
	m.CreatedAt = time.Now()

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal("or trying to serialyse the message to json")
	}
	fmt.Println("=> message: ", string(b))
	_, err = (*jq).enqueuer.Enqueue("requests", work.Q{"message": string(b)})
	if err != nil {
		log.Fatal(err)
	}
}

func GetJobQueue() *JobQueue {
	return jobQueue
}
