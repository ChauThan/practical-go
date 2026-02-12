// Package ui implements the terminal user interface for the YouTube TUI application
// using the Bubble Tea framework
package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"youtube-tui/internal/models"
	"youtube-tui/internal/player"
)

type focusSection int

const (
	focusSearch focusSection = iota
	focusResults
)

// videoItem wraps a models.Video to implement list.Item
type videoItem struct {
	video models.Video
	index int
}

// FilterValue returns the string used for filtering in the list
func (v videoItem) FilterValue() string {
	return v.video.Title
}

// Title returns the title to display in the list item
func (v videoItem) Title() string {
	return fmt.Sprintf("%d. %s", v.index+1, v.video.Title)
}

// Description returns the description for the list item
func (v videoItem) Description() string {
	return fmt.Sprintf("ID: %s | Uploader: %s", v.video.ID, v.video.Uploader)
}

// searchDoneMsg is sent when a search completes successfully
type searchDoneMsg struct {
	results []models.Video
}

// searchErrMsg is sent when a search fails
type searchErrMsg struct {
	err error
}

// Model holds the state for the TUI application
type Model struct {
	focus             focusSection
	searchInput       textinput.Model
	results           []models.Video
	resultsList       list.Model
	loading           bool
	errorMsg          string
	windowWidth       int
	windowHeight      int
	boxWidth          int
	resultsListHeight int
	tooSmall          bool
	player            *player.Player
}

// NewModel creates a new Model with default state
func NewModel(player *player.Player) Model {
	ti := textinput.New()
	ti.Placeholder = "Enter search query..."
	ti.Focus()
	ti.CharLimit = selectorLimit
	ti.Width = minBoxWidth - 4

	delegate := videoListDelegate{}
	li := list.New([]list.Item{}, delegate, minBoxWidth-4, 0)
	li.SetShowHelp(false)
	li.SetShowFilter(false)
	li.SetShowStatusBar(false)
	li.SetShowPagination(true)
	li.SetFilteringEnabled(false)
	li.DisableQuitKeybindings()

	return Model{
		focus:       focusSearch,
		searchInput: ti,
		results:     []models.Video{},
		resultsList: li,
		loading:     false,
		errorMsg:    "",
		player:      player,
	}
}

// Init returns the initial command for the application
func (m Model) Init() tea.Cmd {
	// Force initial layout computation if window size is already set
	if m.windowWidth > 0 && m.windowHeight > 0 {
		RecomputeLayout(&m)
	}
	return nil
}

// Update handles messages and updates the model state
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyCtrlQ, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyRunes:
			if err := HandleKeyRunes(&m, string(msg.Runes)); err != nil {
				return m, tea.Quit
			}
		case tea.KeyEnter:
			cmd = HandleEnterKey(&m, m.player)
			if cmd != nil {
				return m, cmd
			}
		}

	case searchDoneMsg:
		HandleSearchDone(&m, msg.results)
		return m, nil

	case searchErrMsg:
		HandleSearchError(&m, msg.err)
		return m, nil

	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		RecomputeLayout(&m)
		return m, nil
	}

	if m.focus == focusSearch {
		m.searchInput, cmd = m.searchInput.Update(msg)
		return m, cmd
	}

	if m.focus == focusResults {
		m.resultsList, cmd = m.resultsList.Update(msg)
		return m, cmd
	}

	return m, nil
}

// Run starts the TUI application
func Run(p *player.Player) error {
	m := NewModel(p)
	program := tea.NewProgram(m, tea.WithAltScreen())
	_, err := program.Run()
	return err
}

// videoListDelegate implements list.ItemDelegate for styling video items
type videoListDelegate struct{}

// Height returns the height of each list item
func (d videoListDelegate) Height() int {
	return 2
}

// Spacing returns the spacing between items
func (d videoListDelegate) Spacing() int {
	return 1
}

// Update handles messages for an individual list item
func (d videoListDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render renders a single list item
func (d videoListDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	v, ok := item.(videoItem)
	if !ok {
		return
	}

	var (
		titleStyle lipgloss.Style
		descStyle  lipgloss.Style
	)

	if index == m.Index() {
		// Selected item: green + bold, no border
		titleStyle = lipgloss.NewStyle().Foreground(defaultStyles().green).Bold(true).
			MaxWidth(m.Width())
		descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			MaxWidth(m.Width())
	} else {
		titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			MaxWidth(m.Width())
		descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			MaxWidth(m.Width())
	}

	fmt.Fprintf(w, "%s\n%s", titleStyle.Render(v.Title()), descStyle.Render(v.Description()))
}

// styles holds LipGloss style definitions
type styles struct {
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

// defaultStyles returns the default style definitions
func defaultStyles() styles {
	cyan := lipgloss.Color("36")
	yellow := lipgloss.Color("226")
	gray := lipgloss.Color("240")
	green := lipgloss.Color("46")
	red := lipgloss.Color("196")
	white := lipgloss.Color("255")

	return styles{
		cyan:   cyan,
		yellow: yellow,
		gray:   gray,
		green:  green,
		red:    red,
		white:  white,

		focusedBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(cyan).
			Padding(1).
			MarginBottom(1),

		unfocusedBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(gray).
			Padding(1).
			MarginBottom(1),

		focusedTitle: lipgloss.NewStyle().Foreground(cyan).Bold(true),

		unfocusedTitle: lipgloss.NewStyle().Foreground(gray),

		marginTop: lipgloss.NewStyle().MarginTop(1),

		warningText: lipgloss.NewStyle().Foreground(yellow),

		errorText: lipgloss.NewStyle().Foreground(red),

		grayText: lipgloss.NewStyle().Foreground(gray),

		yellowText: lipgloss.NewStyle().Foreground(yellow),
	}
}
