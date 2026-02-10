// Package interfaces defines core abstractions used throughout the application.
// These interfaces enable dependency injection, testability, and loose coupling
// between components.
package interfaces

import "youtube-tui/internal/models"

// VideoSearcher defines the contract for searching YouTube videos.
// Implementations can wrap different search backends (API clients, CLI tools, etc.).
type VideoSearcher interface {
	// Search performs a YouTube search with the given query.
	// It returns a list of matching videos or an error if the search fails.
	Search(query string) ([]models.Video, error)
}

// MediaPlayer defines the contract for controlling audio/video playback.
// Implementations can use different media players (mpv, VLC, etc.).
type MediaPlayer interface {
	// Play starts playback for a given video ID.
	Play(videoID string) error

	// Pause pauses the current playback.
	Pause() error

	// Stop stops the current playback and cleans up resources.
	Stop() error
}

// ConfigProvider defines the contract for accessing configuration values.
// Implementations can source configuration from environment variables, files, etc.
type ConfigProvider interface {
	// Get retrieves a configuration value for the given key.
	// Returns the value and an error if the key is not found or invalid.
	Get(key string) (string, error)

	// GetInt retrieves an integer configuration value for the given key.
	// Returns the value and an error if the key is not found or not an integer.
	GetInt(key string) (int, error)

	// GetBool retrieves a boolean configuration value for the given key.
	// Returns the value and an error if the key is not found or not a boolean.
	GetBool(key string) (bool, error)
}
