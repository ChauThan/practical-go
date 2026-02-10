package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"youtube-tui/internal/models"
	"youtube-tui/internal/player"
)

// VideoPlayMsg is sent when user requests playback of a video
type VideoPlayMsg struct {
	Video models.Video
}

// HandleKeyRunes handles single-character key presses for navigation and control.
// It interprets '1' and '2' as focus switches between search and results,
// 'q' as quit command. When search is focused and has content, keystrokes
// are passed through for normal text entry.
func HandleKeyRunes(m *Model, runes string) error {
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
		RecomputeLayout(m)
	case "q":
		return fmt.Errorf("quit")
	}
	return nil
}

// HandleEnterKey handles the Enter key press.
// When in search focus, it triggers a search operation.
// When in results focus with a selected video, it triggers playback.
func HandleEnterKey(m *Model, player *player.Player) tea.Cmd {
	if m.focus == focusSearch {
		return handleSearchEnter(m)
	}
	if m.focus == focusResults {
		return handlePlayEnter(m, player)
	}
	return nil
}

// handleSearchEnter handles Enter key press in search box.
func handleSearchEnter(m *Model) tea.Cmd {
	query := m.searchInput.Value()
	if query == "" || len(strings.TrimSpace(query)) == 0 {
		m.errorMsg = "Please enter a search query"
		return nil
	}
	m.loading = true
	m.errorMsg = ""
	m.results = []models.Video{}
	return DoSearch(m, strings.TrimSpace(query))
}

// handlePlayEnter handles Enter key press in results list.
// It triggers playback of the currently selected video.
func handlePlayEnter(m *Model, player *player.Player) tea.Cmd {
	selectedIndex := m.resultsList.Index()
	if selectedIndex < 0 || selectedIndex >= len(m.results) {
		return nil
	}

	selectedVideo := m.results[selectedIndex]
	return func() tea.Msg {
		if err := player.Play(selectedVideo.ID); err != nil {
			return searchErrMsg{err: fmt.Errorf("failed to play video: %w", err)}
		}
		return VideoPlayMsg{Video: selectedVideo}
	}
}
