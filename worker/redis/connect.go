package redisHandler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	redis_dto "worker/dto"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

type Context struct {
	customerID int64
}

func NewPool(domain string, port string) {
	redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", domain, port))
		},
	}
}

func GetPool() *redis.Pool {
	return redisPool
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println(" LOG POST() : ", job.Args["message"])

	return next()
}

func (c *Context) VerifyMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	// do something // return error if not valid
	return next()
}

func (c *Context) Scan(job *work.Job) error {

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

	for i := 0; i < 20; i++ {
		// TODO: use .env !!
		fmt.Println(".......", i)
		fmt.Println("=> ", GetJobQueue())
		GetJobQueue().PublishMessage(redis_dto.ScanResponseMessage{
			RequestID: message.Id,
			Message:   strconv.Itoa(i),
		}, "progression")
		time.Sleep(time.Second)
	}

	fmt.Println(".......", message.Address)

	return nil
}

func (c *Context) Export(job *work.Job) error {
	return nil
}
