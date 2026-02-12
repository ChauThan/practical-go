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

	// Add margin between search and results sections
	return lipgloss.JoinVertical(lipgloss.Left, searchBox, s.MarginTop().Render(resultsBox))
}

// RenderSearchBox renders the search section of the UI.
// It displays the title outside the box, then the search input field.
func (m Model) RenderSearchBox(s StyleProvider) string {
	// Render title outside the box
	var title string
	if m.focus == focusSearch {
		title = s.TitleContainer().Render(s.FocusedTitle().Render("Section 1: Search (Press 1, Enter to search, q to quit)"))
	} else {
		title = s.TitleContainer().Render(s.UnfocusedTitle().Render("Section 1: Search (Press 1)"))
	}

	// Render search box (without title, just input + optional warning)
	var searchBox string
	if m.focus == focusSearch {
		if m.tooSmall {
			searchBox = s.FocusedBox().Width(m.boxWidth).Render(
				m.searchInput.View() +
					s.MarginTop().Render(s.YellowText().Render("Warning: Terminal too small (recommended: 80x24)")),
			)
		} else {
			searchBox = s.FocusedBox().Width(m.boxWidth).Render(m.searchInput.View())
		}
	} else {
		if m.tooSmall {
			searchBox = s.UnfocusedBox().Width(m.boxWidth).Render(
				m.searchInput.View() +
					s.MarginTop().Render(s.YellowText().Render("Warning: Terminal too small (recommended: 80x24)")),
			)
		} else {
			searchBox = s.UnfocusedBox().Width(m.boxWidth).Render(m.searchInput.View())
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, title, searchBox)
}

// RenderResultsBox renders the results section of the UI.
// It displays the title outside the box, then the video search results in a list format.
func (m Model) RenderResultsBox(s StyleProvider) string {
	// Render title outside the box
	var title string
	if m.focus == focusResults {
		title = s.TitleContainer().Render(s.FocusedTitle().Render(fmt.Sprintf("Section 2: Results - %d videos (Press 2, q to quit)", len(m.results))))
	} else {
		title = s.TitleContainer().Render(s.UnfocusedTitle().Render(fmt.Sprintf("Section 2: Results - %d videos (Press 2)", len(m.results))))
	}

	// Render results content
	resultsContent := m.RenderResultsContent(s)

	// Render results box (without title, just content)
	// Box height is constrained to m.resultsListHeight which includes borders+padding+content
	var resultsBox string
	if m.focus == focusResults {
		resultsBox = s.FocusedBox().Width(m.boxWidth).Height(m.resultsListHeight).Render(resultsContent)
	} else {
		resultsBox = s.UnfocusedBox().Width(m.boxWidth).Height(m.resultsListHeight).Render(resultsContent)
	}

	return lipgloss.JoinVertical(lipgloss.Left, title, resultsBox)
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
