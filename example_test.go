package loggo

import "os"

func ExampleSimpleUsage() {
	logger := New("GOB", NewStreamHandler(LevelDebug, os.Stdout, NewTextFormatter("(:level:) :message:")))
	logger.Debug("hello debug")
	logger.Info("hello info")

	// Output:
	// (DEBUG) hello debug
	// (INFO) hello info
}

func ExampleBufferEmpty() {
	handler :=  NewStreamHandler(LevelDebug, os.Stdout, NewTextFormatter("(:level:) :message:"))
	logger := New("GOB", NewBufferHandler(handler, LevelWarning))
	logger.Debug("hello debug")
	logger.Info("hello info")

	// Output:
}

func ExampleBuffer() {
	handler :=  NewStreamHandler(LevelDebug, os.Stdout, NewTextFormatter("(:level:) :message:"))
	logger := New("GOB", NewBufferHandler(handler, LevelWarning))
	logger.Debug("hello debug")
	logger.Info("hello info")
	logger.Warning("hello warning")

	// Output:
	// (DEBUG) hello debug
	// (INFO) hello info
	// (WARNING) hello warning
}
