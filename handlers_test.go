package loggo

import (
	"bytes"
	"time"

	. "gopkg.in/check.v1"
)

type HandlersTestSuite struct{
	rawFormatter IFormatter
}

var (
	_ = Suite(&HandlersTestSuite{})
)

func (h *HandlersTestSuite) SetUpSuite(c *C) {
	h.rawFormatter = NewTextFormatter(":message:")
}


func (s *HandlersTestSuite) TestStreamHandlerHandle(c *C) {
	buf := &bytes.Buffer{}
	handler := NewStreamHandler(LevelDebug, buf, s.rawFormatter)
	handler.Handle(NewEntry(LevelDebug, time.Now(), "hello"))
	handler.Handle(NewEntry(LevelInfo, time.Now(), "man"))

	c.Assert(buf.String(), Equals, "hello\nman\n")
}

func (s *HandlersTestSuite) TestStreamHandlerHandleLowLevel(c *C) {
	buf := &bytes.Buffer{}
	handler := NewStreamHandler(LevelInfo, buf, s.rawFormatter)
	handler.Handle(NewEntry(LevelDebug, time.Now(), "hello"))
	handler.Handle(NewEntry(LevelInfo, time.Now(), "man"))

	c.Assert(buf.String(), Equals, "man\n")
}

func (s *HandlersTestSuite) TestStreamHandlerCopy(c *C) {
	buf := &bytes.Buffer{}
	handler := NewStreamHandler(LevelInfo, buf, s.rawFormatter)

	c.Assert(handler.Copy(), Equals, handler)
}

func (s *HandlersTestSuite) TestBufferHandlerHandle(c *C) {
	buf := &bytes.Buffer{}
	streamHandler := NewStreamHandler(LevelDebug, buf, s.rawFormatter)

	handler := NewBufferHandler(streamHandler, LevelWarning)
	handler.Handle(NewEntry(LevelDebug, time.Now(), "debug"))
	handler.Handle(NewEntry(LevelInfo, time.Now(), "info"))
	handler.Handle(NewEntry(LevelWarning, time.Now(), "warning"))

	c.Assert(buf.String(), Equals, "debug\ninfo\nwarning\n")
}

func (s *HandlersTestSuite) TestBufferHandlerHandleLowLevel(c *C) {
	buf := &bytes.Buffer{}
	streamHandler := NewStreamHandler(LevelDebug, buf, s.rawFormatter)

	handler := NewBufferHandler(streamHandler, LevelWarning)
	handler.Handle(NewEntry(LevelDebug, time.Now(), "debug"))
	handler.Handle(NewEntry(LevelInfo, time.Now(), "info"))

	c.Assert(buf.String(), Equals, "")
}

func (s *HandlersTestSuite) TestBufferHandlerCopy(c *C) {
	handler := &handlerForCopy{}
	bufferHandler := NewBufferHandler(handler, LevelWarning)
	copy := bufferHandler.Copy().(*BufferHandler)
	c.Assert(copy, Not(Equals), bufferHandler)
	c.Assert(copy.level, Equals, bufferHandler.level)
	c.Assert(copy.handler.(*handlerForCopy).original, Equals, handler)
}

func (s *HandlersTestSuite) TestMultiHandlerHandle(c *C) {
	buf := &bytes.Buffer{}
	streamHandler := NewStreamHandler(LevelDebug, buf, s.rawFormatter)

	buf2 := &bytes.Buffer{}
	streamHandler2 := NewStreamHandler(LevelInfo, buf2, s.rawFormatter)

	handler := NewMultiHandler(streamHandler, streamHandler2)
	handler.Handle(NewEntry(LevelDebug, time.Now(), "debug"))
	handler.Handle(NewEntry(LevelInfo, time.Now(), "info"))

	c.Assert(buf.String(), Equals, "debug\ninfo\n")
	c.Assert(buf2.String(), Equals, "info\n")
}

func (s *HandlersTestSuite) TestMultiHandlerCopy(c *C) {
	handler := &handlerForCopy{}
	handler2 := &handlerForCopy{}

	multiHandler := NewMultiHandler(handler, handler2)
	copy := multiHandler.Copy().(*MultiHandler)

	c.Assert(copy, Not(Equals), multiHandler)
	c.Assert(copy.handlers[0].(*handlerForCopy).original, Equals, handler)
	c.Assert(copy.handlers[1].(*handlerForCopy).original, Equals, handler2)
}
