package logging

import (
	"log/slog"
	"os"
)

// Logger is the global logger instance
var Logger *slog.Logger

// Init initializes the global logger with a default configuration.
// By default, it logs to stdout with text format and INFO level.
func Init() {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// SetLevel sets the logging level (DEBUG, INFO, WARN, ERROR).
func SetLevel(level slog.Level) {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
}

// Debug logs a debug message with optional key-value pairs.
func Debug(msg string, args ...interface{}) {
	Logger.Debug(msg, args...)
}

// Info logs an informational message with optional key-value pairs.
func Info(msg string, args ...interface{}) {
	Logger.Info(msg, args...)
}

// Warn logs a warning message with optional key-value pairs.
func Warn(msg string, args ...interface{}) {
	Logger.Warn(msg, args...)
}

// Error logs an error message with optional key-value pairs.
func Error(msg string, args ...interface{}) {
	Logger.Error(msg, args...)
}
