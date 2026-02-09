// Package models provides data structures used throughout the application
package models

// Video represents a YouTube video with essential metadata
type Video struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Uploader string `json:"uploader"`
}

// SearchResults contains the response from yt-dlp for video search
// NOTE: Currently unused, SearchVideos returns []Video directly
type SearchResults struct {
	Entries []Video `json:"entries"`
}
