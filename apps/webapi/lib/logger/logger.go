package logger

import (
	"log/slog"

	"provar.se/webapi/lib/config"
)

// Get returns the logger based on the configuration.
func Get() *slog.Logger {
	if config.Get().LogFormat == "json" {
		return JSONLogger
	} else {
		return TextLogger
	}
}
