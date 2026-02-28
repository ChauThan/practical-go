# Phase 5: Polish & Responsive Layout - Plan 1

**Wave:** 1
**Depends on:** Phase 4 (View & Layout Rendering)
**Files Modified:**
- `internal/ui/model.go` (update View to use dynamic widths)
- `internal/ui/styles.go` (update style definitions or remove fixed widths)

**Autonomous:** true
**Estimated Tasks:** 3
**Estimated Time:** 5-10 minutes

---

## Goal

Transform the fixed-width kanban board layout into a responsive layout that adapts gracefully to terminal size changes (80×24, 120×40, 200×50) without visual corruption.

---

## Requirements (from REQUIREMENTS.md)

This plan satisfies the following Phase 5 requirements:
- **RESP-01**: Column widths calculated dynamically from terminal width (width / 3)
- **RESP-02**: Columns stretch to fill available terminal width
- **RESP-03**: Layout respects minimum column width guard for small terminals
- **RESP-04**: Column height fills available vertical space
- **RESP-05**: App title centered horizontally
- **RESP-06**: Layout reflows correctly at 80×24 terminal size without corruption
- **RESP-07**: Layout reflows correctly at 120×40 terminal size
- **RESP-08**: Layout reflows correctly at 200×50 terminal size

---

## Tasks

### Task 1: Add Dynamic Width Calculation Helper

**Description:**
Add a `columnWidth()` helper method to the Model that calculates responsive column width with minimum width guard. This method will be used by View() to determine column width dynamically.

**Acceptance Criteria:**
- [1.1] New method `columnWidth() int` added to Model struct
- [1.2] Method calculates `m.width / 3` for column width
- [1.3] Method enforces minimum width of 25 characters (prevents layout collapse)
- [1.4] Method handles zero/negative width safely (returns default 25)

**Implementation Details:**
```go
// Add to Model in internal/ui/model.go
func (m Model) columnWidth() int {
    if m.width == 0 {
        return 25  // Default for uninitialized state
    }
    columnWidth := m.width / 3
    if columnWidth < 25 {
        return 25  // Minimum width guard
    }
    return columnWidth
}
```

**Files Modified:**
- `internal/ui/model.go` (add columnWidth() method after NewModel())

---

### Task 2: Update View Function to Use Dynamic Widths

**Description:**
Modify the View() function and renderColumn() helper to use dynamic column widths instead of fixed-width styles. Create dynamic styles within View() based on calculated widths.

**Acceptance Criteria:**
- [2.1] renderColumn() helper accepts columnWidth parameter
- [2.2] Column styles created dynamically with Width(columnWidth) in renderColumn()
- [2.3] Title styles created dynamically with Width(columnWidth) for alignment
- [2.4] View() calls m.columnWidth() and passes to renderColumn()
- [2.5] App title centered using lipgloss.Place or dynamic width

**Implementation Details:**

Update renderColumn signature:
```go
func renderColumn(col domain.Column, colIdx int, focusedCol int, focusedCard int, columnWidth int) string
```

Create dynamic styles in renderColumn():
```go
// Dynamic column style based on width
columnStyle := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color(inactiveBorderColor)).
    Padding(0, 1).
    Width(columnWidth)

activeColumnStyle := columnStyle.Copy().
    BorderForeground(lipgloss.Color(activeColumnColor))

// Dynamic title style for proper centering
titleStyle := lipgloss.NewStyle().
    Bold(true).
    Align(lipgloss.Center).
    Width(columnWidth)
```

Update View() to pass columnWidth:
```go
columnWidth := m.columnWidth()
for colIdx, col := range m.columns {
    columnStrings = append(columnStrings, renderColumn(col, colIdx, m.focusedCol, m.focusedCard, columnWidth))
}
```

Center app title:
```go
title := lipgloss.Place(
    m.width, 3,
    lipgloss.Center, lipgloss.Center,
    lipgloss.NewStyle().Bold(true).Render("KANBAN BOARD"),
)
```

**Files Modified:**
- `internal/ui/model.go` (update renderColumn signature and View() function)

---

### Task 3: Manual Verification at Multiple Terminal Sizes

**Description:**
Test the responsive layout at three terminal sizes (80×24, 120×40, 200×50) to verify no visual corruption, proper column width scaling, and successful resize reflow.

**Acceptance Criteria:**
- [3.1] Application builds successfully: `go build ./cmd/kanban`
- [3.2] At 80×24: No clipped borders, no wrapped text, focus indicators visible
- [3.3] At 120×40: Columns fill width evenly, layout balanced
- [3.4] At 200×50: Columns not too wide, content readable
- [3.5] Terminal resize triggers layout reflow without restart
- [3.6] Navigation works correctly at all sizes

**Testing Procedure:**

1. Build verification:
   ```bash
   cd /home/chauthan/projects/practical-go/tui-layout-2
   go build ./cmd/kanban
   ```

2. Test at 80×24:
   - Resize terminal to 80×24
   - Run: `go run ./cmd/kanban`
   - Verify: No clipped borders, title centered, help bar visible
   - Test: Arrow key navigation works

3. Test at 120×40:
   - Resize terminal to 120×40 (should reflow immediately)
   - Verify: Columns wider and fill space evenly
   - Test: Navigation still works

4. Test at 200×50:
   - Resize terminal to 200×50
   - Verify: Columns not excessively wide, content readable

5. Test resize:
   - Resize back to 80×24
   - Verify: No visual corruption during resize

**Files Modified:** None (verification only)

---

## Verification Criteria

### Build Verification
```bash
cd /home/chauthan/projects/practical-go/tui-layout-2
go build ./cmd/kanban
```
**Expected:** Clean build with no errors

### Manual Testing Checklist

**80×24 Terminal:**
- [ ] No clipped borders
- [ ] No wrapped text in cards
- [ ] Focus indicators (purple/amber) visible
- [ ] Title centered horizontally
- [ ] Help bar visible and centered
- [ ] Navigation (arrows + hjkl) works

**120×40 Terminal:**
- [ ] Columns fill width evenly (~40 chars per column)
- [ ] No excessive whitespace
- [ ] All cards visible
- [ ] Layout balanced and professional

**200×50 Terminal:**
- [ ] Columns not too wide (~66 chars per column)
- [ ] Content still readable
- [ ] No horizontal stretch issues

**Resize Testing:**
- [ ] Resize 80→120→200→80
- [ ] Layout reflows without restart
- [ ] No visual corruption during resize
- [ ] Navigation works after resize

---

## Must-Haves (Goal-Backward Verification)

These non-negotiable outcomes must be achieved for Plan 1 to be considered complete:

1. **Dynamic Width Calculation**: View() MUST use `m.width / 3` to calculate column widths (not fixed 25 from styles.go)
2. **Minimum Width Guard**: Column width MUST enforce 25-character minimum to prevent layout collapse
3. **Terminal Reflow**: Layout MUST reflow automatically when terminal resizes (no restart required)
4. **No Visual Corruption**: Layout MUST display correctly at 80×24, 120×40, and 200×50 with no clipped borders or wrapped text
5. **Centered Title**: Application title MUST remain centered regardless of terminal width
6. **Column Height**: Columns MUST fill available vertical space (let terminal handle scrolling if needed)

---

## Notes

### Key Decisions

1. **Dynamic Styles in View()**: Instead of modifying styles.go, create dynamic styles within View()/renderColumn() based on Model.width. This keeps styles responsive and runtime-calculated.

2. **Minimum Width 25**: Established in Phase 3 decision, maintained to ensure readability on small terminals.

3. **No Vertical Truncation**: Let terminal handle vertical scrolling naturally. Focus responsiveness on horizontal (width) adaptation.

4. **Zero Width Handling**: Guard against uninitialized Model (m.width == 0) to return safe default and prevent division errors.

### Dependencies

- **Phase 4 must be complete**: View() function exists and renders fixed-width layout
- **Model.width/height fields**: Already populated by tea.WindowSizeMsg handler in Update()

### Risks & Mitigations

- **Risk**: Division by zero if m.width is 0
  - **Mitigation**: Check for zero width in columnWidth() and return default 25

- **Risk**: Fixed Width() in styles.go conflicts with dynamic widths
  - **Mitigation**: Create new dynamic styles in renderColumn(), don't use exported ColumnStyle/TitleStyle (they have fixed widths)

- **Risk**: lipglass.Place() for centering might not work as expected
  - **Mitigation**: Test centering approach, fallback to simple Width(m.width) with Center alignment if Place() fails

---

**Plan Status:** Ready for execution
**Next Action:** Execute Task 1 (Add columnWidth() helper method)
