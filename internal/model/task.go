package model

import "time"

type Task struct {
	Description string
	Pid         int
	StartedAt   time.Time
	EndedAt     time.Time
}

func NewTaskFromLog(tl TaskLog) (t Task) {
	t.Description = tl.Description
	t.Pid = tl.PID

	switch tl.Entry {
	case EntryStart:
		t.StartedAt = tl.Timestamp
	case EntryEnd:
		t.EndedAt = tl.Timestamp
	}

	return
}
