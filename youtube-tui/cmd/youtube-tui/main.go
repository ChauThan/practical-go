// main.go is the entry point for the YouTube TUI application
package main

import (
	"flag"
	"fmt"
	"log"

	"youtube-tui/pkg/client"
)

func main() {
	query := flag.String("q", "", "Search query for YouTube videos")
	flag.Parse()

	if *query == "" {
		fmt.Println("Usage: youtube-tui -q \"search query\"")
		fmt.Println("Example: youtube-tui -q \"golang tutorial\"")
		return
	}

	videos, err := client.SearchVideos(*query)
	if err != nil {
		log.Fatalf("Error searching videos: %v\n", err)
	}

	fmt.Printf("Found %d videos for: %s\n\n", len(videos), *query)
	for i, video := range videos {
		fmt.Printf("#%d\n", i+1)
		fmt.Printf("Title:   %s\n", video.Title)
		fmt.Printf("ID:      %s\n", video.ID)
		fmt.Printf("Uploader: %s\n", video.Uploader)
		fmt.Println()
	}
}
