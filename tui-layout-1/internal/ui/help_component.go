package ui

func renderHelpBarComponent(width int) string {
	separator := repeat("─", width)
	hints := truncateAndPad("↑/↓: Navigate | Enter: Select | q: Quit", width)
	return separator + "\n" + hints
}
