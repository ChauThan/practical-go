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
	result := renderResultComponent(40, 10, true)
	if result == "" {
		t.Error("Expected non-empty result from renderResultComponent")
	}
	if !contains(result, "Result Section (2)") {
		t.Error("Expected title to contain 'Result Section (2)'")
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
