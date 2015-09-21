package loggo

import . "gopkg.in/check.v1"

type LoggerTestSuite struct {
	handler *testHandler
	logger  *Logger
}

var (
	_ = Suite(&LoggerTestSuite{})
)

func (s *LoggerTestSuite) SetUpTest(c *C) {
	s.handler = &testHandler{}
	s.logger = New("test", s.handler)
}

func (s *LoggerTestSuite) TestLog(c *C) {
	s.logger.Log(LevelDebug, "hello")
	c.Assert(s.handler.entries, HasLen, 1)
	s.assertEntry(c, 0, LevelDebug, "hello")
}

func (s *LoggerTestSuite) TestDebug(c *C) {
	s.logger.Debug("hello")
	s.assertSingleEntry(c, LevelDebug, "hello")
}

func (s *LoggerTestSuite) TestInfo(c *C) {
	s.logger.Info("hello")
	s.assertSingleEntry(c, LevelInfo, "hello")
}

func (s *LoggerTestSuite) TestNotice(c *C) {
	s.logger.Notice("hello")
	s.assertSingleEntry(c, LevelNotice, "hello")
}

func (s *LoggerTestSuite) TestWarning(c *C) {
	s.logger.Warning("hello")
	s.assertSingleEntry(c, LevelWarning, "hello")
}

func (s *LoggerTestSuite) TestError(c *C) {
	s.logger.Error("hello")
	s.assertSingleEntry(c, LevelError, "hello")
}

func (s *LoggerTestSuite) TestCritical(c *C) {
	s.logger.Critical("hello")
	s.assertSingleEntry(c, LevelCritical, "hello")
}

func (s *LoggerTestSuite) TestAlert(c *C) {
	s.logger.Alert("hello")
	s.assertSingleEntry(c, LevelAlert, "hello")
}

func (s *LoggerTestSuite) TestEmergency(c *C) {
	s.logger.Emergency("hello")
	s.assertSingleEntry(c, LevelEmergency, "hello")
}

func (s *LoggerTestSuite) TestLogf(c *C) {
	s.logger.Logf(LevelDebug, "hello, %s", "man")
	c.Assert(s.handler.entries, HasLen, 1)
	s.assertEntry(c, 0, LevelDebug, "hello, man")
}

func (s *LoggerTestSuite) TestDebugf(c *C) {
	s.logger.Debugf("hello, %s", "man")
	s.assertSingleEntry(c, LevelDebug, "hello, man")
}

func (s *LoggerTestSuite) TestInfof(c *C) {
	s.logger.Infof("hello, %s", "man")
	s.assertSingleEntry(c, LevelInfo, "hello, man")
}

func (s *LoggerTestSuite) TestNoticef(c *C) {
	s.logger.Noticef("hello, %s", "man")
	s.assertSingleEntry(c, LevelNotice, "hello, man")
}

func (s *LoggerTestSuite) TestWarningf(c *C) {
	s.logger.Warningf("hello, %s", "man")
	s.assertSingleEntry(c, LevelWarning, "hello, man")
}

func (s *LoggerTestSuite) TestErrorf(c *C) {
	s.logger.Errorf("hello, %s", "man")
	s.assertSingleEntry(c, LevelError, "hello, man")
}

func (s *LoggerTestSuite) TestCriticalf(c *C) {
	s.logger.Criticalf("hello, %s", "man")
	s.assertSingleEntry(c, LevelCritical, "hello, man")
}

func (s *LoggerTestSuite) TestAlertf(c *C) {
	s.logger.Alertf("hello, %s", "man")
	s.assertSingleEntry(c, LevelAlert, "hello, man")
}

func (s *LoggerTestSuite) TestEmergencyf(c *C) {
	s.logger.Emergencyf("hello, %s", "man")
	s.assertSingleEntry(c, LevelEmergency, "hello, man")
}

func (s *LoggerTestSuite) TestAddProcessor(c *C) {
	processor1 := &testProcessor{}
	processor2 := &testProcessor{}
	s.logger.AddProcessor(processor1, processor2)

	c.Assert(s.logger.processors, HasLen, 2)
	c.Assert(s.logger.processors[0], Equals, processor1)
	c.Assert(s.logger.processors[1], Equals, processor2)
}

func (s *LoggerTestSuite) TestCopy(c *C) {
	handler := &handlerForCopy{}
	processor := &testProcessor{}
	logger := New("test", handler)
	logger.AddProcessor(processor)

	copy := logger.Copy()
	c.Assert(copy, Not(Equals), logger)
	c.Assert(copy.name, Equals, logger.name)
	c.Assert(copy.handler.(*handlerForCopy).original, Equals, handler)
	c.Assert(copy.processors[0], Equals, logger.processors[0])
}

func (s *LoggerTestSuite) assertSingleEntry(c *C, level Level, message string) {
	c.Assert(s.handler.entries, HasLen, 1)
	s.assertEntry(c, 0, level, message)
}

func (s *LoggerTestSuite) assertEntry(c *C, entryIndex int, level Level, message string) {
	entry := s.handler.entries[entryIndex]
	c.Assert(entry.Message, Equals, message)
	c.Assert(entry.Level, Equals, level)
}
