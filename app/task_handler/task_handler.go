package task_handler

import (
	"context"
	"fmt"
	"github.com/agniBit/bench-mark/app/config"
	"math/rand"
	"sync"
	"time"
)

var totalTasks = 1
var t = time.Now()
var mutex = &sync.Mutex{}

func DummyTask(ctx context.Context, taskID string) error {
	cfg := config.Get()
	if cfg.Worker.RandomDelay > 0 {
		sleepTime := time.Duration(rand.Intn(config.Get().Worker.RandomDelay)) * time.Millisecond
		time.Sleep(sleepTime)
	}
	if rand.Intn(100) < cfg.Job.ErrorRate {
		return fmt.Errorf("task1 failed")
	}
	mutex.Lock()
	totalTasks++
	if totalTasks >= cfg.Job.TaskPerQueue*cfg.Job.QueueCount {
		fmt.Println("all task completed", time.Since(t))
	}
	if totalTasks%1000 == 0 {
		fmt.Println("task count", totalTasks)
	}
	mutex.Unlock()
	return nil
}
