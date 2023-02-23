package main

import (
	"github.com/agniBit/benchmark/app/config"
	gocraft_worker "github.com/agniBit/benchmark/go_craft/worker"
)

func main() {
	config.Load()
	gocraft_worker.RunWorkers(config.Get())
}
