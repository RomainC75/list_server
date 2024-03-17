package redis_server_handler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/RomainC75/todo2/config"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type Context struct {
	customerID int64
}

func GoSubscribe(redisPool *redis.Pool) {
	go func() {
		config := config.Get()

		redis_namespace := config.Redis.TcpNameSpace
		redis_job_queue := config.Redis.TcpJobQueueProgressionSub

		// WorkerPool => NAMESPACE
		pool := work.NewWorkerPool(Context{}, 10, redis_namespace, redisPool)

		// middlewares execute functions on each job !!
		pool.Middleware((*Context).Log)
		pool.Middleware((*Context).VerifyMiddleware)

		// Job => JOB_QUEUE
		pool.Job(redis_job_queue, (*Context).HandleProgression)

		pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

		// Start processing jobs
		pool.Start()

		// Wait for a signal to quit:
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, os.Kill)

		<-signalChan
		pool.Stop()
	}()
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	// utils.PrettyDisplay(" LOG() : ", job)
	// utils.PrettyDisplay(" LOG POST() : ", job.Args["message"])

	return next()
}

func (c *Context) VerifyMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	// do something // return error if not valid
	return next()
}

func (c *Context) HandleProgression(job *work.Job) error {
	fmt.Println("==> job : ", job.Args)
	return nil
}

func (c *Context) Export(job *work.Job) error {
	return nil
}
