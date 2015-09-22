package loggo

import . "gopkg.in/check.v1"

type LevelTestSuite struct {
}

var (
	_ = Suite(&LevelTestSuite{})
)

func (s *LevelTestSuite) TestString(c *C) {
	levels := map[Level]string{
		LevelDebug:     "DEBUG",
		LevelInfo:      "INFO",
		LevelNotice:    "NOTICE",
		LevelWarning:   "WARNING",
		LevelError:     "ERROR",
		LevelCritical:  "CRITICAL",
		LevelAlert:     "ALERT",
		LevelEmergency: "EMERGENCY",
		Level(10):      "UNKNOWN",
	}

	for level, name := range levels {
		c.Assert(level.String(), Equals, name)
	}
}
