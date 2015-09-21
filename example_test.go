package loggo

func ExampleSimpleUsage() {
	logger := New("MyLogger", NewStreamHandler(LevelDebug, NewTextFormatter("(:level:) :message:")))
	logger.Debug("hello debug")
	logger.Info("hello info")

	// Output:
	// (DEBUG) hello debug
	// (INFO) hello info
}

func ExampleBufferEmpty() {
	handler :=  NewStreamHandler(LevelDebug, NewTextFormatter("(:level:) :message:"))
	logger := New("MyLogger", NewBufferHandler(handler, LevelWarning))
	logger.Debug("hello debug")
	logger.Info("hello info")

	// Output:
}

func ExampleBuffer() {
	handler :=  NewStreamHandler(LevelDebug, NewTextFormatter("(:level:) :message:"))
	logger := New("MyLogger", NewBufferHandler(handler, LevelWarning))
	logger.Debug("hello debug")
	logger.Info("hello info")
	logger.Warning("hello warning")

	// Output:
	// (DEBUG) hello debug
	// (INFO) hello info
	// (WARNING) hello warning
}
