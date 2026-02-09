// Command youtube-tui provides a terminal user interface for searching YouTube videos
package main

import (
	"youtube-tui/internal/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
