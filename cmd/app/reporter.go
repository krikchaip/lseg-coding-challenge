package main

import (
	"log"

	"github.com/krikchaip/lseg-coding-challenge/internal/model"
)

var consoleReporter = &ConsoleReporter{}

type ConsoleReporter struct{}

func (this *ConsoleReporter) Warn(task model.Task, mins float64) {
	log.Printf("[WARNING] Task %d took %.0fmins to complete.", task.Pid, mins)
}

func (this *ConsoleReporter) Error(task model.Task, mins float64) {
	log.Printf("[ERROR] Task %d took %.0fmins to complete.", task.Pid, mins)
}
