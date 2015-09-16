package loggo

import "time"

// Entry is a struct for logging unit
type Entry struct {
	Fields  map[string]interface{}
	Level   Level
	Message string
	Time    time.Time
}

// NewEntry constructor for entry
func NewEntry(level Level, time time.Time, message string) *Entry {
	return &Entry{
		Level:   level,
		Time:    time,
		Message: message,
		Fields:  make(map[string]interface{}),
	}
}
