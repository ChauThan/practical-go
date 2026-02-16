package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func titledTopBorder(width int, title string, style lipgloss.Style) string {
	if width <= 1 {
		return ""
	}

	if width == 2 {
		return "┌┐"
	}

	innerWidth := width - 2
	titleText := " " + title + " "
	styledTitle := style.Render(titleText)
	titleWidth := lipgloss.Width(styledTitle)
	if titleWidth >= innerWidth {
		return "┌" + repeat("─", innerWidth) + "┐"
	}

	left := (innerWidth - titleWidth) / 2
	right := innerWidth - titleWidth - left
	return "┌" + repeat("─", left) + styledTitle + repeat("─", right) + "┐"
}

func bottomBorder(width int) string {
	if width <= 1 {
		return ""
	}
	if width == 2 {
		return "└┘"
	}
	return "└" + repeat("─", width-2) + "┘"
}

func innerTopLine(width int) string {
	if width < 6 {
		return truncateAndPad("││", width)
	}
	innerWidth := width - 4
	return "│ " + "┌" + repeat("─", innerWidth-2) + "┐" + " │"
}

func innerBottomLine(width int) string {
	if width < 6 {
		return truncateAndPad("││", width)
	}
	innerWidth := width - 4
	return "│ " + "└" + repeat("─", innerWidth-2) + "┘" + " │"
}

func innerContentLine(width int, text string) string {
	if width < 6 {
		return truncateAndPad("││", width)
	}
	textWidth := width - 6
	content := truncateAndPad(text, textWidth)
	return "│ " + "│" + content + "│" + " │"
}

func truncateAndPad(text string, width int) string {
	if width <= 0 {
		return ""
	}

	runes := []rune(text)
	if len(runes) > width {
		runes = runes[:width]
	}
	trimmed := string(runes)
	remaining := width - len(runes)
	if remaining <= 0 {
		return trimmed
	}
	return trimmed + repeat(" ", remaining)
}

func repeat(token string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(token, count)
}

func runeLen(text string) int {
	return len([]rune(text))
}

func fitToHeight(view string, height int) string {
	if height <= 0 {
		return ""
	}

	lines := strings.Split(view, "\n")
	if len(lines) > height {
		return strings.Join(lines[:height], "\n")
	}
	if len(lines) == height {
		return view
	}

	paddingCount := height - len(lines)
	padding := make([]string, 0, paddingCount)
	for index := 0; index < paddingCount; index++ {
		padding = append(padding, "")
	}

	return strings.Join(append(lines, padding...), "\n")
}
