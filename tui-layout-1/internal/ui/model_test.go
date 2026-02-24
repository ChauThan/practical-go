package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewModel(t *testing.T) {
	model := NewModel()
	if model.focusState != SearchFocused {
		t.Errorf("Expected focusState to be SearchFocused, got %v", model.focusState)
	}
}

func TestUpdateKey1(t *testing.T) {
	model := NewModel()
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}}
	updatedModel, _ := model.Update(msg)
	if updatedModel.(Model).focusState != SearchFocused {
		t.Errorf("Expected focusState to be SearchFocused after pressing '1', got %v", updatedModel.(Model).focusState)
	}
}

func TestUpdateKey2(t *testing.T) {
	model := NewModel()
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}}
	updatedModel, _ := model.Update(msg)
	if updatedModel.(Model).focusState != ResultFocused {
		t.Errorf("Expected focusState to be ResultFocused after pressing '2', got %v", updatedModel.(Model).focusState)
	}
}

func TestRenderSearchComponent(t *testing.T) {
	result := renderSearchComponent(40, true, "Type to search...", false)
	if result == "" {
		t.Error("Expected non-empty result from renderSearchComponent")
	}
	if !contains(result, "Search Section (1)") {
		t.Error("Expected title to contain 'Search Section (1)'")
	}
}

func TestRenderResultComponent(t *testing.T) {
	items := []string{"• Item 1", "• Item 2", "• Item 3"}
	result := renderResultComponent(40, 10, true, items, 0, 0)
	if result == "" {
		t.Error("Expected non-empty result from renderResultComponent")
	}
	if !contains(result, "Result Section (2)") {
		t.Error("Expected title to contain 'Result Section (2)'")
	}
}

// --- Scroll & selection tests ---

func setupResultModel(items []string, visibleCount int) Model {
	m := NewModel()
	m.resultItems = items
	m.visibleCount = visibleCount
	m.focusState = ResultFocused
	return m
}

func TestScrollDownAutoScroll(t *testing.T) {
	items := make([]string, 10)
	for i := range items {
		items[i] = "item"
	}
	m := setupResultModel(items, 3)
	// Navigate down past visible area
	for i := 0; i < 4; i++ {
		res, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
		m = res.(Model)
	}
	if m.selectedIdx != 4 {
		t.Errorf("Expected selectedIdx=4, got %d", m.selectedIdx)
	}
	if m.scrollOffset != 2 {
		t.Errorf("Expected scrollOffset=2, got %d", m.scrollOffset)
	}
}

func TestScrollUpAutoScroll(t *testing.T) {
	items := make([]string, 10)
	for i := range items {
		items[i] = "item"
	}
	m := setupResultModel(items, 3)
	m.selectedIdx = 5
	m.scrollOffset = 5

	res, _ := m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m = res.(Model)
	if m.selectedIdx != 4 {
		t.Errorf("Expected selectedIdx=4, got %d", m.selectedIdx)
	}
	if m.scrollOffset != 4 {
		t.Errorf("Expected scrollOffset=4, got %d", m.scrollOffset)
	}
}

func TestBoundaryTopNoWrap(t *testing.T) {
	items := make([]string, 5)
	m := setupResultModel(items, 3)
	m.selectedIdx = 0

	res, _ := m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m = res.(Model)
	if m.selectedIdx != 0 {
		t.Errorf("Expected selectedIdx to stay at 0, got %d", m.selectedIdx)
	}
}

func TestBoundaryBottomNoWrap(t *testing.T) {
	items := make([]string, 5)
	m := setupResultModel(items, 5)
	m.selectedIdx = 4

	res, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = res.(Model)
	if m.selectedIdx != 4 {
		t.Errorf("Expected selectedIdx to stay at 4, got %d", m.selectedIdx)
	}
}

func TestScrollbarThumbSizeLargeList(t *testing.T) {
	items := make([]string, 100)
	for i := range items {
		items[i] = "item"
	}
	result := renderResultComponent(40, 12, false, items, 0, 0)
	if !contains(result, "█") {
		t.Error("Expected scrollbar thumb '█' in output for large list")
	}
}

func TestScrollbarAbsentWhenFits(t *testing.T) {
	items := []string{"a", "b", "c"}
	result := renderResultComponent(40, 10, false, items, 0, 0)
	if contains(result, "█") || contains(result, "│") {
		t.Error("Expected no scrollbar when all items fit in view")
	}
}

func TestVisibleCountUpdatesOnResize(t *testing.T) {
	m := NewModel()
	res, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = res.(Model)
	// resultHeight = 24 - 3 - 2 = 19; visibleCount = 18
	if m.visibleCount != 18 {
		t.Errorf("Expected visibleCount=18 for height=24, got %d", m.visibleCount)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
