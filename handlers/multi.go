package handlers

import "github.com/mkabischev/loggo"

type MultiHandler struct {
	handlers []loggo.IHandler
}

func NewMultiHandler(handlers ...loggo.IHandler) *MultiHandler {
	return &MultiHandler{
		handlers: handlers,
	}
}

func (h *MultiHandler) Handle(entry *loggo.Entry) {
	for _, handler := range h.handlers {
		handler.Handle(entry)
	}
}

func (h *MultiHandler) Copy() loggo.IHandler {
	handlers := make([]loggo.IHandler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.Copy()
	}

	return NewMultiHandler(handlers...)
}
