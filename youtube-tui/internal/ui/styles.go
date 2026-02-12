package ui

import "github.com/charmbracelet/lipgloss"

// StyleProvider encapsulates all LipGloss style definitions for the TUI.
// It provides factory methods for creating styled components.
type StyleProvider struct {
	cyan   lipgloss.Color
	yellow lipgloss.Color
	gray   lipgloss.Color
	green  lipgloss.Color
	red    lipgloss.Color
	white  lipgloss.Color

	focusedBox     lipgloss.Style
	unfocusedBox   lipgloss.Style
	focusedTitle   lipgloss.Style
	unfocusedTitle lipgloss.Style
	marginTop      lipgloss.Style
	warningText    lipgloss.Style
	errorText      lipgloss.Style
	grayText       lipgloss.Style
	yellowText     lipgloss.Style
}

// NewStyleProvider creates a new StyleProvider with default color scheme and styles.
func NewStyleProvider() StyleProvider {
	cyan := lipgloss.Color("36")
	yellow := lipgloss.Color("226")
	gray := lipgloss.Color("240")
	green := lipgloss.Color("46")
	red := lipgloss.Color("196")
	white := lipgloss.Color("255")

	return StyleProvider{
		cyan:   cyan,
		yellow: yellow,
		gray:   gray,
		green:  green,
		red:    red,
		white:  white,

		focusedBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(cyan).
			Padding(1),

		unfocusedBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(gray).
			Padding(1),

		focusedTitle: lipgloss.NewStyle().Foreground(white).Bold(true),

		unfocusedTitle: lipgloss.NewStyle().Foreground(gray),

		marginTop: lipgloss.NewStyle().MarginTop(1),

		warningText: lipgloss.NewStyle().Foreground(yellow),

		errorText: lipgloss.NewStyle().Foreground(red),

		grayText: lipgloss.NewStyle().Foreground(gray),

		yellowText: lipgloss.NewStyle().Foreground(yellow),
	}
}

// NewStyle returns a blank lipgloss Style.
func (s StyleProvider) NewStyle() lipgloss.Style {
	return lipgloss.NewStyle()
}

// FocusedBox returns the style for a focused box border and padding.
func (s StyleProvider) FocusedBox() lipgloss.Style {
	return s.focusedBox
}

// UnfocusedBox returns the style for an unfocused box border and padding.
func (s StyleProvider) UnfocusedBox() lipgloss.Style {
	return s.unfocusedBox
}

// FocusedTitle returns the style for a focused section title.
func (s StyleProvider) FocusedTitle() lipgloss.Style {
	return s.focusedTitle
}

// UnfocusedTitle returns the style for an unfocused section title.
func (s StyleProvider) UnfocusedTitle() lipgloss.Style {
	return s.unfocusedTitle
}

// TitleContainer returns the style for title containers (with bottom margin).
func (s StyleProvider) TitleContainer() lipgloss.Style {
	return lipgloss.NewStyle().MarginBottom(1)
}

// MarginTop returns a margin-top style.
func (s StyleProvider) MarginTop() lipgloss.Style {
	return s.marginTop
}

// WarningText returns the style for warning messages.
func (s StyleProvider) WarningText() lipgloss.Style {
	return s.warningText
}

// ErrorText returns the style for error messages.
func (s StyleProvider) ErrorText() lipgloss.Style {
	return s.errorText
}

// GrayText returns the style for gray/informational text.
func (s StyleProvider) GrayText() lipgloss.Style {
	return s.grayText
}

// YellowText returns the style for yellow/highlighted text.
func (s StyleProvider) YellowText() lipgloss.Style {
	return s.yellowText
}

// Cyan returns the cyan color constant.
func (s StyleProvider) Cyan() lipgloss.Color {
	return s.cyan
}

// Yellow returns the yellow color constant.
func (s StyleProvider) Yellow() lipgloss.Color {
	return s.yellow
}

// Gray returns the gray color constant.
func (s StyleProvider) Gray() lipgloss.Color {
	return s.gray
}

// Green returns the green color constant.
func (s StyleProvider) Green() lipgloss.Color {
	return s.green
}

// Red returns the red color constant.
func (s StyleProvider) Red() lipgloss.Color {
	return s.red
}

// White returns the white color constant.
func (s StyleProvider) White() lipgloss.Color {
	return s.white
}
