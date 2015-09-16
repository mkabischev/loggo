package handlers

import "loggo"

type MultiHandler struct {
	handlers []loggo.IHandler
}

func NewMultiHandler(handlers ...loggo.IHandler) *MultiHandler {
	return &MultiHandler{
		handlers: handlers,
	}
}

func (h *MultiHandler) Handle(entry *loggo.Entry) error {
	for _, handler := range h.handlers {
		if err := handler.Handle(entry); err != nil {
			return err
		}
	}

	return nil
}
