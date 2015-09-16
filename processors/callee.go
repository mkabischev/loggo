package processors

import (
	"fmt"
	"loggo"
	"path/filepath"
	"runtime"
	"strings"
)

type CalleeProcessor struct{}

func NewCalleeProcessor() *CalleeProcessor {
	return &CalleeProcessor{}
}

func (p *CalleeProcessor) Process(entry *loggo.Entry) {
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
