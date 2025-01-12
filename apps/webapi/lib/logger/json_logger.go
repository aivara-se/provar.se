package logger

import (
	"log/slog"
	"os"
)

// JSONLogger is a logger that writes JSON logs to the standard output.
var JSONLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
