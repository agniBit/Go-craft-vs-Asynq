package main

import (
	"github.com/agniBit/benchmark/app/config"
	asynqworker "github.com/agniBit/benchmark/asynq/worker"
)

func main() {
	config.Load()
	asynqworker.StartServer(config.Get())
}
