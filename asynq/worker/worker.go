package asynq_worker

import (
	"fmt"
	"github.com/agniBit/bench-mark/app/config"
	"github.com/agniBit/bench-mark/model"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

func StartServer(cfg *config.Config) {
	t := time.Now()
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     cfg.Redis.Addr,
			PoolSize: cfg.Redis.PoolSize,
		},
		asynq.Config{
			Concurrency: cfg.Worker.Concurrency,
			Queues: map[string]int{
				"default": 1,
			},
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
