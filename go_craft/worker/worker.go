package gocraft_worker

import (
	"fmt"
	"github.com/GetSimpl/work"
	"github.com/agniBit/benchmark/app/config"
	"github.com/agniBit/benchmark/model"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Context struct{}

func RunWorkers(cfg *config.Config) {
	t := time.Now()
	redisPool := &redis.Pool{
		MaxActive: cfg.Redis.PoolSize,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Redis.Addr)
		},
	}

	pool := work.NewWorkerPool(Context{}, uint(cfg.Worker.Concurrency), "test", redisPool)

	for i := 0; i < cfg.Job.QueueCount; i++ {
		pool.JobWithOptions(fmt.Sprintf("%s-%d", string(model.DummyJob), i),
			work.JobOptions{Priority: 1000, MaxFails: 3, MaxConcurrency: uint(cfg.Worker.Concurrency)}, handler)
	}

	pool.Start()
	defer pool.Stop()

	fmt.Println("worker started")
	waitForSIGINT()
	fmt.Println("worker stopped", time.Since(t))
}

func waitForSIGINT() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
