// Package ui implements the terminal user interface for the YouTube TUI application
// using the Bubble Tea framework
package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"youtube-tui/internal/models"
	"youtube-tui/pkg/client"
)

const (
	minBoxWidth     = 60
	minTermWidth    = 80
	minTermHeight   = 24
	searchBoxHeight = 3
	hMargin         = 2
	vSectionGap     = 1
	selectorLimit   = 156
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

// recomputeLayout recalculates dimensions based on current window size.
func (m *Model) recomputeLayout() {
	if m.windowWidth == 0 || m.windowHeight == 0 {
		return
	}

	// Determine if terminal is too small
	m.tooSmall = m.windowWidth < minTermWidth || m.windowHeight < minTermHeight

	// Calculate box width: min of terminal width minus margins and minBoxWidth,
	// clamped to minBoxWidth if terminal is very small.
	maxBoxWidth := m.windowWidth - 2*hMargin
	if maxBoxWidth < minBoxWidth {
		m.boxWidth = minBoxWidth
	} else {
		m.boxWidth = maxBoxWidth
	}

	// Search box always occupies exactly searchBoxHeight lines of CONTENT.
	// Actual rendered height includes: borders(2) + padding(2) + bottomMargin(1)
	// Total search height = searchBoxHeight + 5
	// Results section takes the remaining height.
	actualSearchBoxHeight := searchBoxHeight + 5 // borders + padding + margin
	m.resultsListHeight = m.windowHeight - actualSearchBoxHeight - vSectionGap
	if m.resultsListHeight < 1 {
		m.resultsListHeight = 1
	}

	// Inner list content width/height for the list component
	// Need to account for box borders and padding
	// focusedBox/unfocusedBox use Padding(1) and RoundedBorder (1 char on each side)
	innerWidth := m.boxWidth - 4 // 2 for borders, 2 for padding
	if innerWidth < 1 {
		innerWidth = 1
	}

	// Adjust results list height for box chrome inside results section:
	// Results section reserves resultsListHeight lines for the entire box
	// Actual inner space = resultsListHeight - borders(2) - padding(2) = resultsListHeight - 4
	// But we also need space for title line + margin (2 lines)
	innerListHeight := m.resultsListHeight - 6 // borders(2) + padding(2) + title+marginTop(2)
	if innerListHeight < 1 {
		innerListHeight = 1
	}

	// Apply sizing to the list
	m.resultsList.SetSize(innerWidth, innerListHeight)

	// Update search input width to fit inside box
	inputWidth := innerWidth
	if inputWidth > selectorLimit {
		inputWidth = selectorLimit
	}
	m.searchInput.Width = inputWidth
	m.searchInput.PlaceholderStyle = lipgloss.NewStyle().Width(inputWidth)
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
}

// NewModel creates a new Model with default state
func NewModel() Model {
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
	}
}

// Init returns the initial command for the application
func (m Model) Init() tea.Cmd {
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
			if err := handleKeyRunes(&m, string(msg.Runes)); err != nil {
				return m, tea.Quit
			}
		case tea.KeyEnter:
			if m.focus == focusSearch {
				cmd = handleEnterKey(&m)
				if cmd != nil {
					return m, cmd
				}
			}
		}

	case searchDoneMsg:
		m.loading = false
		m.results = msg.results
		m.focus = focusResults
		m.searchInput.Blur()

		items := make([]list.Item, len(msg.results))
		for i, video := range msg.results {
			items[i] = videoItem{video: video, index: i}
		}
		m.resultsList.SetItems(items)
		m.resultsList.Select(0)
		m.recomputeLayout()

		return m, nil

	case searchErrMsg:
		m.loading = false
		m.errorMsg = fmt.Sprintf("Search failed: %v", msg.err)
		return m, nil

	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.recomputeLayout()
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

// View renders the TUI interface
func (m Model) View() string {
	s := defaultStyles()

	searchBox := m.renderSearchBox(s)
	resultsBox := m.renderResultsBox(s)

	return lipgloss.JoinVertical(lipgloss.Left, searchBox, resultsBox)
}

// renderSearchBox renders the search section of the UI
func (m Model) renderSearchBox(s styles) string {
	var searchBox string
	if m.focus == focusSearch {
		var statusText string
		if m.tooSmall {
			statusText = s.yellowText.Render("Warning: Terminal too small (recommended: 80x24)")
		} else {
			statusText = s.yellowText.Render("Status: Press Enter to search")
		}
		
		// Render minimal 3-line content: title, input, status
		searchBox = s.focusedBox.Width(m.boxWidth).Render(
			s.focusedTitle.Render("Section 1: Search (Press 1, Enter to search, q to quit)") +
				s.marginTop.Render(m.searchInput.View()) +
				s.marginTop.Render(statusText),
		)
	} else {
		var statusText string
		if m.tooSmall {
			statusText = s.yellowText.Render("Warning: Terminal too small (recommended: 80x24)")
		} else {
			statusText = ""
		}
		statusLine := s.yellowText.Render(statusText)
		
		// Render minimal 3-line content: title, input, status (empty)
		searchBox = s.unfocusedBox.Width(m.boxWidth).Render(
			s.unfocusedTitle.Render("Section 1: Search (Press 1)") +
				s.marginTop.Render(m.searchInput.View()) +
				s.marginTop.Render(statusLine),
		)
	}
	return searchBox
}

// renderResultsBox renders the results section of the UI
func (m Model) renderResultsBox(s styles) string {
	resultsContent := m.renderResultsContent(s)

	var resultsBox string
	if m.focus == focusResults {
		resultsBox = s.focusedBox.Width(m.boxWidth).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				s.focusedTitle.Render(fmt.Sprintf("Section 2: Results - %d videos (Press 2, q to quit)", len(m.results))),
				s.marginTop.Render(resultsContent),
			),
		)
	} else {
		resultsBox = s.unfocusedBox.Width(m.boxWidth).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				s.unfocusedTitle.Render(fmt.Sprintf("Section 2: Results - %d videos (Press 2)", len(m.results))),
				s.marginTop.Render(resultsContent),
			),
		)
	}
	return resultsBox
}

// renderResultsContent renders the content inside the results box
func (m Model) renderResultsContent(s styles) string {
	if m.loading {
		return s.warningText.Render("Searching... please wait")
	}

	if m.errorMsg != "" {
		return s.errorText.Render(m.errorMsg)
	}

	if len(m.results) == 0 {
		return s.grayText.Render("No results. Enter a search query above and press Enter")
	}

	return m.resultsList.View()
}

// renderVideoList renders the list of video results
func (m Model) renderVideoList(s styles) string {
	resultsLines := make([]string, 0, len(m.results))
	for i, video := range m.results {
		resultText := fmt.Sprintf("%d. %s\n   ID: %s | Uploader: %s",
			i+1, video.Title, video.ID, video.Uploader)

		var videoStyle lipgloss.Style
		if i == 0 && m.focus == focusResults {
			videoStyle = lipgloss.NewStyle().Foreground(s.green).Bold(true)
		} else {
			videoStyle = lipgloss.NewStyle().Foreground(s.white)
		}
		resultsLines = append(resultsLines, videoStyle.Render(resultText))
	}
	return lipgloss.JoinVertical(lipgloss.Left, resultsLines...)
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
		titleStyle = lipgloss.NewStyle().
			Foreground(defaultStyles().green).
			Bold(true).
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

// doSearch performs a YouTube search asynchronously
func (m Model) doSearch(query string) tea.Cmd {
	return func() tea.Msg {
		results, err := client.SearchVideos(query)
		if err != nil {
			return searchErrMsg{err: err}
		}
		return searchDoneMsg{results: results}
	}
}

// handleKeyRunes handles single-character key presses
func handleKeyRunes(m *Model, runes string) error {
	// "Typing wins": only treat '1', '2', 'q' as commands when search is not focused
	// or the search input is empty, allowing normal text entry
	if m.focus == focusSearch && m.searchInput.Value() != "" {
		return nil
	}

	switch runes {
	case "1":
		m.focus = focusSearch
		m.searchInput.Focus()
		m.errorMsg = ""
	case "2":
		m.focus = focusResults
		m.searchInput.Blur()
		m.recomputeLayout()
	case "q":
		return fmt.Errorf("quit")
	}
	return nil
}

// handleEnterKey handles the Enter key press
func handleEnterKey(m *Model) tea.Cmd {
	query := m.searchInput.Value()
	if query == "" || len(strings.TrimSpace(query)) == 0 {
		m.errorMsg = "Please enter a search query"
		return nil
	}
	m.loading = true
	m.errorMsg = ""
	m.results = []models.Video{}
	return m.doSearch(strings.TrimSpace(query))
}

// Run starts the TUI application
func Run() error {
	m := NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
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
