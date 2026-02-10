// Command youtube-tui provides a terminal user interface for searching YouTube videos
package main

import (
	"fmt"
	"os"
	"log/slog"

	"youtube-tui/internal/logging"
	"youtube-tui/internal/player"
	"youtube-tui/internal/ui"
)

func main() {
	// Initialize logging
	logging.Init()
	logging.Info("starting youtube-tui application")

	// Apply log level from environment if set
	if levelStr := os.Getenv("YOUTUBE_TUI_LOG_LEVEL"); levelStr != "" {
		var level slog.Level
		switch levelStr {
		case "debug", "DEBUG":
			level = slog.LevelDebug
		case "info", "INFO":
			level = slog.LevelInfo
		case "warn", "WARN":
			level = slog.LevelWarn
		case "error", "ERROR":
			level = slog.LevelError
		default:
			logging.Warn("unknown log level, using INFO", "level", levelStr)
		}
		logging.SetLevel(level)
		logging.Debug("log level set", "level", level)
	}

	// Initialize player with default configuration
	mediaPlayer := player.NewPlayer(nil)
	logging.Debug("player initialized", "executable", os.Getenv("YOUTUBE_TUI_PLAYER_EXECUTABLE"))

	// Run the TUI with the player
	if err := ui.Run(mediaPlayer); err != nil {
		logging.Error("application error", "error", err)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	logging.Info("application exited normally")
}

