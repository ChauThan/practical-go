package main

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
	"tui-layout-2/internal/ui"
)

func main() {
	// Create initial model with mock data
	model := ui.NewModel()

	// Create and run the program
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
