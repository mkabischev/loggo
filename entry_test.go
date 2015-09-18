package loggo

import (
	"time"

	. "gopkg.in/check.v1"
)

type EntryTestSuite struct {}

var (
	_ = Suite(&EntryTestSuite{})
)

func (s *EntryTestSuite) TestNewEntry(c *C) {
	now := time.Now()
	entry := NewEntry(LevelWarning, now, "some message")

	c.Assert(entry.Level, Equals, LevelWarning)
	c.Assert(entry.Message, Equals, "some message")
	c.Assert(entry.Time, Equals, now)
	c.Assert(entry.Fields, HasLen, 0)
}
