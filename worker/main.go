package main

import (
	"log"
	"os"
	"os/signal"
	redisHandler "worker/redis"

	"github.com/gocraft/work"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	redis_domain := os.Getenv("REDIS_DOMAIN")
	redis_port := os.Getenv("REDIS_PORT")
	redis_namespace := os.Getenv("REDIS_NAMESPACE")
	redis_job_queue := os.Getenv("REDIS_JOB_QUEUE")
	if redis_domain == "" || redis_port == "" || redis_namespace == "" || redis_job_queue == "" {
		log.Fatal("Error loading .env file - check the variables !!!")
		return
	}

	redisHandler.NewPool(redis_domain, redis_port)

	// WorkerPool => NAMESPACE
	pool := work.NewWorkerPool(redisHandler.Context{}, 10, redis_namespace, redisHandler.GetPool())
	// Pub
	redisHandler.CreateMessagePublisher(redisHandler.GetPool())

	// middlewares execute functions on each job !!
	pool.Middleware((*redisHandler.Context).Log)
	pool.Middleware((*redisHandler.Context).VerifyMiddleware)

	// Job => JOB_QUEUE
	pool.Job(redis_job_queue, (*redisHandler.Context).Scan)

	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*redisHandler.Context).Export)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	pool.Stop()
}
