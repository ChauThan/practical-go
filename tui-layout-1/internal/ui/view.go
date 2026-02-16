package ui

import "strings"

func (m Model) View() string {
	if m.width <= 0 || m.height <= 0 {
		return ""
	}

	searchHeight := 5
	helpHeight := 2
	resultHeight := m.height - searchHeight - helpHeight
	if resultHeight < 5 {
		resultHeight = 5
	}

	search := renderSearchComponent(m.width, m.focusState == SearchFocused)
	result := renderResultComponent(m.width, resultHeight, m.focusState == ResultFocused)
	help := renderHelpBarComponent(m.width)

	view := strings.Join([]string{search, result, help}, "\n")
	return fitToHeight(view, m.height)
}
