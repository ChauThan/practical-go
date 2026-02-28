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
