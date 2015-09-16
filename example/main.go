package main

import (
	"loggo"
	"loggo/handlers"
	"loggo/processors"
	"os"
)

func main() {
	sh := handlers.NewStreamHandler(loggo.LevelDebug, os.Stdout, &loggo.TextFormatter{})
	logger := loggo.New("GOB", handlers.NewBufferHandler(sh, loggo.LevelWarning))
	logger.AddProcessor(processors.NewCalleeProcessor(), processors.NewFieldsProcessor(map[string]interface{}{"_traceID":123}))
	logger.Debug("hello debug")
	logger.Info("hello warning")
	logger.Warning("foo")
}
