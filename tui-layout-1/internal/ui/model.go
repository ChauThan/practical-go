package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type FocusState int

const (
	SearchFocused FocusState = iota
	InputFocused
	ResultFocused
)

type Model struct {
	width        int
	height       int
	focusState   FocusState
	textInput    TextInput
	resultItems  []string
	scrollOffset int
	selectedIdx  int
	visibleCount int
}

type TextInput struct {
	value       string
	cursor      int
	focused     bool
	width       int
	placeholder string
}

func NewTextInput(placeholder string) TextInput {
	return TextInput{
		placeholder: placeholder,
	}
}

func (t TextInput) Focus() TextInput {
	t.focused = true
	return t
}

func (t TextInput) Blur() TextInput {
	t.focused = false
	return t
}

func (t TextInput) SetValue(value string) TextInput {
	t.value = value
	t.cursor = len(value)
	return t
}

func (t TextInput) SetWidth(width int) TextInput {
	t.width = width
	return t
}

func (t TextInput) Update(msg tea.Msg) (TextInput, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok && t.focused {
		switch msg.Type {
		case tea.KeyRunes:
			t.value += string(msg.Runes)
			t.cursor = len(t.value)
		case tea.KeyBackspace:
			if t.cursor > 0 {
				t.value = t.value[:t.cursor-1]
				t.cursor--
			}
		case tea.KeyDelete:
			if t.cursor < len(t.value) {
				t.value = t.value[:t.cursor] + t.value[t.cursor+1:]
			}
		case tea.KeyLeft:
			if t.cursor > 0 {
				t.cursor--
			}
		case tea.KeyRight:
			if t.cursor < len(t.value) {
				t.cursor++
			}
		case tea.KeyHome:
			t.cursor = 0
		case tea.KeyEnd:
			t.cursor = len(t.value)
		}
	}
	return t, nil
}

func (t TextInput) View() string {
	display := t.value

	if len(display) == 0 && !t.focused {
		display = t.placeholder
	} else if t.focused {
		// Add cursor at current position
		if t.cursor >= len(display) {
			display = display + "▋"
		} else {
			display = display[:t.cursor] + "▋" + display[t.cursor:]
		}
	}

	if t.width > 0 && len(display) > t.width {
		display = display[:t.width]
	}

	return display
}

func mockResultItems() []string {
	items := make([]string, 30)
	for i := range items {
		items[i] = fmt.Sprintf("• Result item %d", i+1)
	}
	return items
}

func NewModel() Model {
	ti := NewTextInput("Type to search...")

	return Model{
		focusState:  SearchFocused,
		textInput:   ti,
		resultItems: mockResultItems(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch typed := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = typed.Width
		m.height = typed.Height
		// title line takes 1 row; search=3, help=2
		resultHeight := m.height - 3 - 2
		if resultHeight < 1 {
			resultHeight = 1
		}
		m.visibleCount = resultHeight - 1
		if m.visibleCount < 1 {
			m.visibleCount = 1
		}
	case tea.KeyMsg:
		// Handle special keys before passing to text input
		// Arrow navigation in result section
		if m.focusState == ResultFocused {
			total := len(m.resultItems)
			switch typed.Type {
			case tea.KeyUp:
				if m.selectedIdx > 0 {
					m.selectedIdx--
				}
				if m.selectedIdx < m.scrollOffset {
					m.scrollOffset = m.selectedIdx
				}
				return m, nil
			case tea.KeyDown:
				if m.selectedIdx < total-1 {
					m.selectedIdx++
				}
				if m.visibleCount > 0 && m.selectedIdx >= m.scrollOffset+m.visibleCount {
					m.scrollOffset = m.selectedIdx - m.visibleCount + 1
				}
				return m, nil
			}
		}

		switch typed.String() {
		case "q":
			if m.focusState == InputFocused {
				m.focusState = SearchFocused
				m.textInput = m.textInput.Blur()
				return m, cmd
			} else {
				return m, tea.Quit
			}
		case "1", "2":
			wasInputFocused := m.focusState == InputFocused
			if wasInputFocused {
				m.focusState = SearchFocused
				m.textInput = m.textInput.Blur()
			}
			if typed.String() == "1" {
				m.focusState = SearchFocused
			} else if typed.String() == "2" {
				m.focusState = ResultFocused
			}
			if wasInputFocused {
				return m, nil
			}
		case "i":
			if m.focusState == SearchFocused {
				m.focusState = InputFocused
				m.textInput = m.textInput.Focus()
				m.textInput = m.textInput.SetWidth(m.width - 4)
				return m, nil
			}
		}
	}

	// Only pass messages to text input if it's focused
	if m.focusState == InputFocused {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}
