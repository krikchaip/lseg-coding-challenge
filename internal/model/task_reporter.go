package model

type TaskReporter interface {
	Warn(t Task, mins float64)
	Error(t Task, mins float64)
}
