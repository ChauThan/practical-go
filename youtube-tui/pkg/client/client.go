// Package client provides functionality for interacting with YouTube
// through the yt-dlp command line tool.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"youtube-tui/internal/models"
)

const maxSearchResults = 5

// executeCommand runs a command with the given arguments and returns stdout, stderr, and error
func executeCommand(name string, args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("command '%s' failed: %w", name, err)
	}

	return stdout.String(), stderr.String(), nil
}

// SearchVideos searches YouTube for videos matching the query using yt-dlp
// It returns a list of up to five videos with Title, ID, and Uploader
func SearchVideos(query string) ([]models.Video, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	stdout, stderr, err := executeCommand("yt-dlp", "--dump-single-json", "--flat-playlist", fmt.Sprintf("ytsearch%d:%s", maxSearchResults, query))
	if err != nil {
		return nil, fmt.Errorf("failed to execute yt-dlp: %w\n%s", err, stderr)
	}

	var playlist struct {
		Entries []models.Video `json:"entries"`
	}

	if err := json.Unmarshal([]byte(stdout), &playlist); err != nil {
		return nil, fmt.Errorf("failed to parse yt-dlp JSON output: %w", err)
	}

	if len(playlist.Entries) > maxSearchResults {
		playlist.Entries = playlist.Entries[:maxSearchResults]
	}

	if len(playlist.Entries) == 0 {
		return nil, fmt.Errorf("no videos found for query: %s", query)
	}

	return playlist.Entries, nil
}
