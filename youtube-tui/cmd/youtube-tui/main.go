// Command youtube-tui provides a terminal user interface for searching YouTube videos
package main

import (
	"fmt"
	"os"

	"youtube-tui/internal/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
