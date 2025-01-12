package logger

import (
	"log/slog"
	"os"
)

// TextLogger is a logger that writes text logs to the standard output.
var TextLogger = slog.New(slog.NewTextHandler(os.Stdout, nil))
