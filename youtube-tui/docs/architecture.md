# Architecture Documentation

## Overview

YouTube TUI is a terminal-based user interface for searching and playing YouTube videos. The application follows a clean architecture with clear separation of concerns, dependency injection, and environment-based configuration.

## Core Components

### 1. Entry Point ([cmd/youtube-tui/main.go](cmd/youtube-tui/main.go))

The `main` function serves as the application bootstrap:

1. **Initialize Configuration**: Loads configuration from environment variables
2. **Initialize Logging**: Sets up structured logging with file-based output
3. **Create Components**: Instantiates player and UI components
4. **Run Application**: Starts the TUI event loop

### 2. Configuration ([internal/config/config.go](internal/config/config.go))

The configuration system uses a Provider pattern that reads from environment variables:

- **Prefix**: All variables use `YOUTUBE_TUI_` prefix
- **Provider Type**: Implements `ConfigProvider` interface from `interfaces`
- **Type Safety**: Separate methods forGetString, GetInt, GetBool operations
- **Defaults**: Built-in defaults for all configuration values

**Key Design Decisions**:
- Environment variables only (no config files currently)
- Sensible defaults for all settings
- Explicit accessor methods vs generic `Get()` methods

### 3. Logging ([internal/logging/logger.go](internal/logging/logger.go))

The logging system provides file-based structured logging with rotation:

**Features**:
- **File-Based Output**: Logs to file by default to avoid TUI interference
- **Automatic Rotation**: Rotates logs when reaching max size
- **Backup Management**: Keeps configurable number of rotated logs
- **Graceful Degradation**: Falls back to stderr if file logging fails
- **Console Mode**: Optional development mode for debugging

**Configuration**:
- `YOUTUBE_TUI_LOG_FILE`: Path to log file (default: `$HOME/.youtube-tui/youtube-tui.log`)
- `YOUTUBE_TUI_LOG_TO_CONSOLE`: Enable console output (dev only)
- `YOUTUBE_TUI_LOG_FILE_MAX_SIZE`: Max size in MB (default: 10)
- `YOUTUBE_TUI_LOG_FILE_MAX_BACKUPS`: Number of backups to keep (default: 3)
- `YOUTUBE_TUI_LOG_LEVEL`: Log level (DEBUG, INFO, WARN, ERROR)

**Log Rotation**:
- Rotation triggers when file size exceeds max
- Backup files named with timestamp: `youtube-tui.20260112-132045.log`
- Oldest backups deleted when exceeding max backups count
- Uses mutex for thread-safe writes

**Usage**:
```go
// Initialize with configuration
logging.Init(logFilePath, logToConsole, maxSizeMB, maxBackups)

// Set log level from environment
logging.SetLevelFromEnv(levelStr)

// Log messages
logging.Info("application started", "version", "1.0.0")
logging.Error("failed to connect", "error", err)
```

### 4. Interfaces ([internal/interfaces/interfaces.go](internal/interfaces/interfaces.go))

Core abstractions enable dependency injection and testability:

- **VideoSearcher**: Contract for YouTube search operations
- **MediaPlayer**: Contract for media playback control
- **ConfigProvider**: Contract for configuration access

These interfaces allow swapping implementations (e.g., different players, search backends) without modifying core logic.

### 5. User Interface ([internal/ui/](internal/ui/))

The TUI implementation uses Bubble Tea framework in a modular structure:

- **[ui.go](internal/ui/ui.go)**: Main TUI model and initialization
- **[views.go](internal/ui/views.go)**: View components (search results, video details)
- **[handlers.go](internal/ui/handlers.go)**: User input event handlers
- **[layout.go](internal/ui/layout.go)**: Layout and viewport management
- **[search.go](internal/ui/search.go)**: Search input handling
- **[styles.go](internal/ui/styles.go)**: Style definitions and theming

### 6. Player ([internal/player/](internal/player/))

Media Player abstraction enabling multiple media player backends:

- Configuration-based player selection
- MPV integration as default
- Control methods: Play, Pause, Stop

### 7. YouTube Client ([pkg/client/](pkg/client/))

YouTube API client wrapping yt-dlp:

- Executes yt-dlp subprocess for search
- Parses output into structured Video models
- Handles errors and format conversion

### 8. Models ([internal/models/](internal/models/))

Data structures:

- **Video**: Represents a YouTube video with metadata
- **VideoList**: Collection of videos with pagination

## Application Flow

```
┌─────────────┐
│   main()    │ Application entry
└──────┬──────┘
       │
       ├──────────────────────────┐
       │                          │
       ▼                          ▼
┌──────────────┐         ┌──────────────┐
│   Config     │         │    Logging   │
│   Provider   │         │   Init()     │
└──────┬───────┘         └──────┬───────┘
       │                         │
       └─────────────┬───────────┘
                     ▼
              ┌──────────────┐
              │   Player     │
              │   New()      │
              └──────┬───────┘
                     │
                     ▼
              ┌──────────────┐
              │   ui.Run()   │
              └──────────────┘
                     │
                     ▼
              ┌──────────────┐
              │  Bubble Tea  │
              │  Event Loop  │
              └──────────────┘
```

## Data Flow

### Search Flow

1. User enters query in search box
2. [search.go](internal/ui/search.go) captures input
3. [handlers.go](internal/ui/handlers.go) triggers search command
4. [client.go](pkg/client/client.go) calls yt-dlp
5. Results parsed into [Video](internal/models/models.go) models
6. [views.go](internal/ui/views.go) renders results
7. TUI updates display

### Playback Flow

1. User selects video and presses Enter
2. [handlers.go](internal/ui/handlers.go) captures event
3. [player.go](internal/player/player.go) receives video ID
4. Player spawns subprocess with media player
5. Playback starts in separate process
6. TUI continues to accept user input

## Dependencies

### External

- **Bubble Tea**: TUI framework for terminal interfaces
- **Lip Gloss**: Styling for terminal UI
- **yt-dlp**: YouTube downloader (external dependency)

### Internal

- Go standard library:
  - `log/slog`: Structured logging
  - `os/exec`: Subprocess management
  - `encoding/json`: JSON parsing
  - `os`: Environment variables and file I/O

## Design Patterns

1. **Provider Pattern**: Configuration access via Provider interface
2. **Interface Segregation**: Small, focused interfaces for each component
3. **Dependency Injection**: Components receive dependencies via constructors
4. **Model-View-Update (MVU)**: Bubble Tea's reactive architecture
5. **Environment Configuration**: All settings via environment variables

## Logging System Details

### Why File-Based Logging?

The TUI uses the terminal for rendering, so logging to stdout would:
- Mix log output with UI rendering
- Cause visual artifacts and screen corruption
- Degrade user experience

File-based logging solves this by:
- Separating application logs from UI output
- Maintaining clean terminal display
- Providing searchable log history
- Supporting log rotation for disk management

### Log Writer Implementation

The `logWriter` type wraps file operations with:

- **Thread Safety**: Mutex-protected writes
- **Rotation Logic**: Checks size before each write
- **Clean Up**: Removes old backup files
- **Error Handling**: Falls back to stderr on failure

### Extending the Logging System

To add new logging features:

1. **JSON Format**: Replace `NewTextHandler` with `NewJSONHandler`
2. **Remote Logging**: Implement `io.Writer` that sends to remote service
3. **Filtering**: Add custom handler that filters by component or level
4. **Metadata**: Add source location, request ID, or user context

## Future Considerations

### Potential Improvements

1. **Configuration File Support**: Add YAML/TOML config file parsing
2. **Plugin System**: Allow external player and search backends
3. **Async Logging**: Buffered writes for better performance
4. **Log Compression**: Compress rotated log files
5. **Structured Metrics**: Add performance and usage metrics

### Scalability Concerns

- **Concurrent Use**: Current implementation is single-user
- **Performance**: Large result sets may need pagination optimizations
- **Network**: No caching of search results currently

## Testing Strategy

- **Unit Tests**: Test individual components in isolation
- **Integration Tests**: Test component interactions
- **E2E Tests**: Test full user workflows
- **Mock Implementations**: Use interface mocks for external dependencies

