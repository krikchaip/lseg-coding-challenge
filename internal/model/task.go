package model

import "time"

type Task struct {
	Description string
	Pid         int
	StartedAt   time.Time
	EndedAt     time.Time
}
