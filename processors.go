package loggo

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type IProcessor interface {
	Process(entry *Entry)
}

// CalleeProcessor adds packageName and filename to entry fields
type CalleeProcessor struct{}

// NewCalleeProcessor constructor for CalleeProcessor
func NewCalleeProcessor() *CalleeProcessor {
	return &CalleeProcessor{}
}

// Process adds two fields to entry:
// _package - name of package where logger was called
// _file - file:line where logger was called
func (p *CalleeProcessor) Process(entry *Entry) {
	entry.Fields["_package"] = getPackageName()
	entry.Fields["_file"] = getFileName()
}

func getPackageName() string {
	v := "???"
	if pc, _, _, ok := runtime.Caller(4); ok {
		if f := runtime.FuncForPC(pc); f != nil {
			v = formatFuncName(f.Name())
		}
	}

	return v
}

func getFileName() string {
	_, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)
	}

	return fmt.Sprintf("%s:%d", file, line)
}

func formatFuncName(f string) string {
	i := strings.LastIndex(f, "/")
	j := strings.Index(f[i+1:], ".")
	if j < 1 {
		return "???"
	}

	return f[:i+j+1]
}

type FieldsProcessor struct {
	fields map[string]interface{}
}

func NewFieldsProcessor(fields map[string]interface{}) *FieldsProcessor {
	return &FieldsProcessor{
		fields: fields,
	}
}

func (p *FieldsProcessor) Process(entry *Entry) {
	for key, value := range p.fields {
		entry.Fields[key] = value
	}
}
