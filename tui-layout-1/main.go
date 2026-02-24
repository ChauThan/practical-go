package main

import (
	"fmt"
	"os"

	"tui-layout/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run program: %v\n", err)
		os.Exit(1)
	}
}
