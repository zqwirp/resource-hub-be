package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

func NewSlog(writer *io.WriteCloser) *slog.Logger {
	if writer == nil {
		writer = os.Stdout // Default to standard output if no writer is provided
	} else {
		writer = io.MultiWriter(os.Stdout, writer)
	}

	// Initialize a new text handler with the writer.
	// Adjust the second argument to NewTextHandler if it requires specific configuration or error handling.
	h := slog.NewTextHandler(*writer, nil)
	l := slog.New(h)
	return l
}

func NewLog(file *os.File) *log.Logger {
	l := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return l
}
