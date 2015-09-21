package loggo

import (
	"fmt"
	"time"
)

var DefaultHandler = NewStreamHandler(LevelDebug, NewTextFormatter("[:time:] (:level:) :message:"))

type ILogger interface {
    Debugf(format string, args ...interface{})

    Infof(format string, args ...interface{})

    Noticef(format string, args ...interface{})

    Warningf(format string, args ...interface{})

    Errorf(format string, args ...interface{})

    Alertf(format string, args ...interface{})

    Criticalf(format string, args ...interface{})

    Emergencyf(format string, args ...interface{})

    Logf(level Level, format string, args ...interface{})

    Debug(args ...interface{})

    Info(args ...interface{})

    Notice(args ...interface{})

    Warning(args ...interface{})

    Error(args ...interface{})

    Alert(args ...interface{})

    Critical(args ...interface{})

    Emergency(args ...interface{})

    Log(level Level, args ...interface{})
}

// Logger is a default implementation of ILogger interface
type Logger struct {
	name       string
	handler    IHandler
	processors []IProcessor
}

// New Returns new Logger instance
func New(name string, handler IHandler) *Logger {
	return &Logger{
		name:       name,
		handler:    handler,
		processors: make([]IProcessor, 0),
	}
}

func (l *Logger) Copy() *Logger {
	return &Logger{
		name:       l.name,
		handler:    l.handler.Copy(),
		processors: l.processors,
	}
}

// AddProcessor adds entry processor to logger
func (l *Logger) AddProcessor(processors ...IProcessor) {
	l.processors = append(l.processors, processors...)
}

// Log logs new entry with specified level
func (l *Logger) Log(level Level, args ...interface{}) {
	entry := NewEntry(level, time.Now(), fmt.Sprint(args...))
	for _, processor := range l.processors {
		processor.Process(entry)
	}
	entry.Fields["_module"] = l.name
	l.handler.Handle(entry)
}

// Logf logs new entry with specified level
func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	l.Log(level, fmt.Sprintf(format, args...))
}

// Debug alias for log with debug level
func (l *Logger) Debug(args ...interface{}) {
	l.Log(LevelDebug, args...)
}

// Info alias for log with Info level
func (l *Logger) Info(args ...interface{}) {
	l.Log(LevelInfo, args...)
}

// Notice alias for log with notice level
func (l *Logger) Notice(args ...interface{}) {
	l.Log(LevelNotice, args...)
}

// Warning alias for log with warning level
func (l *Logger) Warning(args ...interface{}) {
	l.Log(LevelWarning, args...)
}

// Error alias for log with error level
func (l *Logger) Error(args ...interface{}) {
	l.Log(LevelError, args...)
}

// Critical alias for log with critical level
func (l *Logger) Critical(args ...interface{}) {
	l.Log(LevelCritical, args...)
}

// Alert alias for log with alert level
func (l *Logger) Alert(args ...interface{}) {
	l.Log(LevelAlert, args...)
}

// Emergency alias for log with emergency level
func (l *Logger) Emergency(args ...interface{}) {
	l.Log(LevelEmergency, args...)
}

// Debugf alias for log with debug level
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logf(LevelDebug, format, args...)
}

// Infof alias for log with info level
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logf(LevelInfo, format, args...)
}

// Noticef alias for log with notice level
func (l *Logger) Noticef(format string, args ...interface{}) {
	l.Logf(LevelNotice, format, args...)
}

// Warningf alias for log with warning level
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Logf(LevelWarning, format, args...)
}

// Errorf alias for log with error level
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logf(LevelError, format, args...)
}

// Criticalf alias for log with critical level
func (l *Logger) Criticalf(format string, args ...interface{}) {
	l.Logf(LevelCritical, format, args...)
}

// Alertf alias for log with alert level
func (l *Logger) Alertf(format string, args ...interface{}) {
	l.Logf(LevelAlert, format, args...)
}

// Emergencyf alias for log with emergency level
func (l *Logger) Emergencyf(format string, args ...interface{}) {
	l.Logf(LevelEmergency, format, args...)
}
