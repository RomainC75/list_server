package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	redis_dto "worker/dto"

	"github.com/RomainC75/todo2/utils"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

var redisPool *redis.Pool

type Context struct {
	customerID int64
}

func NewPool(domain string, port string) *redis.Pool {
	return &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", domain, port))
		},
	}
}

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

	redisPool = NewPool(redis_domain, redis_port)

	// WorkerPool => NAMESPACE
	pool := work.NewWorkerPool(Context{}, 10, redis_namespace, redisPool)

	// middlewares execute functions on each job !!
	pool.Middleware((*Context).Log)
	pool.Middleware((*Context).VerifyMiddleware)

	// Job => JOB_QUEUE
	pool.Job(redis_job_queue, (*Context).Scan)

	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	pool.Stop()
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	utils.PrettyDisplay(" LOG() : ", job)
	utils.PrettyDisplay(" LOG POST() : ", job.Args["message"])

	return next()
}

func (c *Context) VerifyMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	// do something // return error if not valid
	return next()
}

func (c *Context) Scan(job *work.Job) error {
	utils.PrettyDisplay("data for scanning: ", job.Args["message"])

	fmt.Println(".......")
	addr := job.ArgString("message")
	fmt.Println("=> message: ", addr)
	if err := job.ArgError(); err != nil {
		fmt.Println("=> error: ", err)
		return err
	}

	var message redis_dto.ScanRequestMessage
	if err := json.Unmarshal([]byte(addr), &message); err != nil {
		return err
	}
	utils.PrettyDisplay("OBJECT ! ", message)

	fmt.Println(".......", message.Address)

	return nil
}

func (c *Context) Export(job *work.Job) error {
	return nil
}
