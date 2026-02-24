package ui

func renderHelpBarComponent(width int, focusState FocusState) string {
	separator := repeat("─", width)

	var hints string
	switch focusState {
	case SearchFocused:
		hints = truncateAndPad("i: Focus input field | 1-2: Navigate sections | q: Quit", width)
	case InputFocused:
		hints = truncateAndPad("Type to search | q: Quit input field", width)
	case ResultFocused:
		hints = truncateAndPad("↑/↓: Navigate | Enter: Select | 1-2: Navigate sections | q: Quit", width)
	default:
		hints = truncateAndPad("q: Quit", width)
	}

	return separator + "\n" + hints
}
