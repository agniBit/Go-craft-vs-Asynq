package main

import (
	"github.com/agniBit/bench-mark/app/config"
	gocraft_worker "github.com/agniBit/bench-mark/go_craft/worker"
)

func main() {
	config.Load()
	gocraft_worker.RunWorkers(config.Get())
}
