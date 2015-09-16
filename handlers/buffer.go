package handlers

import (
	"github.com/mkabischev/loggo"
	"sync"
)

type BufferHandler struct {
	handler loggo.IHandler
	level   loggo.Level
	buffer  []*loggo.Entry
	lock    sync.Mutex
}

func NewBufferHandler(handler loggo.IHandler, level loggo.Level) *BufferHandler {
	return &BufferHandler{
		handler: handler,
		level:   level,
		buffer:  make([]*loggo.Entry, 0, 128),
	}
}

func (h *BufferHandler) Handle(entry *loggo.Entry) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.buffer = append(h.buffer, entry)
	if entry.Level >= h.level {
		for _, e := range h.buffer {
			h.handler.Handle(e)
		}

		h.buffer = make([]*loggo.Entry, 0, 128)
	}
}

func (h *BufferHandler) Copy() loggo.IHandler {
	return NewBufferHandler(h.handler.Copy(), h.level)
}
