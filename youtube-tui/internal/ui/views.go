package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View renders the complete TUI interface.
// It joins the search box and resultsbox vertically.
func (m Model) View() string {
	s := NewStyleProvider()

	searchBox := m.RenderSearchBox(s)
	resultsBox := m.RenderResultsBox(s)

	return lipgloss.JoinVertical(lipgloss.Left, searchBox, resultsBox)
}

// RenderSearchBox renders the search section of the UI.
// It displays the search input field and warnings if terminal is too small.
func (m Model) RenderSearchBox(s StyleProvider) string {
	var searchBox string
	if m.focus == focusSearch {
		// Render minimal 2-line content: title, input (plus warning if terminal too small)
		if m.tooSmall {
			searchBox = s.FocusedBox().Width(m.boxWidth).Render(
				s.FocusedTitle().Render("Section 1: Search (Press 1, Enter to search, q to quit)") +
					s.MarginTop().Render(m.searchInput.View()) +
					s.MarginTop().Render(s.YellowText().Render("Warning: Terminal too small (recommended: 80x24)")),
			)
		} else {
			searchBox = s.FocusedBox().Width(m.boxWidth).Render(
				s.FocusedTitle().Render("Section 1: Search (Press 1, Enter to search, q to quit)") +
					s.MarginTop().Render(m.searchInput.View()),
			)
		}
	} else {
		// Render minimal 2-line content: title, input (plus warning if terminal too small)
		if m.tooSmall {
			searchBox = s.UnfocusedBox().Width(m.boxWidth).Render(
				s.UnfocusedTitle().Render("Section 1: Search (Press 1)") +
					s.MarginTop().Render(m.searchInput.View()) +
					s.MarginTop().Render(s.YellowText().Render("Warning: Terminal too small (recommended: 80x24)")),
			)
		} else {
			searchBox = s.UnfocusedBox().Width(m.boxWidth).Render(
				s.UnfocusedTitle().Render("Section 1: Search (Press 1)") +
					s.MarginTop().Render(m.searchInput.View()),
			)
		}
	}
	return searchBox
}

// RenderResultsBox renders the results section of the UI.
// It displays video search results in a list format.
func (m Model) RenderResultsBox(s StyleProvider) string {
	resultsContent := m.RenderResultsContent(s)

	var resultsBox string
	if m.focus == focusResults {
		resultsBox = s.FocusedBox().Width(m.boxWidth).Height(m.resultsListHeight).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				s.FocusedTitle().Render(fmt.Sprintf("Section 2: Results - %d videos (Press 2, q to quit)", len(m.results))),
				s.MarginTop().Render(resultsContent),
			),
		)
	} else {
		resultsBox = s.UnfocusedBox().Width(m.boxWidth).Height(m.resultsListHeight).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				s.UnfocusedTitle().Render(fmt.Sprintf("Section 2: Results - %d videos (Press 2)", len(m.results))),
				s.MarginTop().Render(resultsContent),
			),
		)
	}
	return resultsBox
}

// RenderResultsContent renders the content inside the results box.
// It shows loading state, error messages, or the list of results.
func (m Model) RenderResultsContent(s StyleProvider) string {
	if m.loading {
		return s.WarningText().Render("Searching... please wait")
	}

	if m.errorMsg != "" {
		return s.ErrorText().Render(m.errorMsg)
	}

	if len(m.results) == 0 {
		return s.GrayText().Render("No results. Enter a search query above and press Enter")
	}

	return m.resultsList.View()
}
