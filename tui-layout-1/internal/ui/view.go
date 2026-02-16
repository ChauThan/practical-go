package ui

import "strings"

func (m Model) View() string {
	if m.width <= 0 || m.height <= 0 {
		return ""
	}

	searchHeight := 3
	helpHeight := 2
	resultHeight := m.height - searchHeight - helpHeight
	if resultHeight < 5 {
		resultHeight = 5
	}

	searchFocused := m.focusState == SearchFocused || m.focusState == InputFocused
	search := renderSearchComponent(m.width, searchFocused, m.textInput.View(), m.focusState == InputFocused)
	result := renderResultComponent(m.width, resultHeight, m.focusState == ResultFocused)
	help := renderHelpBarComponent(m.width, m.focusState)

	view := strings.Join([]string{search, result, help}, "\n")
	return fitToHeight(view, m.height)
}
