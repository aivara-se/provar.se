package logger

import (
	"io"
	"log/slog"
	"os"

	"provar.se/webapi/lib/config"
)

var (
	// JSONLogger is a logger that writes JSON logs to the standard output.
	JSONLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// TextLogger is a logger that writes text logs to the standard output.
	TextLogger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	// SilentLogger is a logger that does not write any logs.
	SilentLogger = slog.New(slog.NewTextHandler(io.Discard, nil))
)

// Get returns the logger based on the configuration.
func Get() *slog.Logger {
	logFormat := config.Get().LogFormat
	if logFormat == "json" {
		return JSONLogger
	} else if logFormat == "text" {
		return TextLogger
	} else if logFormat == "file" {
		return NewFileLogger()
	} else {
		return SilentLogger
	}
}

// NewFileLogger creates a new logger that writes text logs to a file.
func NewFileLogger() *slog.Logger {
	logOutput := config.Get().LogOutput
	logFile, err := os.OpenFile(logOutput, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Error opening log file: " + err.Error())
	}
	return slog.New(slog.NewTextHandler(logFile, nil))
}
