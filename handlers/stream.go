package handlers

import (
	"io"
	"loggo"
)

type StreamHandler struct {
	level     loggo.Level
	out       io.Writer
	formatter loggo.IFormatter
}

func NewStreamHandler(level loggo.Level, out io.Writer, formatter loggo.IFormatter) *StreamHandler {
	return &StreamHandler{
		level:     level,
		out:       out,
		formatter: formatter,
	}
}

func (h *StreamHandler) Handle(entry *loggo.Entry) error {
	if entry.Level < h.level {
		return nil
	}

	if _, err := h.out.Write(h.formatter.Format(entry)); err != nil {
		return err
	}

	return nil
}
