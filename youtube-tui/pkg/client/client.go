// Package client provides YouTube API client functionality
package client

import (
	"encoding/json"
	"fmt"
	"youtube-tui/internal/models"
	"youtube-tui/pkg/utils"
)

const maxSearchResults = 5

// SearchVideos searches YouTube for videos matching the query using yt-dlp
// It returns a list of up to 5 videos with Title, ID, and Uploader
func SearchVideos(query string) ([]models.Video, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	// Construct yt-dlp command to search for 5 videos and dump JSON output
	// Use --flat-playlist to get simpler JSON output without full metadata
	// and --print to specify exactly what we need
	searchTerm := fmt.Sprintf("ytsearch%d:%s", maxSearchResults, query)
	stdout, stderr, err := utils.ExecuteCommand("yt-dlp", "--dump-single-json", "--flat-playlist", searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to execute yt-dlp: %w\n%s", err, stderr)
	}

	// Parse the playlist JSON output
	var playlist struct {
		Entries []models.Video `json:"entries"`
	}

	if err := json.Unmarshal([]byte(stdout), &playlist); err != nil {
		return nil, fmt.Errorf("failed to parse yt-dlp JSON output: %w", err)
	}

	// Limit results to maxSearchResults
	if len(playlist.Entries) > maxSearchResults {
		playlist.Entries = playlist.Entries[:maxSearchResults]
	}

	if len(playlist.Entries) == 0 {
		return nil, fmt.Errorf("no videos found for query: %s", query)
	}

	return playlist.Entries, nil
}
