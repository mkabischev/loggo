# loggo

Logger inspired by Monolog

Simple usage (writing to stdout)
```go
handler := handlers.NewStreamHandler(loggo.LevelDebug, os.Stdout, &loggo.TextFormatter{})
logger := loggo.New("EXAMPLE", handler)
logger.AddProcessor(&processors.CalleeProcessor{})

logger.Debugf("hello, %s", "user")
logger.Error("some error")

// Output:
// [2015-09-16 12:02:25.289628] 0 0 0 EXAMPLE (DEBUG) loggo/example loggo_test.go:15: hello, user {}
// [2015-09-16 12:02:25.290890] 0 0 0 EXAMPLE (ERROR) loggo/example loggo_test.go:16: some error {}
```

Bufferized output (aka monolog fingercross)
```go
handler := handlers.NewStreamHandler(loggo.LevelDebug, os.Stdout, &loggo.TextFormatter{})
logger := loggo.New("GOB", handlers.NewBufferHandler(handler, loggo.LevelWarning))	logger.AddProcessor(processors.NewCalleeProcessor(), processors.NewFieldsProcessor(map[string]interface{}{"_traceID": 123}))
logger.Debug("hello debug")
logger.Info("hello warning")

// Output is empty because no Warning entry.
```

But 
```go
handler := handlers.NewStreamHandler(loggo.LevelDebug, os.Stdout, &loggo.TextFormatter{})
logger := loggo.New("GOB", handlers.NewBufferHandler(handler, loggo.LevelWarning))
logger.AddProcessor(processors.NewCalleeProcessor(), processors.NewFieldsProcessor(map[string]interface{}{"_traceID": 123}))
logger.Debug("hello debug")
logger.Info("hello warning")
logger.Warning("foo")

// Output:
// [2015-09-16 12:09:20.992744] 7B 0 0 GOB (DEBUG) main main.go:14: hello debug {}
// [2015-09-16 12:09:20.992795] 7B 0 0 GOB (INFO) main main.go:15: hello warning {}
// [2015-09-16 12:09:20.992818] 7B 0 0 GOB (WARNING) main main.go:16: foo {}
```

Roadmap:
- [ ] tests
- [ ] configurable TextFormatter with placeholders
- [ ] support builtin log package interface
