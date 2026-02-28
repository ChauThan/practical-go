package ui

import (
	tea "charm.land/bubbletea/v2"
	"tui-layout-2/internal/domain"
)

// Model holds the application state
type Model struct {
	columns    []domain.Column
	focusedCol int
	focusedCard int
	width      int
	height     int
	ready      bool
}

// Init returns the initial command
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles incoming messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
	}
	return m, nil
}

// View renders the UI
func (m Model) View() tea.View {
	var content string
	if !m.ready {
		content = "Initializing..."
	} else {
		content = "Press 'q' to quit.\n"
	}
	return tea.NewView(content)
}
