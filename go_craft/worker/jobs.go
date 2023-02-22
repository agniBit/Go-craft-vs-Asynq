package gocraft_worker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/GetSimpl/work"
	"github.com/agniBit/bench-mark/app/task_handler"
	"github.com/agniBit/bench-mark/model"
)

type Job struct {
	queue              string
	jobOptions         work.JobOptions
	function           func(job *work.Job) error
	productionSchedule bool
}

func handler(job *work.Job) error {
	payload := &model.Payload{}
	p, err := base64.StdEncoding.DecodeString(job.ArgString("payload"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(p, payload)
	if err != nil {
		return err
	}
	return task_handler.DummyTask(context.Background(), payload.TaskID)
}
