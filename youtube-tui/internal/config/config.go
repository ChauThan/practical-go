package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

const (
	// Environment variable prefixes
	envPrefix = "YOUTUBE_TUI_"

	// UI configuration keys
	EnvMinBoxWidth     = envPrefix + "MIN_BOX_WIDTH"
	EnvMinTermWidth    = envPrefix + "MIN_TERM_WIDTH"
	EnvMinTermHeight   = envPrefix + "MIN_TERM_HEIGHT"
	EnvSearchBoxHeight = envPrefix + "SEARCH_BOX_HEIGHT"
	EnvHMargin         = envPrefix + "H_MARGIN"
	EnvVSectionGap     = envPrefix + "V_SECTION_GAP"
	EnvSelectorLimit   = envPrefix + "SELECTOR_LIMIT"

	// Color configuration keys
	EnvColorCyan   = envPrefix + "COLOR_CYAN"
	EnvColorYellow = envPrefix + "COLOR_YELLOW"
	EnvColorGray   = envPrefix + "COLOR_GRAY"
	EnvColorGreen  = envPrefix + "COLOR_GREEN"
	EnvColorRed    = envPrefix + "COLOR_RED"
	EnvColorWhite  = envPrefix + "COLOR_WHITE"

	// Player configuration keys
	EnvPlayerExecutable = envPrefix + "PLAYER_EXECUTABLE"
	EnvPlayerAutoStop   = envPrefix + "PLAYER_AUTO_STOP"

	// Search configuration keys
	EnvSearchMaxResults = envPrefix + "SEARCH_MAX_RESULTS"

	// Logging configuration keys
	EnvLogFile           = envPrefix + "LOG_FILE"
	EnvLogToConsole      = envPrefix + "LOG_TO_CONSOLE"
	EnvLogFileMaxSize    = envPrefix + "LOG_FILE_MAX_SIZE"
	EnvLogFileMaxBackups = envPrefix + "LOG_FILE_MAX_BACKUPS"
)

// Default configuration values
const (
	DefaultMinBoxWidth     = 60
	DefaultMinTermWidth    = 80
	DefaultMinTermHeight   = 24
	DefaultSearchBoxHeight = 3
	DefaultHMargin         = 2
	DefaultVSectionGap     = 1
	DefaultSelectorLimit   = 156
)

// Default color values (ANSI color codes)
const (
	DefaultColorCyan   = "36"
	DefaultColorYellow = "226"
	DefaultColorGray   = "240"
	DefaultColorGreen  = "46"
	DefaultColorRed    = "196"
	DefaultColorWhite  = "255"
)

// Default player configuration
const (
	DefaultPlayerExecutable = "mpv"
	DefaultPlayerAutoStop   = true
)

// Default search configuration
const (
	DefaultSearchMaxResults = 25
)

// Default logging configuration
const (
	DefaultLogFileMaxSize    = 10 // in MB
	DefaultLogFileMaxBackups = 3
)

// Provider implements the ConfigProvider interface using environment variables.
type Provider struct{}

// NewProvider creates a new configuration provider that reads from environment variables.
func NewProvider() *Provider {
	return &Provider{}
}

// Get retrieves a string configuration value for the given key.
// Returns the value from environment, or an empty string if the key is not set.
// This is the public interface method that matches interfaces.ConfigProvider.
func (p *Provider) Get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("configuration key '%s' not set", key)
	}
	return value, nil
}

// GetWithDefault retrieves a string configuration value for the given key,
// returning the provided default if the key is not set.
func (p *Provider) GetWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetInt retrieves an integer configuration value for the given key.
// Returns an error if the key is not set or not a valid integer.
func (p *Provider) GetInt(key string) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("config: environment variable '%s' not set", key)
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("config: environment variable '%s' has invalid integer value '%s': %w", key, value, err)
	}
	return result, nil
}

// GetIntWithDefault retrieves an integer configuration value for the given key,
// returning the provided default if the key is not set or invalid.
func (p *Provider) GetIntWithDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if result, err := strconv.Atoi(value); err == nil {
			return result
		}
	}
	return defaultValue
}

// GetBool retrieves a boolean configuration value for the given key.
// Returns an error if the key is not set or not a valid boolean.
// Valid boolean values: true, false, 1, 0 (case-insensitive).
func (p *Provider) GetBool(key string) (bool, error) {
	value := os.Getenv(key)
	if value == "" {
		return false, fmt.Errorf("config: environment variable '%s' not set", key)
	}

	switch value {
	case "true", "True", "TRUE", "1":
		return true, nil
	case "false", "False", "FALSE", "0":
		return false, nil
	default:
		return false, fmt.Errorf("config: environment variable '%s' has invalid boolean value '%s' (expected: true/false/1/0)", key, value)
	}
}

// GetBoolWithDefault retrieves a boolean configuration value for the given key,
// returning the provided default if the key is not set or invalid.
func (p *Provider) GetBoolWithDefault(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		switch value {
		case "true", "True", "TRUE", "1":
			return true
		case "false", "False", "FALSE", "0":
			return false
		}
	}
	return defaultValue
}

// GetMinBoxWidth returns the minimum box width configuration.
func (p *Provider) GetMinBoxWidth() int {
	return p.GetIntWithDefault(EnvMinBoxWidth, DefaultMinBoxWidth)
}

// GetMinTermWidth returns the minimum terminal width configuration.
func (p *Provider) GetMinTermWidth() int {
	return p.GetIntWithDefault(EnvMinTermWidth, DefaultMinTermWidth)
}

// GetMinTermHeight returns the minimum terminal height configuration.
func (p *Provider) GetMinTermHeight() int {
	return p.GetIntWithDefault(EnvMinTermHeight, DefaultMinTermHeight)
}

// GetSearchBoxHeight returns the search box height configuration.
func (p *Provider) GetSearchBoxHeight() int {
	return p.GetIntWithDefault(EnvSearchBoxHeight, DefaultSearchBoxHeight)
}

// GetHMargin returns the horizontal margin configuration.
func (p *Provider) GetHMargin() int {
	return p.GetIntWithDefault(EnvHMargin, DefaultHMargin)
}

// GetVSectionGap returns the vertical section gap configuration.
func (p *Provider) GetVSectionGap() int {
	return p.GetIntWithDefault(EnvVSectionGap, DefaultVSectionGap)
}

// GetSelectorLimit returns the selector text limit configuration.
func (p *Provider) GetSelectorLimit() int {
	return p.GetIntWithDefault(EnvSelectorLimit, DefaultSelectorLimit)
}

// GetColorCyan returns the cyan ANSI color code.
func (p *Provider) GetColorCyan() string {
	return p.GetWithDefault(EnvColorCyan, DefaultColorCyan)
}

// GetColorYellow returns the yellow ANSI color code.
func (p *Provider) GetColorYellow() string {
	return p.GetWithDefault(EnvColorYellow, DefaultColorYellow)
}

// GetColorGray returns the gray ANSI color code.
func (p *Provider) GetColorGray() string {
	return p.GetWithDefault(EnvColorGray, DefaultColorGray)
}

// GetColorGreen returns the green ANSI color code.
func (p *Provider) GetColorGreen() string {
	return p.GetWithDefault(EnvColorGreen, DefaultColorGreen)
}

// GetColorRed returns the red ANSI color code.
func (p *Provider) GetColorRed() string {
	return p.GetWithDefault(EnvColorRed, DefaultColorRed)
}

// GetColorWhite returns the white ANSI color code.
func (p *Provider) GetColorWhite() string {
	return p.GetWithDefault(EnvColorWhite, DefaultColorWhite)
}

// GetPlayerExecutable returns the player executable path.
func (p *Provider) GetPlayerExecutable() string {
	return p.GetWithDefault(EnvPlayerExecutable, DefaultPlayerExecutable)
}

// GetPlayerAutoStop returns whether to auto-stop the player.
func (p *Provider) GetPlayerAutoStop() bool {
	return p.GetBoolWithDefault(EnvPlayerAutoStop, DefaultPlayerAutoStop)
}

// GetSearchMaxResults returns the maximum number of search results.
func (p *Provider) GetSearchMaxResults() int {
	return p.GetIntWithDefault(EnvSearchMaxResults, DefaultSearchMaxResults)
}

// GetDefaultLogFile returns the default log file path based on the platform.
func GetDefaultLogFile() string {
	var basePath string
	if runtime.GOOS == "windows" {
		basePath = filepath.Join(os.Getenv("USERPROFILE"), ".youtube-tui")
	} else {
		basePath = filepath.Join(os.Getenv("HOME"), ".youtube-tui")
	}
	return filepath.Join(basePath, "youtube-tui.log")
}

// GetLogFile returns the log file path to use, or the default if not configured.
func (p *Provider) GetLogFile() string {
	if value := os.Getenv(EnvLogFile); value != "" {
		return value
	}
	return GetDefaultLogFile()
}

// GetLogToConsole returns whether to log to console (dev mode).
func (p *Provider) GetLogToConsole() bool {
	return p.GetBoolWithDefault(EnvLogToConsole, false)
}

// GetLogFileMaxSize returns the maximum log file size in MB.
func (p *Provider) GetLogFileMaxSize() int {
	return p.GetIntWithDefault(EnvLogFileMaxSize, DefaultLogFileMaxSize)
}

// GetLogFileMaxBackups returns the number of old log files to keep.
func (p *Provider) GetLogFileMaxBackups() int {
	return p.GetIntWithDefault(EnvLogFileMaxBackups, DefaultLogFileMaxBackups)
}
