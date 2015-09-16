package handlers

import (
	"loggo"
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

func (h *BufferHandler) Handle(entry *loggo.Entry) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.buffer = append(h.buffer, entry)
	if entry.Level >= h.level {
		for _, e := range h.buffer {
			if err := h.handler.Handle(e); err != nil {
				return err
			}
		}

		h.buffer = make([]*loggo.Entry, 0, 128)
	}

	return nil
}
