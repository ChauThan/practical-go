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

	// activeColumnStyle highlights focused column with purple border
	activeColumnStyle = columnStyle.Copy().
		BorderForeground(lipgloss.Color(activeColumnColor))

	// activeCardStyle highlights focused card with amber border and bold text
	activeCardStyle = cardStyle.Copy().
		BorderForeground(lipgloss.Color(activeCardColor)).
		Bold(true)

	// titleStyle renders bold, centered column headers
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Align(lipgloss.Center).
		Width(25)

	// appTitleStyle renders bold, centered application title
	appTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Align(lipgloss.Center).
		Width(80).
		Margin(1, 0)

	// helpStyle renders dimmed text for footer help bar
	helpStyle = lipgloss.NewStyle().
		Faint(true).
		Align(lipgloss.Center)
)

// Exported styles for use by View implementation
var (
	// ColumnStyle provides border, padding, and minimum width for columns
	ColumnStyle = columnStyle

	// ActiveColumnStyle highlights focused column with purple border
	ActiveColumnStyle = activeColumnStyle

	// CardStyle provides subtle border and padding for task cards
	CardStyle = cardStyle

	// ActiveCardStyle highlights focused card with amber border and bold text
	ActiveCardStyle = activeCardStyle

	// TitleStyle renders bold, centered column headers
	TitleStyle = titleStyle

	// AppTitleStyle renders bold, centered application title
	AppTitleStyle = appTitleStyle

	// HelpStyle renders dimmed text for footer help bar
	HelpStyle = helpStyle
)
