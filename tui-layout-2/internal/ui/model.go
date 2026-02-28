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

// NewModel creates and initializes a new Model with mock kanban data
func NewModel() Model {
	return Model{
		columns: []domain.Column{
			{
				Title: "To Do",
				Cards: []domain.Card{
					{Title: "Fix login bug"},
					{Title: "Write unit tests"},
					{Title: "Update README"},
				},
			},
			{
				Title: "In Progress",
				Cards: []domain.Card{
					{Title: "Refactor auth module"},
					{Title: "Code review PR #42"},
				},
			},
			{
				Title: "Done",
				Cards: []domain.Card{
					{Title: "Setup CI pipeline"},
					{Title: "Deploy v1.0"},
				},
			},
		},
		focusedCol:  0,
		focusedCard: 0,
		width:       80,
		height:      24,
		ready:       false,
	}
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
		case "left", "h":
			if m.focusedCol > 0 {
				m.focusedCol--
				m.focusedCard = 0 // Reset to top card when changing columns
			}
		case "right", "l":
			if m.focusedCol < len(m.columns)-1 {
				m.focusedCol++
				m.focusedCard = 0 // Reset to top card when changing columns
			}
		case "up", "k":
			// Guard against empty columns
			if len(m.columns) == 0 {
				return m, nil
			}
			currentCol := m.columns[m.focusedCol]
			if len(currentCol.Cards) == 0 {
				return m, nil
			}
			if m.focusedCard > 0 {
				m.focusedCard--
			}
		case "down", "j":
			// Guard against empty columns
			if len(m.columns) == 0 {
				return m, nil
			}
			currentCol := m.columns[m.focusedCol]
			if len(currentCol.Cards) == 0 {
				return m, nil
			}
			if m.focusedCard < len(currentCol.Cards)-1 {
				m.focusedCard++
			}
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
