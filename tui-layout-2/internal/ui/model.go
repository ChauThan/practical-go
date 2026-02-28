package ui

import (
	tea "charm.land/bubbletea/v2"
	"tui-layout-2/internal/domain"

	"charm.land/lipgloss/v2"
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

// columnWidth calculates responsive column width with minimum width guard
func (m Model) columnWidth() int {
	if m.width == 0 {
		return 25 // Default for uninitialized state
	}
	columnWidth := m.width / 3
	if columnWidth < 25 {
		return 25 // Minimum width guard
	}
	return columnWidth
}

// renderCard renders a single card with appropriate styling based on focus state
func renderCard(card domain.Card, cardIdx int, colIdx int, focusedCol int, focusedCard int) string {
	// Check if this card is focused
	if colIdx == focusedCol && cardIdx == focusedCard {
		return ActiveCardStyle.Render(card.Title)
	}
	return CardStyle.Render(card.Title)
}

// renderColumn renders a single column with title and cards
func renderColumn(col domain.Column, colIdx int, focusedCol int, focusedCard int) string {
	// Render column title
	title := TitleStyle.Render(col.Title)

	// Render all cards in this column
	var cardStrings []string
	for cardIdx, card := range col.Cards {
		cardStrings = append(cardStrings, renderCard(card, cardIdx, colIdx, focusedCol, focusedCard))
	}

	// Stack cards vertically
	cards := lipgloss.JoinVertical(lipgloss.Left, cardStrings...)

	// Combine title and cards
	columnContent := lipgloss.JoinVertical(lipgloss.Left, title, cards)

	// Apply appropriate column style
	if colIdx == focusedCol {
		return ActiveColumnStyle.Render(columnContent)
	}
	return ColumnStyle.Render(columnContent)
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
		// Render application title
		title := AppTitleStyle.Render("KANBAN BOARD")

		// Render all columns
		var columnStrings []string
		for colIdx, col := range m.columns {
			columnStrings = append(columnStrings, renderColumn(col, colIdx, m.focusedCol, m.focusedCard))
		}

		// Arrange columns horizontally
		board := lipgloss.JoinHorizontal(lipgloss.Top, columnStrings...)

		// Render help bar
		help := HelpStyle.Render("←/→: Move column | ↑/↓: Move card | q: Quit")

		// Combine all sections vertically
		content = lipgloss.JoinVertical(lipgloss.Left, title, board, help)
	}
	return tea.NewView(content)
}
