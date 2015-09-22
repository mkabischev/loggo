package loggo

import (
	"io"
	"os"
	"sync"
)

// IHandler interface
type IHandler interface {
	Handle(entry *Entry)
	Copy() IHandler
	IsEnabledFor(level Level) bool
}

type BufferHandler struct {
	handler    IHandler
	flushLevel Level
	buffer     []*Entry
	lock       sync.Mutex
}

func NewBufferHandler(handler IHandler, flushLevel Level) *BufferHandler {
	return &BufferHandler{
		handler:    handler,
		flushLevel: flushLevel,
		buffer:     make([]*Entry, 0, 128),
	}
}

func (h *BufferHandler) Handle(entry *Entry) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.buffer = append(h.buffer, entry)
	if entry.Level >= h.flushLevel {
		for _, e := range h.buffer {
			h.handler.Handle(e)
		}

		h.buffer = make([]*Entry, 0, 128)
	}
}

func (h *BufferHandler) Copy() IHandler {
	return NewBufferHandler(h.handler.Copy(), h.flushLevel)
}

func (h *BufferHandler) IsEnabledFor(level Level) bool {
	return h.handler.IsEnabledFor(level)
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

func (h *MultiHandler) IsEnabledFor(level Level) bool {
	for _, handler := range h.handlers {
		if handler.IsEnabledFor(level) {
			return true
		}
	}

	return false
}

type StreamHandler struct {
	level        Level
	outputWriter io.Writer
	formatter    IFormatter
}

// NewStreamHandler returns new StreamHandler.
// If out is not passed - stdout will be used
func NewStreamHandler(level Level, formatter IFormatter, outputWriter ...io.Writer) *StreamHandler {
	var outputW io.Writer

	if len(outputWriter) > 1 {
		panic("You can't pass more than one outputWriter")
	}

	if len(outputWriter) == 1 {
		outputW = outputWriter[0]
	} else {
		outputW = os.Stdout
	}

	return &StreamHandler{
		level:        level,
		outputWriter: outputW,
		formatter:    formatter,
	}
}

func (h *StreamHandler) Handle(entry *Entry) {
	if entry.Level < h.level {
		return
	}

	h.outputWriter.Write(h.formatter.Format(entry))
}

func (h *StreamHandler) Copy() IHandler {
	return h
}

func (h *StreamHandler) IsEnabledFor(level Level) bool {
	return level >= h.level
}
