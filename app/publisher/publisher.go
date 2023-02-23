package publisher

import (
	"encoding/json"
	"fmt"
	"github.com/GetSimpl/work"
	"github.com/agniBit/benchmark/app/config"
	"github.com/agniBit/benchmark/model"
	"github.com/gomodule/redigo/redis"
	"github.com/hibiken/asynq"
	"math/rand"
	"sync"
)

func PublishTasks(cfg *config.Config, enqueueClient string) {
	wg := &sync.WaitGroup{}
	var asynqClient *asynq.Client
	var goCraftClient *work.Enqueuer

	// initialize client
	if enqueueClient == "asynq" {
		asynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.Redis.Addr})
	} else if enqueueClient == "gocraft" {
		redisPool := &redis.Pool{
			MaxActive: cfg.Redis.PoolSize,
			Wait:      true,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", cfg.Redis.Addr)
			},
		}

		goCraftClient = work.NewEnqueuer("test", redisPool)
	}

	queuePriority := []model.QueuePriority{model.QueuePriorityCritical, model.QueuePriorityDefault, model.QueuePriorityLow}

	// run 100 go-routines to publish tasks
	for i := 0; i < cfg.Job.PublisherCount; i++ {
		wg.Add(1)
		go func(publisherID int) {
			defer wg.Done()
			// publish 100 tasks
			for i := 0; i < cfg.Job.TaskPerQueue/cfg.Job.PublisherCount; i++ {
				for j := 0; j < cfg.Job.QueueCount; j++ {
					p := model.Payload{
						TaskID: fmt.Sprintf("publisher-%d-task-%d", publisherID, i),
					}
					payload, err := json.Marshal(p)
					if err != nil {
						fmt.Println("unable to marshal payload", err.Error())
					}
					if enqueueClient == "asynq" {
						p := string(queuePriority[rand.Intn(3)])
						task := asynq.NewTask(fmt.Sprintf("%s-%d", string(model.DummyJob), j), payload, asynq.Queue(p), asynq.MaxRetry(3))
						_, err = asynqClient.Enqueue(task)
					} else if enqueueClient == "gocraft" {
						_, err = goCraftClient.Enqueue(fmt.Sprintf("%s-%d", string(model.DummyJob), j), work.Q{"payload": payload})
					}
					if err != nil {
						fmt.Println("unable to enqueue task", err.Error())
					}
				}
			}
		}(i)
	}

	wg.Wait()
}
