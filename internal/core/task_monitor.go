package core

import (
	"fmt"

	"github.com/krikchaip/lseg-coding-challenge/internal/model"
)

type TaskMonitor struct {
	reporter model.TaskReporter

	pendingEntries map[int]model.TaskLog
}

func NewTaskMonitor(reporter model.TaskReporter) *TaskMonitor {
	return &TaskMonitor{
		reporter: reporter,

		pendingEntries: make(map[int]model.TaskLog),
	}
}

func (this *TaskMonitor) AppendLog(tl model.TaskLog) error {
	switch tl.Entry {
	case model.EntryStart:
		return this.addEntry(tl)
	case model.EntryEnd:
		task, err := this.closeEntry(tl)
		if err != nil {
			return err
		}

		this.reportTask(task)
	}

	return nil
}

func (this *TaskMonitor) addEntry(tl model.TaskLog) error {
	if this.pendingEntries == nil {
		this.pendingEntries = make(map[int]model.TaskLog)
	}

	entry, ok := this.pendingEntries[tl.PID]

	if !ok {
		this.pendingEntries[tl.PID] = tl
		return nil
	}

	return fmt.Errorf("START entry for task %d is already existed.\n", entry.PID)
}

func (this *TaskMonitor) closeEntry(tl model.TaskLog) (model.Task, error) {
	var task model.Task

	entry, ok := this.pendingEntries[tl.PID]

	if !ok {
		return task, fmt.Errorf("START entry for task %d is missing.\n", entry.PID)
	}

	task.Description = entry.Description
	task.Pid = entry.PID
	task.StartedAt = entry.Timestamp
	task.EndedAt = tl.Timestamp

	delete(this.pendingEntries, entry.PID)

	return task, nil
}

func (this *TaskMonitor) reportTask(t model.Task) {
	switch minsTaken := t.EndedAt.Sub(t.StartedAt).Minutes(); {
	case minsTaken <= 5:
		return
	case 5 < minsTaken && minsTaken <= 10:
		this.reporter.Warn(t, minsTaken)
	case minsTaken > 10:
		this.reporter.Error(t, minsTaken)
	}
}
