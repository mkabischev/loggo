package loggo

type Level uint8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelAlert
	LevelEmergency
)

func (level Level) String() string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelNotice:
		return "NOTICE"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelCritical:
		return "CRITICAL"
	case LevelAlert:
		return "ALERT"
	case LevelEmergency:
		return "EMERGENCY"
	}

	return "UNKNOWN"
}
