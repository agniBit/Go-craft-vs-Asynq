package main

import (
	"github.com/agniBit/bench-mark/app/config"
	asynqworker "github.com/agniBit/bench-mark/asynq/worker"
)

func main() {
	config.Load()
	asynqworker.StartServer(config.Get())
}
