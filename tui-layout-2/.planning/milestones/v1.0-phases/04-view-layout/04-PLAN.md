# Phase 04: View & Layout Rendering - Plan 1

**Wave:** 1
**Depends On:** Phase 2 (Plan 1), Phase 3 (Plan 1)
**Files Modified:**
- internal/ui/model.go (replace debug View with full kanban layout)
**Autonomous:** true

---

## Overview

Implement the complete View function to render a 3-column kanban board layout with dynamic highlighting of focused elements. This phase brings together the Model state (Phase 2) and Visual Styles (Phase 3) to create a fully functional TUI interface.

**Goal:** Transform the debug View into a polished kanban board that renders columns, cards, titles, and navigation feedback using the styles defined in Phase 3.

---

## Requirements Coverage

All 8 requirements from Phase 4 must be satisfied:

- **VIEW-01**: View(m Model) string function renders complete TUI layout
- **VIEW-02**: Each column rendered with appropriate style (active vs inactive)
- **VIEW-03**: Column titles rendered using titleStyle
- **VIEW-04**: Cards rendered with appropriate style (focused vs unfocused)
- **VIEW-05**: Cards stacked vertically using lipgloss.JoinVertical
- **VIEW-06**: Three columns arranged horizontally using lipgloss.JoinHorizontal
- **VIEW-07**: Application title rendered above board using appTitleStyle
- **VIEW-08**: Help bar rendered below board using helpStyle

---

## Tasks

### Task 1: Create Card Rendering Helper

**Active Form:** Creating card rendering helper function

**Description:**
Create a helper function `renderCard(card domain.Card, cardIdx int, colIdx int, focusedCol int, focusedCard int) string` that:
- Takes a card and its position indices
- Checks if the card is focused (colIdx == focusedCol && cardIdx == focusedCard)
- Returns ActiveCardStyle.Render(card.Title) if focused, else CardStyle.Render(card.Title)
- Adds padding around the card title for visual spacing

**Acceptance Criteria:**
- Function compiles without errors
- Returns amber-bordered bold text when card is focused
- Returns gray-bordered regular text when card is not focused
- Card title is centered with padding

**Files Modified:** internal/ui/model.go

**Dependencies:** None

---

### Task 2: Create Column Rendering Helper

**Active Form:** Creating column rendering helper function

**Description:**
Create a helper function `renderColumn(col domain.Column, colIdx int, focusedCol int, focusedCard int) string` that:
- Renders the column title using TitleStyle.Render(col.Title)
- Iterates through col.Cards and calls renderCard for each
- Stacks all cards vertically using lipgloss.JoinVertical(lipgloss.Left, cardStrings...)
- Combines title + cards using lipgloss.JoinVertical(lipgloss.Left, title, cards)
- Applies ActiveColumnStyle if colIdx == focusedCol, else ColumnStyle

**Acceptance Criteria:**
- Function compiles without errors
- Column title renders at top of column (bold, centered)
- All cards in column stack vertically
- Focused column has purple border (#7C3AED)
- Unfocused column has gray border (#245)

**Files Modified:** internal/ui/model.go

**Dependencies:** Task 1 (renderCard helper must exist)

---

### Task 3: Implement Full View Layout

**Active Form:** Implementing full View layout with title, board, and help bar

**Description:**
Replace the debug View implementation with a complete kanban layout:
- Keep the `!m.ready` check that returns "Initializing..."
- Create three sections: app title, board, help bar
- Render app title: AppTitleStyle.Render("KANBAN BOARD")
- Render board: iterate through m.columns, call renderColumn for each, join horizontally with lipgloss.JoinHorizontal(lipgloss.Top, columnStrings...)
- Render help bar: HelpStyle.Render("←/→: Move column | ↑/↓: Move card | q: Quit")
- Combine all sections with lipgloss.JoinVertical(lipgloss.Left, title, board, help)
- Return tea.NewView(content)

**Acceptance Criteria:**
- View function compiles without errors
- Application runs without errors
- 3 columns render side by side filling terminal width
- Active column has purple border
- Active card has amber border + bold text
- Application title appears centered above board
- Help bar appears centered below board
- Keyboard navigation updates visual highlighting in real-time

**Files Modified:** internal/ui/model.go

**Dependencies:** Task 2 (renderColumn helper must exist)

---

## Verification Criteria

### Build Verification
```bash
go build ./cmd/kanban  # Must pass
go run ./cmd/kanban    # Must start without errors
```

### Visual Verification
1. **Layout Structure:**
   - [ ] Application title appears at top ("KANBAN BOARD")
   - [ ] 3 columns render side-by-side (To Do, In Progress, Done)
   - [ ] Cards stack vertically within each column
   - [ ] Help bar appears at bottom with navigation hints

2. **Focus Indicators:**
   - [ ] Active column has purple border (#7C3AED)
   - [ ] Active card has amber border (#F59E0B) + bold text
   - [ ] Inactive columns have gray border (#245)
   - [ ] Inactive cards have gray border + regular text

3. **Navigation Feedback:**
   - [ ] Left/right arrow keys change focused column (border color updates)
   - [ ] Up/down arrow keys change focused card (border color + bold updates)
   - [ ] hjkl keys work identically to arrow keys
   - [ ] No visual corruption on navigation

### Edge Case Handling
- [ ] No panic when navigating to empty columns (defensive checks exist)
- [ ] No panic when column has no cards (defensive checks exist)
- [ ] Terminal resize triggers window size update without breaking layout

---

## Must Haves (Goal-Backward Verification)

Derived from Phase 4 goal: "Implement the View function to render the complete 3-column kanban board layout with dynamic highlighting of focused elements."

**Critical Success Factors:**

1. **View Function Works:** The View() method must return tea.View with tea.NewView(content) and render without errors
2. **3-Column Layout:** lipgloss.JoinHorizontal must arrange 3 columns side by side
3. **Card Stacking:** lipgloss.JoinVertical must stack cards within each column
4. **Style Application:** ActiveColumnStyle, ActiveCardStyle must be applied conditionally based on focusedCol/focusedCard state
5. **Visual Feedback:** Purple border for active column, amber border + bold for active card must be visible
6. **Real-Time Updates:** Navigation must trigger visual changes immediately (no rebuild required)
7. **Complete Layout:** Title + board + help bar all render in correct positions

**Non-Negotiable:**
- Must use bubbletea v2 API (View() returns tea.View, not string)
- Must use lipgloss v2 API (charm.land/lipgloss/v2)
- Must use exported styles from styles.go (ColumnStyle, ActiveColumnStyle, CardStyle, ActiveCardStyle, TitleStyle, AppTitleStyle, HelpStyle)
- Must not introduce new dependencies
- Must not break existing keyboard navigation from Phase 2

---

## Implementation Notes

### Style Import
The View function should use the exported styles defined in styles.go:
```go
// Use these exported variables:
ColumnStyle       // lipgloss.Style for inactive columns
ActiveColumnStyle // lipgloss.Style for focused column (purple border)
CardStyle         // lipgloss.Style for inactive cards
ActiveCardStyle   // lipgloss.Style for focused card (amber border + bold)
TitleStyle        // lipgloss.Style for column titles
AppTitleStyle     // lipgloss.Style for "KANBAN BOARD" title
HelpStyle         // lipgloss.Style for help bar text
```

### Lipgloss Layout Pattern
```go
// Vertical stacking (cards in column, or title+board+help)
content := lipgloss.JoinVertical(lipgloss.Left, item1, item2, item3)

// Horizontal arrangement (columns side by side)
board := lipgloss.JoinHorizontal(lipgloss.Top, col1, col2, col3)
```

### Conditional Style Application
```go
// For columns
var styledColumn string
if colIdx == focusedCol {
    styledColumn = ActiveColumnStyle.Render(columnContent)
} else {
    styledColumn = ColumnStyle.Render(columnContent)
}

// For cards
var styledCard string
if colIdx == focusedCol && cardIdx == focusedCard {
    styledCard = ActiveCardStyle.Render(card.Title)
} else {
    styledCard = CardStyle.Render(card.Title)
}
```

### Layout Hierarchy
```
View() returns:
├── App Title (AppTitleStyle)
├── Board (JoinHorizontal)
│   ├── Column 0 (ActiveColumnStyle or ColumnStyle)
│   │   ├── Title (TitleStyle)
│   │   └── Cards (JoinVertical)
│   │       ├── Card 0 (ActiveCardStyle or CardStyle)
│   │       ├── Card 1
│   │       └── Card 2
│   ├── Column 1
│   └── Column 2
└── Help Bar (HelpStyle)
```

### Defensive Programming
- Always check `len(m.columns) > 0` before accessing columns
- Always check `len(col.Cards) > 0` before accessing cards
- Always check index bounds before accessing focusedCol/focusedCard
- Handle empty columns gracefully (render title only)

---

## Testing Strategy

### Manual Testing Steps
1. Build and run: `go run ./cmd/kanban`
2. Verify layout renders: 3 columns visible side by side
3. Press left/right arrows: verify focused column border changes color
4. Press up/down arrows: verify focused card border changes color and becomes bold
5. Verify hjkl keys work identically to arrow keys
6. Press 'q': verify application exits cleanly

### Expected Visual Output
```
        KANBAN BOARD

┌─────────────────┬─────────────────┬─────────────────┐
│     To Do       │  In Progress    │      Done       │
├─────────────────┼─────────────────┼─────────────────┤
│┌───────────────┐│┌───────────────┐│┌───────────────┐│
││  Fix login bug│││Refactor auth │││Setup CI pipe ││
│└───────────────┘│└───────────────┘│└───────────────┘│
│┌───────────────┐│┌───────────────┐│┌───────────────┐│
││Write unit test│││Code review PR │││  Deploy v1.0  ││
│└───────────────┘│└───────────────┘│└───────────────┘│
│┌───────────────┐││               │││               ││
││  Update README│││               │││               ││
│└───────────────┘││               │││               ││
└─────────────────┴─────────────────┴─────────────────┘

  ←/→: Move column | ↑/↓: Move card | q: Quit
```

*(Note: Active elements would have colored borders)*

---

## Dependencies

### Internal Dependencies
- **Phase 2 (Plan 1)**: Model structure with columns, focusedCol, focusedCard state must exist
- **Phase 3 (Plan 1)**: All 7 styles (ColumnStyle, ActiveColumnStyle, CardStyle, ActiveCardStyle, TitleStyle, AppTitleStyle, HelpStyle) must be defined

### External Dependencies
- bubbletea v2.0.0 (tea.View, tea.NewView)
- lipgloss v2.0.0 (JoinVertical, JoinHorizontal)

---

## Risk Mitigation

### Risk 1: Breaking Existing Navigation
**Mitigation:** Do NOT modify the Update() method. Only replace the View() method implementation.

### Risk 2: Incorrect Style Application
**Mitigation:** Use exported style variables from styles.go, do NOT redefine styles in View function.

### Risk 3: Layout Corruption on Small Terminals
**Mitigation:** Phase 4 uses fixed column width (25 chars) from styles. Responsive sizing deferred to Phase 5. Test at minimum 80×24 terminal.

### Risk 4: Index Out of Bounds
**Mitigation:** Add defensive checks before accessing m.columns[m.focusedCol] and col.Cards[m.focusedCard].

---

## Success Metrics

1. **Build Success:** `go build ./cmd/kanban` completes without errors
2. **Runtime Success:** `go run ./cmd/kanban` starts and displays kanban board
3. **Visual Correctness:** 3 columns render side-by-side with proper styling
4. **Navigation Feedback:** Keyboard input triggers immediate visual updates
5. **Requirements Coverage:** All 8 VIEW requirements satisfied (VIEW-01 through VIEW-08)

---

## Post-Completion Checklist

- [ ] All 8 VIEW requirements marked complete in REQUIREMENTS.md
- [ ] ROADMAP.md updated with Phase 4 status
- [ ] STATE.md updated with Phase 4 completion summary
- [ ] Code committed with descriptive commit message
- [ ] Manual verification completed (layout renders, navigation works)
- [ ] Ready for Phase 5 (Polish & Responsive Layout)

---

**Plan created:** 2026-02-28
**Estimated execution time:** 5-10 minutes
**Complexity:** Low (well-researched, clear patterns from Phase 3 research)
