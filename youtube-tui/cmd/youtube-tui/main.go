// Command youtube-tui provides a terminal user interface for searching YouTube videos
package main

import (
	"fmt"
	"os"

	"youtube-tui/internal/config"
	"youtube-tui/internal/logging"
	"youtube-tui/internal/player"
	"youtube-tui/internal/ui"
)

func main() {
	// Initialize logging with configuration from environment
	provider := config.NewProvider()
	logging.Init(
		provider.GetLogFile(),
		provider.GetLogToConsole(),
		provider.GetLogFileMaxSize(),
		provider.GetLogFileMaxBackups(),
	)
	logging.Info("starting youtube-tui application", "logFile", provider.GetLogFile())

	// Apply log level from environment if set
	if levelStr := os.Getenv("YOUTUBE_TUI_LOG_LEVEL"); levelStr != "" {
		logging.SetLevelFromEnv(levelStr)
		logging.Debug("log level set from environment", "level", levelStr)
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
