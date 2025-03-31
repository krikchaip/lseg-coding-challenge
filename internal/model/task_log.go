package model

import "time"

type TaskLog struct {
	Timestamp   time.Time
	Description string
	Entry       LogEntry
	PID         int
}
