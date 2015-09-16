package loggo

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// IFormatter formatters convert entries to []byte and used by handlers
type IFormatter interface {
	Format(entry *Entry) []byte
}

// TextFormatter simple formatter
type TextFormatter struct{}

// Format formats entry to common format
func (f *TextFormatter) Format(entry *Entry) []byte {
	b := &bytes.Buffer{}
	time := entry.Time.Format("2006-01-02 15:04:05.000000")
	traceID := f.getEntryField(entry, "_traceID", 0)
	parentSpanID := f.getEntryField(entry, "_parentSpanID", 0)
	spanID := f.getEntryField(entry, "_spanID", 0)
	module := f.getEntryField(entry, "_module", "???")
	level := entry.Level
	pkg := f.getEntryField(entry, "_package", "???")
	file := f.getEntryField(entry, "_file", "???")

	fmt.Fprintf(b, "[%s] %X %X %X %s (%s) %s %s: %s ", time, traceID, parentSpanID, spanID, module, level, pkg, file, entry.Message)

	data := f.filterEntryFields(entry)
	if marshaledData, err := json.Marshal(data); err == nil {
		b.Write(marshaledData)
	}

	b.WriteByte('\n')

	return b.Bytes()
}

func (f *TextFormatter) filterEntryFields(entry *Entry) map[string]interface{} {
	result := make(map[string]interface{}, len(entry.Fields))

	for key, value := range entry.Fields {
		if key[0] != '_' {
			result[key] = value
		}
	}

	return result
}

func (f *TextFormatter) getEntryField(entry *Entry, field string, defaultValue interface{}) interface{} {
	if value, ok := entry.Fields[field]; ok {
		return value
	}

	return defaultValue
}
