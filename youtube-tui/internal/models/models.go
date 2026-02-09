// Package models defines data structures for the application
package models

// Video represents a YouTube video extracted from yt-dlp JSON output
type Video struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Uploader string `json:"uploader"`
}

// SearchResults contains the response from yt-dlp for video search
type SearchResults struct {
	Entries []Video `json:"entries"`
}
