package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func renderSearchComponent(width int, focused bool, textInput string, inputFocused bool) string {
	title := "Search Section (1)"
	var titleStyle lipgloss.Style
	if focused {
		titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
	} else {
		titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	}

	lines := []string{
		titledTopBorder(width, title, titleStyle),
		simpleContentLine(width, textInput),
		bottomBorder(width),
	}

	return strings.Join(lines, "\n")
}
