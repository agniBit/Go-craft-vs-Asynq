package asynq_worker

import (
	"context"
	"encoding/json"
	"github.com/agniBit/bench-mark/app/task_handler"
	"github.com/hibiken/asynq"

	"github.com/agniBit/bench-mark/model"
)

func handleFunc(ctx context.Context, t *asynq.Task) error {
	var payload model.Payload
	err := json.Unmarshal(t.Payload(), &payload)
	if err != nil {
		return err
	}
	return task_handler.DummyTask(ctx, payload.TaskID)
}
