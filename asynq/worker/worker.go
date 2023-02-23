package asynq_worker

import (
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"

	"github.com/agniBit/benchmark/app/config"
	"github.com/agniBit/benchmark/model"
)

var QueuePriorityMap = map[string]int{
	string(model.QueuePriorityCritical): 6,
	string(model.QueuePriorityDefault):  3,
	string(model.QueuePriorityLow):      1,
}

func StartServer(cfg *config.Config) {
	t := time.Now()
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     cfg.Redis.Addr,
			PoolSize: cfg.Redis.PoolSize,
		},
		asynq.Config{
			Concurrency: cfg.Worker.Concurrency,
			Queues:      QueuePriorityMap,
		},
	)

	mux := asynq.NewServeMux()

	// register worker jobs
	for i := 0; i < cfg.Job.QueueCount; i++ {
		mux.HandleFunc(fmt.Sprintf("%s-%d", string(model.DummyJob), i), handleFunc)
	}

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

	fmt.Println("shutting down consumer worker server...")
	fmt.Println("worker stopped", time.Since(t))
	srv.Shutdown()
}
