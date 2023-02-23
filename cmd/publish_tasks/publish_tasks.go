package main

import (
	"flag"
	"fmt"
	"github.com/agniBit/benchmark/app/config"
	"github.com/agniBit/benchmark/app/publisher"
	"time"
)

func main() {
	wordPtr := flag.String("publishTo", "asynq", "please provide a client name")
	flag.Parse()
	if wordPtr == nil {
		fmt.Println("Please provide a client name")
		return
	}

	config.Load()
	t := time.Now()
	if *wordPtr == "asynq" || *wordPtr == "gocraft" {
		publisher.PublishTasks(config.Get(), *wordPtr)
	} else {
		fmt.Println("Please provide a valid client name")
		return
	}
	cfg := config.Get()
	fmt.Printf("Time taken to publish %d tasks: %v\n", cfg.Job.TaskPerQueue*cfg.Job.QueueCount, time.Since(t))
}
