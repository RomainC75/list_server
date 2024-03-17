package redisHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	redis_dto "worker/dto"

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

func (jq *JobQueue) PublishMessage(m redis_dto.ScanResponseMessage, queueName string) {

	m.Id = uuid.New()
	m.CreatedAt = time.Now()

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal("or trying to serialyse the message to json")
	}
	fmt.Println("=> message: ", string(b))
	_, err = (*jq).enqueuer.Enqueue(queueName, work.Q{"message": string(b)})
	if err != nil {
		log.Fatal(err)
	}
}

func GetJobQueue() *JobQueue {
	return jobQueue
}
