package asynq_scheduler

import (
	"github.com/agniBit/bench-mark/app/config"
	"github.com/hibiken/asynq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var TotalJobScheduledType1 int64 = 0
var TotalJobScheduledType2 int64 = 0

func StartScheduler(cfg *config.Config) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     cfg.Redis.Addr,
			PoolSize: cfg.Redis.PoolSize,
		},
		asynq.Config{
			Concurrency: cfg.Worker.Concurrency,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: cfg.Redis.Addr},
		&asynq.SchedulerOpts{Location: time.Local},
	)

	registerCrons(scheduler)

	// register worker jobs
	if err := scheduler.Run(); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

	// shutdown gracefully on SIGINT or SIGTERM
	sings := make(chan os.Signal, 1)
	signal.Notify(sings, syscall.SIGINT, syscall.SIGTERM)
	<-sings
	log.Println("shutting down scheduler worker server...")
	log.Printf("total job scheduled type 1: %d", TotalJobScheduledType1)
	log.Printf("total job scheduled type 2: %d", TotalJobScheduledType2)
	srv.Shutdown()
}
