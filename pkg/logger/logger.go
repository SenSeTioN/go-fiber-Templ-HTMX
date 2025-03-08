package logger

import (
	"log/slog"
	"os"

	"sensetion/go-fiber/config"
)

func NewLogger(config *config.LogConfig) *slog.Logger {
	var logger *slog.Logger
	level := slog.Level(config.Level)

	switch config.Format {
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))

	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))
	}

	return logger
}
