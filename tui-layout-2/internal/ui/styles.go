package ui

import "charm.land/lipgloss/v2"

// Color constants for visual feedback
const (
	activeColumnColor   = "#7C3AED" // Purple for focused column
	activeCardColor     = "#F59E0B" // Amber for focused card
	inactiveBorderColor = "245"     // Gray for unfocused elements
)

// Base styles for columns and cards (inactive state)
var (
	// columnStyle provides border, padding, and minimum width for columns
	columnStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(inactiveBorderColor)).
		Padding(0, 1).
		Width(25)

	// cardStyle provides subtle border and padding for task cards
	cardStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(inactiveBorderColor)).
		Padding(0)
)
