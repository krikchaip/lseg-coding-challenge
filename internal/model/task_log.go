package model

import "time"

const (
	TIMESTAMP_FORMAT = "15:04:05"
)

type TaskLog struct {
	Timestamp   time.Time
	Description string
	Entry       LogEntry
	PID         int
}
