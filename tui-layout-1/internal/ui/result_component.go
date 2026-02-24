package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func renderResultComponent(width int, height int, focused bool, items []string, scrollOffset int, selectedIdx int) string {
	if height < 2 {
		height = 2
	}

	// Title line
	title := "Result Section (2)"
	var titleStyle lipgloss.Style
	if focused {
		titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
	} else {
		titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	}

	visibleRows := height - 1 // subtract title line
	if visibleRows < 1 {
		visibleRows = 1
	}

	total := len(items)

	// Scrollbar calculation
	showScrollbar := total > visibleRows
	thumbSize := 1
	thumbPos := 0
	if showScrollbar && total > 0 {
		thumbSize = visibleRows * visibleRows / total
		if thumbSize < 1 {
			thumbSize = 1
		}
		freeSpace := visibleRows - thumbSize
		maxScroll := total - visibleRows
		if maxScroll > 0 && freeSpace > 0 {
			thumbPos = scrollOffset * freeSpace / maxScroll
		}
	}

	// Content width (leave 1 col for scrollbar if needed)
	contentWidth := width
	if showScrollbar {
		contentWidth = width - 1
	}
	if contentWidth < 1 {
		contentWidth = 1
	}

	selectedStyle := lipgloss.NewStyle().Reverse(true)
	normalStyle := lipgloss.NewStyle()

	lines := make([]string, 0, height)

	// Title line — padded to full width
	titleText := titleStyle.Render(title)
	titlePad := width - lipgloss.Width(titleText)
	if titlePad < 0 {
		titlePad = 0
	}
	lines = append(lines, titleText+strings.Repeat(" ", titlePad))

	// Content rows
	for row := 0; row < visibleRows; row++ {
		itemIdx := scrollOffset + row
		var text string
		if itemIdx < total {
			text = items[itemIdx]
		}

		// Truncate or pad to contentWidth
		runes := []rune(text)
		if len(runes) > contentWidth {
			runes = runes[:contentWidth]
		}
		text = string(runes)
		padRight := contentWidth - len([]rune(text))
		if padRight < 0 {
			padRight = 0
		}
		text = text + strings.Repeat(" ", padRight)

		var styledLine string
		if itemIdx == selectedIdx && focused {
			styledLine = selectedStyle.Render(text)
		} else {
			styledLine = normalStyle.Render(text)
		}

		// Append scrollbar column
		if showScrollbar {
			var scrollChar string
			if row >= thumbPos && row < thumbPos+thumbSize {
				scrollChar = "█"
			} else {
				scrollChar = "│"
			}
			styledLine = styledLine + scrollChar
		}

		lines = append(lines, styledLine)
	}

	return strings.Join(lines, "\n")
}
