package model

type Payload struct {
	TaskID string `json:"task_id"`
}

type QueuePriority string

const (
	QueuePriorityCritical QueuePriority = "critical"
	QueuePriorityDefault  QueuePriority = "default"
	QueuePriorityLow      QueuePriority = "low"
)
