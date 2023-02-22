package asynq_scheduler

import (
	"fmt"
	"github.com/agniBit/bench-mark/model"
	"github.com/hibiken/asynq"
	"time"
)

func registerCrons(scheduler *asynq.Scheduler) {
	_, err := scheduler.Register("*/1 * * * *", asynq.NewTask(string(model.DummyJob2), []byte(fmt.Sprintf("cron1-%v", time.Now().Local()))))
	if err != nil {
		fmt.Println("unable to register cron", err.Error())
		panic(err)
	}
	_, err = scheduler.Register("*/2 * * * *", asynq.NewTask(string(model.DummyJob1), []byte(fmt.Sprintf("cron2-%v", time.Now().Local()))))
	if err != nil {
		fmt.Println("unable to register cron2", err.Error())
		panic(err)
	}
}
