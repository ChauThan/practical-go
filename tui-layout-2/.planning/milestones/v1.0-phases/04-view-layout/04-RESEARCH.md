# Phase 4: View & Layout Rendering - Research

**Phase:** 4
**Researched:** 2026-02-28
**Tech Stack:** bubbletea v2.0.0, lipgloss v2.0.0

---

## Research Objective

Answer: "What do I need to know to implement a complete View function for a 3-column kanban board using bubbletea v2 and lipgloss v2?"

---

## Bubbletea v2 View API

### View Function Signature

```go
func (m Model) View() tea.View {
    // ... build view content
    return tea.NewView(content)
}
```

**Key changes from v1:**
- View() returns `tea.View` struct, not `string`
- Content wrapped with `tea.NewView(content)`
- tea.View is a struct with rendering configuration

### View Pattern

```go
func (m Model) View() tea.View {
    var content string

    if !m.ready {
        content = "Initializing..."
    } else {
        // Build actual TUI layout
        content = renderBoard(m)
    }

    return tea.NewView(content)
}
```

---

## Lipgloss Layout Functions

### JoinVertical - Stack Content Vertically

```go
import "charm.land/lipgloss/v2"

// Stack multiple strings vertically
var columnContent = lipgloss.JoinVertical(
    lipgloss.Left,   // Alignment
    titleText,       // Top
    card1Text,       // Middle items
    card2Text,
    card3Text,
)
```

**Use case:** Stack cards within a column

### JoinHorizontal - Arrange Content Horizontally

```go
// Arrange columns side by side
var boardLayout = lipgloss.JoinHorizontal(
    lipgloss.Top,    // Vertical alignment
    column1Content,  // Left
    column2Content,  // Middle
    column3Content,  // Right
)
```

**Use case:** Arrange 3 columns horizontally

---

## Conditional Style Application

### Pattern: Apply Style Based on Focus State

```go
// For columns
var columnRendered string
if i == m.focusedCol {
    columnRendered = activeColumnStyle.Render(columnContent)
} else {
    columnRendered = columnStyle.Render(columnContent)
}

// For cards
var cardRendered string
if i == m.focusedCard && colIdx == m.focusedCol {
    cardRendered = activeCardStyle.Render(cardText)
} else {
    cardRendered = cardStyle.Render(cardText)
}
```

### Helper Function Pattern

```go
func renderColumn(col domain.Column, colIdx int, focusedCol int, focusedCard int) string {
    var cards []string

    // Render each card with appropriate style
    for i, card := range col.Cards {
        cardText := "  " + card.Title + "  "
        if colIdx == focusedCol && i == focusedCard {
            cards = append(cards, activeCardStyle.Render(cardText))
        } else {
            cards = append(cards, cardStyle.Render(cardText))
        }
    }

    // Stack cards vertically
    cardsContent := lipgloss.JoinVertical(lipgloss.Left, cards...)

    // Add column title
    titleContent := titleStyle.Render(col.Title)

    // Combine title + cards
    columnContent := lipgloss.JoinVertical(lipgloss.Left, titleContent, cardsContent)

    // Apply column style
    if colIdx == focusedCol {
        return activeColumnStyle.Render(columnContent)
    }
    return columnStyle.Render(columnContent)
}
```

---

## View Layout Structure

### Complete TUI Layout

```
┌─────────────────────────────────────────────┐
│           KANBAN BOARD TITLE                │  ← appTitleStyle
├──────────────┬──────────────┬───────────────┤
│   TO DO      │ IN PROGRESS  │     DONE      │  ← titleStyle for headers
├──────────────┼──────────────┼───────────────┤
│  Card 1      │  Card 4      │   Card 7      │  ← cardStyle / activeCardStyle
│  Card 2      │  Card 5      │   Card 8      │
│  Card 3      │  Card 6      │   Card 9      │
└──────────────┴──────────────┴───────────────┘
│  ←/→: Move column | ↑/↓: Move card | q: Quit │  ← helpStyle
└─────────────────────────────────────────────┘
```

### Implementation Order

1. **Render individual cards** — Apply activeCardStyle vs cardStyle based on focus
2. **Stack cards in columns** — Use JoinVertical within each column
3. **Add column titles** — Prepend titleStyle.Render(col.Title) to each column
4. **Arrange columns horizontally** — Use JoinHorizontal for 3-column layout
5. **Add app title** — Prepend appTitleStyle.Render("KANBAN BOARD")
6. **Add help bar** — Append helpStyle.Render(help text)

---

## String Building Patterns

### Pattern 1: Build String Directly

```go
func (m Model) View() tea.View {
    var content strings.Builder

    // App title
    content.WriteString(appTitleStyle.Render("KANBAN BOARD"))
    content.WriteString("\n\n")

    // Columns
    var columns []string
    for i, col := range m.columns {
        columns = append(columns, renderColumn(col, i, m.focusedCol, m.focusedCard))
    }
    content.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, columns...))
    content.WriteString("\n\n")

    // Help bar
    content.WriteString(helpStyle.Render("←/→: Move column | ↑/↓: Move card | q: Quit"))

    return tea.NewView(content.String())
}
```

### Pattern 2: Build Hierarchically

```go
func (m Model) View() tea.View {
    var sections []string

    // App title
    sections = append(sections, appTitleStyle.Render("KANBAN BOARD"))

    // Board
    var columns []string
    for i, col := range m.columns {
        columns = append(columns, renderColumn(col, i, m.focusedCol, m.focusedCard))
    }
    board := lipgloss.JoinHorizontal(lipgloss.Top, columns...)
    sections = append(sections, board)

    // Help bar
    sections = append(sections, helpStyle.Render("←/→: Move column | ↑/↓: Move card | q: Quit"))

    // Combine all sections
    content := lipgloss.JoinVertical(lipgloss.Left, sections...)

    return tea.NewView(content)
}
```

**Recommendation:** Pattern 2 (hierarchical) is cleaner and more maintainable.

---

## Height/Width Considerations

### Terminal Size Awareness

Model has `width` and `height` fields (updated by tea.WindowSizeMsg):

```go
type Model struct {
    width  int
    height int
    // ... other fields
}
```

**Current approach (Phase 4):**
- Use fixed layout with lipgloss.Width() from styles
- Trust styles to handle spacing (columnStyle has Width(25))
- Focus on visual correctness, not dynamic sizing

**Dynamic sizing (Phase 5):**
- Calculate column widths: `width / 3`
- Handle small terminals with minimum width guards
- Implement responsive reflow

---

## Style Integration

### Import Styles from styles.go

```go
package ui

import (
    "charm.land/lipgloss/v2"
    "strings"
    // ... other imports
)

func (m Model) View() tea.View {
    // Use exported styles:
    // - ColumnStyle, ActiveColumnStyle
    // - CardStyle, ActiveCardStyle
    // - TitleStyle, AppTitleStyle, HelpStyle
}
```

### Style Rendering

```go
// Apply style to text
styledText := ColumnStyle.Render("content")

// Style returns a string that includes ANSI codes
// No need to manually track colors/borders
```

---

## Common Pitfalls

### 1. Forgetting to wrap content in tea.NewView()

**Wrong:**
```go
func (m Model) View() string {
    return content  // v1 API
}
```

**Correct:**
```go
func (m Model) View() tea.View {
    return tea.NewView(content)  // v2 API
}
```

### 2. Not handling empty columns/cards

```go
if len(m.columns) == 0 {
    return tea.NewView("No columns")
}

if len(col.Cards) == 0 {
    // Render empty column with title only
}
```

### 3. Index out of bounds on focus state

```go
// Always check bounds before accessing
if m.focusedCol >= 0 && m.focusedCol < len(m.columns) {
    // Safe to access m.columns[m.focusedCol]
}
```

### 4. Not using JoinVertical/JoinHorizontal correctly

**Wrong:**
```go
cards := strings.Join([]string{card1, card2, card3}, "\n")
```

**Correct:**
```go
cards := lipgloss.JoinVertical(lipgloss.Left, card1, card2, card3)
```

---

## Testing Strategy

### Manual Testing

1. **Visual inspection:** Run `go run ./cmd/kanban` and verify:
   - 3 columns render side by side
   - Active column has purple border
   - Active card has amber border + bold text
   - App title appears at top
   - Help bar appears at bottom

2. **Navigation testing:** Press arrow keys and verify:
   - Left/right changes focused column (border color changes)
   - Up/down changes focused card (border color + bold changes)
   - No visual corruption on navigation

3. **Terminal resize:** Resize terminal and verify:
   - Layout doesn't break (Phase 4: basic test)
   - Content remains readable

### Build Verification

```bash
go build ./cmd/kanban  # Must pass
go run ./cmd/kanban    # Must start without errors
```

---

## Dependencies

No new dependencies required. Using existing:
- bubbletea v2.0.0 (tea.View, tea.NewView)
- lipgloss v2.0.0 (JoinVertical, JoinHorizontal, styles)

---

## Validation Strategy

1. **Compile check:** `go build ./cmd/kanban` must pass
2. **View function returns:** tea.View (not string)
3. **Layout correctness:** 3 columns render side-by-side
4. **Style application:** Active vs inactive elements visually distinct
5. **Navigation feedback:** Focus state changes visible in real-time

---

## Open Questions

None - View implementation is straightforward with lipgloss layout functions and bubbletea v2 API.

---

## RESEARCH COMPLETE

**Status:** Ready for planning

**Key findings:**
- Use bubbletea v2 API: `View() tea.View` with `tea.NewView(content)`
- Use lipgloss.JoinVertical for stacking cards in columns
- Use lipgloss.JoinHorizontal for arranging 3 columns
- Apply styles conditionally based on focusedCol/focusedCard state
- Hierarchical layout pattern (sections → board → columns → cards) is most maintainable
- Handle edge cases (empty columns, index bounds) with defensive checks

---

*Phase: 04-view-layout*
*Research completed: 2026-02-28*
