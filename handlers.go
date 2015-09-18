package loggo

import (
	"io"
	"sync"
)

// IHandler interface
type IHandler interface {
	Handle(entry *Entry)
	Copy() IHandler
}

type BufferHandler struct {
	handler IHandler
	level   Level
	buffer  []*Entry
	lock    sync.Mutex
}

func NewBufferHandler(handler IHandler, level Level) *BufferHandler {
	return &BufferHandler{
		handler: handler,
		level:   level,
		buffer:  make([]*Entry, 0, 128),
	}
}

func (h *BufferHandler) Handle(entry *Entry) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.buffer = append(h.buffer, entry)
	if entry.Level >= h.level {
		for _, e := range h.buffer {
			h.handler.Handle(e)
		}

		h.buffer = make([]*Entry, 0, 128)
	}
}

func (h *BufferHandler) Copy() IHandler {
	return NewBufferHandler(h.handler.Copy(), h.level)
}

type MultiHandler struct {
	handlers []IHandler
}

func NewMultiHandler(handlers ...IHandler) *MultiHandler {
	return &MultiHandler{
		handlers: handlers,
	}
}

func (h *MultiHandler) Handle(entry *Entry) {
	for _, handler := range h.handlers {
		handler.Handle(entry)
	}
}

func (h *MultiHandler) Copy() IHandler {
	handlers := make([]IHandler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.Copy()
	}

	return NewMultiHandler(handlers...)
}

type StreamHandler struct {
	level     Level
	out       io.Writer
	formatter IFormatter
}

func NewStreamHandler(level Level, out io.Writer, formatter IFormatter) *StreamHandler {
	return &StreamHandler{
		level:     level,
		out:       out,
		formatter: formatter,
	}
}

func (h *StreamHandler) Handle(entry *Entry) {
	if entry.Level < h.level {
		return
	}

	h.out.Write(h.formatter.Format(entry))
}

func (h *StreamHandler) Copy() IHandler {
	return h
}
