# YouTube TUI

A terminal user interface for YouTube, built with Go. Search and play YouTube videos from the comfort of your terminal.

## Prerequisites

This project requires `yt-dlp` to be installed on your system.

### Installing yt-dlp

**On Linux (Ubuntu/Debian):**
```bash
sudo apt install yt-dlp
```

**On macOS:**
```bash
brew install yt-dlp
```

**On Windows:**
```bash
choco install yt-dlp
```

Or install via pip:
```bash
pip install yt-dlp
```

For more installation methods, visit [yt-dlp's official GitHub repository](https://github.com/yt-dlp/yt-dlp).

## Building

```bash
go build
```

## Running

```bash
./youtube-tui
```

## Configuration

The application can be configured via environment variables. All configuration options have sensible defaults, so configuration is optional.

### UI Configuration

| Variable | Default | Description |
|----------|----------|-------------|
| `YOUTUBE_TUI_MIN_BOX_WIDTH` | 60 | Minimum width for UI boxes |
| `YOUTUBE_TUI_MIN_TERM_WIDTH` | 80 | Minimum terminal width required |
| `YOUTUBE_TUI_MIN_TERM_HEIGHT` | 24 | Minimum terminal height required |
| `YOUTUBE_TUI_SEARCH_BOX_HEIGHT` | 3 | Height of the search input box |
| `YOUTUBE_TUI_H_MARGIN` | 2 | Horizontal margin around boxes |
| `YOUTUBE_TUI_V_SECTION_GAP` | 1 | Vertical gap between sections |
| `YOUTUBE_TUI_SELECTOR_LIMIT` | 156 | Maximum character limit for search input |

### Color Configuration

Customize the color scheme using ANSI color codes (0-255):

| Variable | Default | Description |
|----------|----------|-------------|
| `YOUTUBE_TUI_COLOR_CYAN` | 36 | Cyan color for focused elements |
| `YOUTUBE_TUI_COLOR_YELLOW` | 226 | Yellow color for warnings |
| `YOUTUBE_TUI_COLOR_GRAY` | 240 | Gray color for unfocused elements |
| `YOUTUBE_TUI_COLOR_GREEN` | 46 | Green color for selection highlights |
| `YOUTUBE_TUI_COLOR_RED` | 196 | Red color for errors |
| `YOUTUBE_TUI_COLOR_WHITE` | 255 | White color for text |

### Player Configuration

| Variable | Default | Description |
|----------|----------|-------------|
| `YOUTUBE_TUI_PLAYER_EXECUTABLE` | mpv | Media player executable |
| `YOUTUBE_TUI_PLAYER_AUTO_STOP` | true | Auto-stop player on video end |

### Search Configuration

| Variable | Default | Description |
|----------|----------|-------------|
| `YOUTUBE_TUI_SEARCH_MAX_RESULTS` | 25 | Maximum number of search results |

### Logging Configuration

| Variable | Default | Description |
|----------|----------|-------------|
| `YOUTUBE_TUI_LOG_LEVEL` | INFO | Logging level (DEBUG, INFO, WARN, ERROR) |

### Example Configuration

```bash
# Set a custom terminal width
export YOUTUBE_TUI_MIN_TERM_WIDTH=100

# Use a different media player
export YOUTUBE_TUI_PLAYER_EXECUTABLE=vlc

# Enable debug logging
export YOUTUBE_TUI_LOG_LEVEL=DEBUG

# Change the selection color
export YOUTUBE_TUI_COLOR_GREEN=34
```

## Usage

1. **Start the application**: Run `./youtube-tui` in your terminal
2. **Search**: Type your query in the search box and press Enter
3. **Browse**: Press `2` to navigate to results, use arrow keys to select
4. **Play**: Press Enter on a video to play it in MPV (or your configured player)
5. **Quit**: Press `q` or `Ctrl+C` to exit

## Architecture

The application follows a clean architecture with separated concerns:

- **`cmd/youtube-tui/`**: Application entry point
- **`internal/ui/`**: Terminal user interface implementation using Bubble Tea
- **`internal/player/`**: Media player abstraction and MPV integration
- **`internal/config/`**: Environment-based configuration provider
- **`internal/logging/`**: Structured logging using Go's log/slog
- **`internal/interfaces/`**: Core interface definitions for dependency injection
- **`pkg/client/`**: YouTube API client wrapping yt-dlp
- **`internal/models/`**: Data models for videos and search results

For detailed architecture documentation, see [docs/architecture.md](docs/architecture.md).

## Contributing

Contributions are welcome! Please ensure:
- Code follows Go conventions and formatting
- Error messages include appropriate context
- Comments explain complex logic when necessary
- Changes are tested for basic functionality

## License

 TBD
