package loggo

import (
	"fmt"
	"time"
)

// DefaultLogger default implementation of ILogger interface
type DefaultLogger struct {
	name       string
	handler    IHandler
	processors []IProcessor
}

// New Returns new DefaultLogger instance
func New(name string, handler IHandler) *DefaultLogger {
	return &DefaultLogger{
		name:       name,
		handler:    handler,
		processors: make([]IProcessor, 0),
	}
}

// AddProcessor adds entry processor to logger
func (l *DefaultLogger) AddProcessor(processors ...IProcessor) {
	l.processors = append(l.processors, processors...)
}

// Log logs new entry with specified level
func (l *DefaultLogger) Log(level Level, args ...interface{}) {
	entry := NewEntry(level, time.Now(), fmt.Sprint(args...))
	for _, processor := range l.processors {
		processor.Process(entry)
	}
	entry.Fields["_module"] = l.name
	l.handler.Handle(entry)
}

// Logf logs new entry with specified level
func (l *DefaultLogger) Logf(level Level, format string, args ...interface{}) {
	l.Log(level, fmt.Sprintf(format, args...))
}

// Debug alias for log with debug level
func (l *DefaultLogger) Debug(args ...interface{}) {
	l.Log(LevelDebug, args...)
}

// Info alias for log with Info level
func (l *DefaultLogger) Info(args ...interface{}) {
	l.Log(LevelInfo, args...)
}

// Notice alias for log with notice level
func (l *DefaultLogger) Notice(args ...interface{}) {
	l.Log(LevelNotice, args...)
}

// Warning alias for log with warning level
func (l *DefaultLogger) Warning(args ...interface{}) {
	l.Log(LevelWarning, args...)
}

// Error alias for log with error level
func (l *DefaultLogger) Error(args ...interface{}) {
	l.Log(LevelError, args...)
}

// Critical alias for log with critical level
func (l *DefaultLogger) Critical(args ...interface{}) {
	l.Log(LevelCritical, args...)
}

// Alert alias for log with alert level
func (l *DefaultLogger) Alert(args ...interface{}) {
	l.Log(LevelAlert, args...)
}

// Emergency alias for log with emergency level
func (l *DefaultLogger) Emergency(args ...interface{}) {
	l.Log(LevelEmergency, args...)
}

// Debugf alias for log with debug level
func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	l.Log(LevelDebug, args...)
}

// Infof alias for log with info level
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	l.Log(LevelInfo, args...)
}

// Noticef alias for log with notice level
func (l *DefaultLogger) Noticef(format string, args ...interface{}) {
	l.Log(LevelNotice, args...)
}

// Warningf alias for log with warning level
func (l *DefaultLogger) Warningf(format string, args ...interface{}) {
	l.Log(LevelWarning, args...)
}

// Errorf alias for log with error level
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	l.Log(LevelError, args...)
}

// Criticalf alias for log with critical level
func (l *DefaultLogger) Criticalf(format string, args ...interface{}) {
	l.Log(LevelCritical, args...)
}

// Alertf alias for log with alert level
func (l *DefaultLogger) Alertf(format string, args ...interface{}) {
	l.Log(LevelAlert, args...)
}

// Emergencyf alias for log with emergency level
func (l *DefaultLogger) Emergencyf(format string, args ...interface{}) {
	l.Log(LevelEmergency, args...)
}
