package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Logger is the global logger instance
var Logger *slog.Logger

// currentWriter tracks the current log writer for level changes
var currentWriter io.Writer

// logWriter handles writing logs to file with support for log rotation.
type logWriter struct {
	mu          sync.Mutex
	file        *os.File
	logFilePath string
	maxSize     int64 // max size in bytes
	maxBackups  int
	currentSize int64
}

// newLogWriter creates a new log writer that handles log rotation.
func newLogWriter(logFilePath string, maxSizeMB int, maxBackups int) (*logWriter, error) {
	// Ensure the log directory exists
	logDir := filepath.Dir(logFilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Get current file size
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to get log file info: %w", err)
	}

	maxSize := int64(maxSizeMB) * 1024 * 1024

	return &logWriter{
		file:        file,
		logFilePath: logFilePath,
		maxSize:     maxSize,
		maxBackups:  maxBackups,
		currentSize: info.Size(),
	}, nil
}

// Write implements io.Writer for the logWriter.
// It checks for rotation before writing each log entry.
func (lw *logWriter) Write(p []byte) (n int, err error) {
	lw.mu.Lock()
	defer lw.mu.Unlock()

	// Check if we need to rotate
	if lw.currentSize+int64(len(p)) >= lw.maxSize {
		if err := lw.rotate(); err != nil {
			// If rotation fails, write to stderr as fallback
			fmt.Fprintf(os.Stderr, "Log rotation failed: %v\n", err)
		}
	}

	n, err = lw.file.Write(p)
	if err == nil {
		lw.currentSize += int64(n)
	}
	return n, err
}

// rotate rotates the log file by renaming the current file and creating a new one.
func (lw *logWriter) rotate() error {
	// Close current file
	if err := lw.file.Close(); err != nil {
		return fmt.Errorf("failed to close current log file: %w", err)
	}

	// Generate timestamp for backup file
	timestamp := time.Now().Format("20060102-150405")
	backupPath := fmt.Sprintf("%s.%s", lw.logFilePath, timestamp)

	// Rename current log file
	if err := os.Rename(lw.logFilePath, backupPath); err != nil {
		return fmt.Errorf("failed to rotate log file: %w", err)
	}

	// Clean up old backups
	lw.cleanupOldBackups()

	// Create new log file
	file, err := os.OpenFile(lw.logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to create new log file: %w", err)
	}

	lw.file = file
	lw.currentSize = 0

	return nil
}

// cleanupOldBackups removes old log backup files exceeding maxBackups limit.
func (lw *logWriter) cleanupOldBackups() {
	logDir := filepath.Dir(lw.logFilePath)
	baseName := filepath.Base(lw.logFilePath)

	// Find all backup files
	entries, err := os.ReadDir(logDir)
	if err != nil {
		return
	}

	var backups []os.FileInfo
	var backupNames []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		// Check if this is a backup file (starts with base name followed by .)
		if len(name) > len(baseName)+1 && name[:len(baseName)] == baseName && name[len(baseName)] == '.' {
			info, err := entry.Info()
			if err == nil {
				backups = append(backups, info)
				backupNames = append(backupNames, name)
			}
		}
	}

	// If we have too many backups, remove the oldest ones
	if len(backups) > lw.maxBackups {
		// Sort by modification time (oldest first)
		for i := 0; i < len(backups); i++ {
			for j := i + 1; j < len(backups); j++ {
				if backups[i].ModTime().After(backups[j].ModTime()) {
					backups[i], backups[j] = backups[j], backups[i]
					backupNames[i], backupNames[j] = backupNames[j], backupNames[i]
				}
			}
		}

		// Remove oldest backups beyond the limit
		for i := 0; i < len(backups)-lw.maxBackups; i++ {
			backupPath := filepath.Join(logDir, backupNames[i])
			os.Remove(backupPath)
		}
	}
}

// Close closes the log file.
func (lw *logWriter) Close() error {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	if lw.file != nil {
		return lw.file.Close()
	}
	return nil
}

// Init initializes the global logger with a default configuration.
// By default, it logs to a file with text format and INFO level.
// If log file creation fails, it falls back to stderr.
func Init(logFilePath string, logToConsole bool, maxSizeMB int, maxBackups int) {
	var handler slog.Handler
	var writer io.Writer

	// Determine the output destination
	if logToConsole {
		// Console mode - write to stderr to avoid interfering with stdout
		writer = os.Stderr
		handler = slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		// File mode - create or open log file
		logWriter, err := newLogWriter(logFilePath, maxSizeMB, maxBackups)
		if err != nil {
			// Fallback to stderr if file logging fails
			fmt.Fprintf(os.Stderr, "Warning: Failed to initialize file logging (%v). Using console output.\n", err)
			writer = os.Stderr
			handler = slog.NewTextHandler(writer, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			})
		} else {
			writer = logWriter
			handler = slog.NewTextHandler(writer, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			})
		}
	}

	// Save the writer for future level changes
	currentWriter = writer
	Logger = slog.New(handler)
}

// LegacyInit initializes the logger with default configuration for backward compatibility.
// This function is kept for compatibility with existing code.
// Deprecated: Use Init() instead.
func LegacyInit() {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// SetLevel sets the logging level (DEBUG, INFO, WARN, ERROR).
// This preserves the current log writer (file or console).
func SetLevel(level slog.Level) {
	// Use the saved writer to preserve log destination
	writer := currentWriter
	if writer == nil {
		writer = os.Stderr
	}

	Logger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level: level,
	}))
}

// getDefaultLogFilePath returns the default log file path for the current platform.
func getDefaultLogFilePath() string {
	var basePath string
	if runtime.GOOS == "windows" {
		// Windows: %USERPROFILE%\.youtube-tui\youtube-tui.log
		basePath = os.Getenv("USERPROFILE")
		if basePath == "" {
			homeDrive := os.Getenv("HOMEDRIVE")
			homePath := os.Getenv("HOMEPATH")
			if homeDrive != "" && homePath != "" {
				basePath = filepath.Join(homeDrive+homePath, ".youtube-tui")
			} else {
				basePath = ".youtube-tui"
			}
		}
	} else {
		// Linux/macOS: $HOME/.youtube-tui/youtube-tui.log
		basePath = os.Getenv("HOME")
	}

	if basePath == "" {
		// Fallback to current directory
		basePath = "."
	}

	return filepath.Join(basePath, ".youtube-tui", "youtube-tui.log")
}

// InitWithConfig initializes the logger using configuration values from the config package.
func InitWithConfig(logFile, logToConsole, maxSizeMB, maxBackups string, level string) {
	// Convert string values to appropriate types
	filePath := logFile
	if filePath == "" {
		filePath = getDefaultLogFilePath()
	}

	toConsole := false
	if logToConsole == "true" || logToConsole == "1" {
		toConsole = true
	}

	maxSize := 10
	if maxSizeMB != "" {
		if val, err := strconv.Atoi(maxSizeMB); err == nil && val > 0 {
			maxSize = val
		}
	}

	backups := 3
	if maxBackups != "" {
		if val, err := strconv.Atoi(maxBackups); err == nil && val >= 0 {
			backups = val
		}
	}

	Init(filePath, toConsole, maxSize, backups)

	// Set log level
	if level != "" {
		SetLevelFromEnv(level)
	}
}

// SetLevelFromEnv sets the logging level from environment variable string.
func SetLevelFromEnv(levelStr string) {
	var level slog.Level
	switch levelStr {
	case "debug", "DEBUG":
		level = slog.LevelDebug
	case "info", "INFO":
		level = slog.LevelInfo
	case "warn", "WARN", "warning", "WARNING":
		level = slog.LevelWarn
	case "error", "ERROR":
		level = slog.LevelError
	default:
		// Log warning using stderr directly since Logger might not be set yet
		fmt.Fprintf(os.Stderr, "Warning: unknown log level '%s', using INFO\n", levelStr)
		level = slog.LevelInfo
	}

	SetLevel(level)
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
