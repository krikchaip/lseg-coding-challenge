package core_test

import (
	"testing"
	"time"

	"github.com/krikchaip/lseg-coding-challenge/internal/core"
	"github.com/krikchaip/lseg-coding-challenge/internal/model"
)

func TestTaskMonitor(t *testing.T) {
	t.Run("Don't report if a job took <= 5mins", func(t *testing.T) {
		var reporter TestReporter

		timestamp := toTimestamp("00:00:00")
		monitor := core.NewTaskMonitor(&reporter)

		monitor.AppendLog(model.TaskLog{
			PID:       1,
			Entry:     model.EntryStart,
			Timestamp: timestamp,
		})
		monitor.AppendLog(model.TaskLog{
			PID:       1,
			Entry:     model.EntryEnd,
			Timestamp: timestamp.Add(5 * time.Minute),
		})

		if reporter.WarnCalled || reporter.ErrorCalled {
			t.Error("expect reporter not to be called.")
		}
	})

	t.Run("Warn if a job took > 5mins, but <= 10mins", func(t *testing.T) {
		var reporter TestReporter

		timestamp := toTimestamp("00:00:00")
		monitor := core.NewTaskMonitor(&reporter)

		monitor.AppendLog(model.TaskLog{
			PID:       1,
			Entry:     model.EntryStart,
			Timestamp: timestamp,
		})
		monitor.AppendLog(model.TaskLog{
			PID:       1,
			Entry:     model.EntryEnd,
			Timestamp: timestamp.Add(10 * time.Minute),
		})

		if !reporter.WarnCalled {
			t.Error("expect reporter.Warn() to be called.")
		}

		if reporter.ErrorCalled {
			t.Error("expect reporter.Error() not to be called.")
		}
	})
}

func toTimestamp(t string) time.Time {
	ts, _ := time.Parse(model.TIMESTAMP_FORMAT, t)
	return ts
}

type TestReporter struct {
	WarnCalled  bool
	ErrorCalled bool
}

func (this *TestReporter) Warn(_ model.Task, _ float64) {
	this.WarnCalled = true
}

func (this *TestReporter) Error(_ model.Task, _ float64) {
	this.ErrorCalled = true
}
