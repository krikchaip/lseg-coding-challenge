package model

import "fmt"

type LogEntry string

const (
	EntryStart LogEntry = "START"
	EntryEnd   LogEntry = "END"
)

func ToLogEntry(s string) (LogEntry, error) {
	switch LogEntry(s) {
	case EntryStart:
		return EntryStart, nil
	case EntryEnd:
		return EntryEnd, nil
	default:
		return "", fmt.Errorf("Unknown emtry %q", s)
	}
}
