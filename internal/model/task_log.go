package model

import (
	"strconv"
	"time"
)

const (
	TIMESTAMP_FORMAT = "15:04:05"
)

type TaskLog struct {
	Timestamp   time.Time
	Description string
	Entry       LogEntry
	PID         int
}

func NewTaskLogFromStrings(
	timestamp,
	description,
	entry,
	pid string,
) (tl TaskLog, err error) {
	ts, err := time.Parse(TIMESTAMP_FORMAT, timestamp)
	if err != nil {
		return tl, err
	}

	id, err := strconv.Atoi(pid)
	if err != nil {
		return tl, err
	}

	ent, err := ToLogEntry(entry)
	if err != nil {
		return tl, err
	}

	tl.Timestamp = ts
	tl.Description = description
	tl.Entry = ent
	tl.PID = id

	return
}
