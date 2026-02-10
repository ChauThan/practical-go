package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"

	"youtube-tui/internal/models"
	"youtube-tui/pkg/client"
)

// DoSearch performs a YouTube search asynchronously.
// It returns a Bubble Tea command that spawns a goroutine to execute the search.
// The goroutine communicates results back via returned messages.
func DoSearch(m *Model, query string) tea.Cmd {
	return func() tea.Msg {
		results, err := client.SearchVideos(query)
		if err != nil {
			return searchErrMsg{err: fmt.Errorf("search failed for query '%s': %w", query, err)}
		}
		return searchDoneMsg{results: results}
	}
}

// HandleSearchDone processes a successful search completion.
// It updates the model with results, switches focus to the results section,
// and populates the list component with video items.
func HandleSearchDone(m *Model, results []models.Video) {
	m.loading = false
	m.results = results
	m.focus = focusResults
	m.searchInput.Blur()

	items := make([]list.Item, len(results))
	for i, video := range results {
		items[i] = videoItem{video: video, index: i}
	}
	m.resultsList.SetItems(items)
	m.resultsList.Select(0)
	RecomputeLayout(m)
}

// HandleSearchError processes a failed search operation.
// It updates the model with an error message to display to the user.
func HandleSearchError(m *Model, err error) {
	m.loading = false
	m.errorMsg = fmt.Sprintf("Search failed: %v", err)
}
