package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gocraft/work/webui"
	"github.com/gomodule/redigo/redis"

	"github.com/agniBit/bench-mark/app/config"
)

func main() {
	config.Load()
	cfg := config.Get()
	redisPool := &redis.Pool{
		MaxActive: cfg.Redis.PoolSize,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Redis.Addr)
		},
	}
	server := webui.NewServer("test", redisPool, ":9875")
	server.Start()
	fmt.Println("worker started on port 9875, see on http://localhost:9875/worker")
	waitForSIGINT()

	defer server.Stop()
}

func waitForSIGINT() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
