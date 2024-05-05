package logger

import (
	"io"
	"log"
	"log/slog"
)

func NewSlog(writer io.WriteCloser) *slog.Logger {
	h := slog.NewTextHandler(writer, nil)
	l := slog.New(h)
	return l
}

func NewLog(writer io.WriteCloser) *log.Logger {
	l := log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return l
}
