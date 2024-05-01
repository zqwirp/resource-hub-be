package logger

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
	*os.File
}

func New(path, name string) (*Logger, error) {
	err := os.MkdirAll("./"+path, 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path+"/"+name+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	h := slog.NewTextHandler(io.MultiWriter(os.Stdout, file), nil)

	l := slog.New(h)
	return &Logger{Logger: l, File: file}, nil
}
