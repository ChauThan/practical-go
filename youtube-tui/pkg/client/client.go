// Package client provides functionality for interacting with YouTube
// through the yt-dlp command line tool.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"youtube-tui/internal/logging"
	"youtube-tui/internal/models"
)

const maxSearchResults = 25

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

// SearchVideos searches YouTube for videos matching the query using yt-dlp.
// It returns a list of up to maxSearchResults (25) videos with Title, ID, and Uploader.
func SearchVideos(query string) ([]models.Video, error) {
	logging.Info("searching YouTube", "query", query)

	if query == "" {
		logging.Warn("empty search query received")
		return nil, fmt.Errorf("search query cannot be empty")
	}

	logging.Debug("executing yt-dlp command", "query", query, "maxResults", maxSearchResults)

	stdout, stderr, err := executeCommand("yt-dlp", "--dump-single-json", "--flat-playlist", fmt.Sprintf("ytsearch%d:%s", maxSearchResults, query))
	if err != nil {
		logging.Error("yt-dlp command failed", "query", query, "error", err)
		return nil, fmt.Errorf("failed to execute yt-dlp: %w\nstderr: %s", err, stderr)
	}

	logging.Debug("parsing yt-dlp JSON output", "query", query)

	var playlist struct {
		Entries []models.Video `json:"entries"`
	}
	if err := json.Unmarshal([]byte(stdout), &playlist); err != nil {
		logging.Error("failed to parse yt-dlp output", "query", query, "error", err)
		return nil, fmt.Errorf("failed to parse yt-dlp JSON output: %w", err)
	}

	// Limit results to maxSearchResults
	if len(playlist.Entries) > maxSearchResults {
		playlist.Entries = playlist.Entries[:maxSearchResults]
	}

	logging.Info("search completed successfully", "query", query, "resultCount", len(playlist.Entries))
	return playlist.Entries, nil
}
