# Phase 5: Polish & Responsive Layout - Research

**Phase:** 5
**Researched:** 2026-02-28
**Tech Stack:** bubbletea v2.0.0, lipgloss v2.0.0

---

## Research Objective

Answer: "What do I need to know to make the kanban board layout responsive to terminal size changes and handle edge cases without visual corruption?"

---

## Terminal Size Awareness in Bubbletea

### WindowSizeMsg Handling

Bubbletea sends `tea.WindowSizeMsg` when terminal resizes:

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        return m, nil
    // ... other cases
    }
}
```

**Current implementation (Phase 2):**
- Model has `width` and `height` fields
- Update() method already handles tea.WindowSizeMsg
- Fields are updated on terminal resize

**What's missing:** Using these dimensions in View()

---

## Dynamic Width Calculation

### Current Fixed Width Approach (Phase 4)

```go
// In styles.go (Phase 3)
var columnStyle = lipgloss.NewStyle().
    Width(25)  // Fixed width!

// Problem: Doesn't adapt to terminal size
```

### Responsive Width Pattern

```go
// In View(), calculate column width dynamically
columnWidth := m.width / 3

// Apply width to each column
columnStyle := lipgloss.NewStyle().
    Width(columnWidth)

// Or use lipgloss dynamically
styledColumn := lipgloss.NewStyle().
    Width(columnWidth).
    Render(columnContent)
```

### Centering Content

```go
// Calculate margin for centering
margin := (m.width - boardWidth) / 2

// Or use lipgloss alignment
centeredBoard := lipgloss.Place(
    m.width, m.height,
    lipgloss.Center, lipgloss.Center,
    boardContent,
)
```

---

## Minimum Width Guards

### Problem: Small Terminals

If terminal is too narrow (e.g., 60 chars wide), 60/3 = 20 chars per column. This might break borders or clip content.

### Solution: Minimum Column Width

```go
const minColumnWidth = 25  // Minimum readable width

maxColumns := m.width / minColumnWidth
if maxColumns < 1 {
    maxColumns = 1  // Always show at least 1 column
}

// Or: Calculate column width with minimum
columnWidth := m.width / 3
if columnWidth < minColumnWidth {
    columnWidth = minColumnWidth
}

// Result: Total width might exceed terminal, but content is readable
// User scrolls horizontally if needed
}
```

### Alternative: Horizontal Scroll

```go
// If too narrow, enable horizontal scroll
if m.width < minBoardWidth {
    // Clip content, let terminal handle scroll
    // Or show "Terminal too narrow" message
}
```

**Recommendation:** For this demo, use minimum column width guard. Let content exceed terminal width rather than breaking layout.

---

## Vertical Space Handling

### Current Approach (Phase 4)

```go
// View renders content without considering m.height
// Problem: Content might be taller than terminal
```

### Responsive Height Pattern

```go
// Calculate available height for cards
headerHeight := 3  // App title + margins
footerHeight := 2  // Help bar
availableHeight := m.height - headerHeight - footerHeight

// Truncate cards if too tall
if len(col.Cards) > maxVisibleCards {
    visibleCards := col.Cards[:maxVisibleCards]
    // Render visible cards only
}
```

**Alternative:** Let terminal handle scrolling (default behavior)

**Recommendation:** For this demo, let terminal handle vertical scrolling. Focus on horizontal responsiveness.

---

## Terminal Size Testing

### Test Terminal Sizes

| Size | Description | Challenge |
|------|-------------|-----------|
| 80×24 | Minimum supported | Tight fit, test minimum widths |
| 120×40 | Common medium | Ample space |
| 200×50 | Large terminal | Test maximum width handling |

### Manual Testing Approach

```bash
# Resize terminal before running
go run ./cmd/kanban

# Test navigation and visual check:
# - No clipped borders
# - No wrapped text in cards
# - Focus indicators visible
# - Title centered
# - Help bar visible
```

### Automated Verification

```go
// In tests, simulate different terminal sizes
testCases := []struct {
    width  int
    height int
}{
    {80, 24},
    {120, 40},
    {200, 50},
}

for _, tc := range testCases {
    m := Model{width: tc.width, height: tc.height}
    view := m.View()
    // Assert: no layout corruption
}
```

---

## Lipgloss Responsive Patterns

### Dynamic Style Application

```go
// Calculate width based on Model state
func (m Model) getColumnWidth() int {
    columnWidth := m.width / 3
    if columnWidth < 25 {
        return 25  // Minimum
    }
    return columnWidth
}

// Apply in View
func (m Model) View() tea.View {
    columnWidth := m.getColumnWidth()

    // Create dynamic style
    dynamicColumnStyle := lipgloss.NewStyle().
        Width(columnWidth).
        Border(lipgloss.NormalBorder())

    // Render each column with dynamic width
    // ...
}
```

### Width and Height

```go
// lipgloss.Width() sets content width
// lipgloss.Height() sets content height
// lipgloss.MaxWidth() limits maximum width
// lipgloss.MaxHeight() limits maximum height

// Example: Truncate long card titles
truncatedTitle := lipgloss.NewStyle().
    MaxWidth(columnWidth - 4).  // Account for border
    Render(card.Title)
```

---

## Implementation Strategy

### Step 1: Update View() to Use Dynamic Widths

**Current (Phase 4):**
```go
func (m Model) View() tea.View {
    // Uses fixed ColumnStyle.Width(25) from styles.go
}
```

**Proposed (Phase 5):**
```go
func (m Model) View() tea.View {
    // Calculate column width dynamically
    columnWidth := m.width / 3
    if columnWidth < 25 {
        columnWidth = 25
    }

    // Create dynamic styles in View()
    // Don't rely on fixed Width() from styles.go
}
```

### Step 2: Handle Edge Cases

**Empty columns:**
```go
if len(m.columns) == 0 {
    return tea.NewView("No columns to display")
}
```

**Empty cards:**
```go
if len(col.Cards) == 0 {
    // Render column with title only, no cards
}
```

**Zero dimensions (initial state):**
```go
if m.width == 0 || m.height == 0 {
    return tea.NewView("Initializing...")
}
```

### Step 3: Test at Multiple Sizes

**Manual test plan:**
1. Run application at 80×24
2. Resize to 120×40 (should reflow without restart)
3. Resize to 200×50 (should reflow)
4. Resize back to 80×24 (should handle gracefully)
5. Test navigation at each size
6. Check for visual corruption

---

## Layout Reflow on Resize

### Current Behavior

- Update() method captures width/height on tea.WindowSizeMsg
- View() is called automatically after Update()
- **Issue:** View() doesn't use the updated dimensions

### Proposed Fix

```go
func (m Model) View() tea.View {
    // Use m.width and m.height for layout
    // View is called after each Update() on resize
    // Layout reflows automatically

    columnWidth := m.width / 3
    // ... rest of View uses dynamic widths
}
```

**Key insight:** No special handling needed! Just use `m.width` and `m.height` in View(). Bubbletea handles the rest.

---

## Common Pitfalls

### 1. Not Using Model Dimensions in View()

**Wrong:**
```go
func (m Model) View() tea.View {
    // Ignores m.width, m.height
    // Uses fixed styles.Width(25)
}
```

**Correct:**
```go
func (m Model) View() tea.View {
    columnWidth := m.width / 3
    // Use columnWidth for layout
}
```

### 2. Division by Zero

```go
// Safe check
if m.width == 0 {
    return tea.NewView("Initializing...")
}

columnWidth := m.width / 3  // Now safe
```

### 3. Negative Widths After Calculations

```go
columnWidth := m.width / 3 - padding
if columnWidth < 0 {
    columnWidth = minColumnWidth
}
```

### 4. Forgetting to Handle Edge Cases

```go
// Always check for empty slices
if len(m.columns) == 0 {
    return tea.NewView("No columns")
}

if m.focusedCol >= len(m.columns) {
    m.focusedCol = len(m.columns) - 1  // Clamp
}
```

---

## Testing Strategy

### Build Verification

```bash
go build ./cmd/kanban  # Must pass
```

### Manual Testing Checklist

**At 80×24:**
- [ ] No clipped borders
- [ ] No wrapped text
- [ ] Focus indicators visible
- [ ] Title centered
- [ ] Help bar visible
- [ ] Navigation works

**At 120×40:**
- [ ] Columns fill width evenly
- [ ] No excessive whitespace
- [ ] All cards visible
- [ ] Layout balanced

**At 200×50:**
- [ ] Columns not too wide
- [ ] Content readable
- [ ] No horizontal stretch issues

**Resize Testing:**
- [ ] Resize from 80→120→200→80
- [ ] Layout reflows without restart
- [ ] No visual corruption during resize
- [ ] Navigation still works after resize

---

## Dependencies

No new dependencies required. Using existing:
- bubbletea v2.0.0 (tea.WindowSizeMsg handling)
- lipgloss v2.0.0 (dynamic Width(), MaxWidth())

---

## Validation Strategy

1. **Compile check:** `go build ./cmd/kanban` must pass
2. **Dimension usage:** View() must use m.width for calculations
3. **Minimum width:** Columns must not collapse below 25 chars
4. **Manual testing:** Test at 80×24, 120×40, 200×50
5. **Resize testing:** Resize terminal and verify reflow

---

## Open Questions

None - Responsive layout is straightforward with dynamic width calculations and minimum guards.

---

## RESEARCH COMPLETE

**Status:** Ready for planning

**Key findings:**
- Model already has width/height fields updated by tea.WindowSizeMsg
- View() must use m.width for dynamic column calculations
- Set minimum column width (25 chars) to prevent layout collapse
- Use `columnWidth := m.width / 3` with minimum guard
- Test at 3 terminal sizes: 80×24, 120×40, 200×50
- Handle edge cases (zero dimensions, empty columns)
- No special resize handling needed - bubbletea calls View() after Update()

---

*Phase: 05-polish-responsive*
*Research completed: 2026-02-28*
