package ui

import tea "github.com/charmbracelet/bubbletea"

type FocusState int

const (
	SearchFocused FocusState = iota
	ResultFocused
)

type Model struct {
	width      int
	height     int
	focusState FocusState
}

func NewModel() Model {
	return Model{
		focusState: SearchFocused,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch typed := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = typed.Width
		m.height = typed.Height
	case tea.KeyMsg:
		switch typed.String() {
		case "q":
			return m, tea.Quit
		case "1":
			m.focusState = SearchFocused
		case "2":
			m.focusState = ResultFocused
		}
	}

	return m, nil
}
