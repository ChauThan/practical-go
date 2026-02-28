package main

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
	"tui-layout-2/internal/ui"
)

func main() {
	// Create initial model
	// Note: Model fields will be initialized in Phase 2
	var model ui.Model

	// Create and run the program
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
