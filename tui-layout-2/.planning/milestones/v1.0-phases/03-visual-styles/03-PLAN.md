# Phase 3: Visual Styles - Execution Plan

**Phase:** 3
**Created:** 2026-02-28
**Status:** Ready for Execution

---

## Frontmatter

```yaml
wave: 1
depends_on:
  - phase: 1
    reason: Project scaffold and lipgloss dependency must exist
files_modified:
  - internal/ui/styles.go
autonomous: true
```

---

## Phase Goal

Define all visual styles using lipgloss v2 to create a polished, professional terminal UI appearance with clear visual feedback for focused states (columns and cards).

---

## Requirements Coverage

All 7 phase requirements must be satisfied:

- **STYLE-01:** columnStyle provides border, padding, and minimum width for columns
- **STYLE-02:** activeColumnStyle highlights focused column with distinct border color (#7C3AED purple)
- **STYLE-03:** cardStyle provides subtle border and padding for task cards
- **STYLE-04:** activeCardStyle highlights focused card with distinct border color (#F59E0B amber)
- **STYLE-05:** titleStyle renders bold, centered column headers
- **STYLE-06:** appTitleStyle renders bold, centered application title
- **STYLE-07:** helpStyle renders dimmed text for footer help bar

---

## Must-Haves (Goal-Backward Verification)

**Goal:** Production-ready style definitions for Phase 4 View rendering

**Verification:**
1. [ ] All 7 styles compile without errors in `internal/ui/styles.go`
2. [ ] Styles are exported (capitalized) for use by Phase 4 View code
3. [ ] columnStyle and activeColumnStyle are visually distinct (different border colors)
4. [ ] cardStyle and activeCardStyle are visually distinct (different border colors)
5. [ ] Color constants defined for maintainability (active colors)
6. [ ] `go build ./internal/ui` passes with no errors
7. [ ] No hardcoded styles exist outside of styles.go (single source of truth)

---

## Tasks

### Task 1: Define Color Constants and Base Styles

**ID:** 03-01
**Type:** auto
**Estimate:** 15 minutes

**Description:**
Establish color constants and implement column/card base styles (inactive state) to provide the foundation for visual hierarchy.

**Files:**
- `/home/chauthan/projects/practical-go/tui-layout-2/internal/ui/styles.go`

**Action:**
Replace the placeholder styles.go content with:
1. Define package-level color constants:
   - `activeColumnColor` = `#7C3AED` (purple for focused column)
   - `activeCardColor` = `#F59E0B` (amber for focused card)
   - `inactiveBorderColor` = `245` (gray for unfocused elements)
2. Implement `columnStyle`:
   - Use `lipgloss.NormalBorder()` for standard borders
   - Set `BorderForeground(lipgloss.Color(inactiveBorderColor))`
   - Add `Padding(0, 1)` for horizontal spacing
   - Set `Width(25)` minimum width (fits 80×24 terminal: 25×3=75 chars)
3. Implement `cardStyle`:
   - Use `lipgloss.NormalBorder()`
   - Set `BorderForeground(lipgloss.Color(inactiveBorderColor))`
   - Add `Padding(0)` to save vertical space (cards stack densely)
   - Do NOT set width (cards fill column width dynamically)

DO NOT create active/focused styles yet (Task 2). DO NOT create text styles yet (Task 3).

**Verify:**
```xml
<verify>
  <automated>cd /home/chauthan/projects/practical-go/tui-layout-2 && go build ./internal/ui</automated>
</verify>
```

**Done:**
- `go build ./internal/ui` completes with no errors
- Color constants exist as package-level vars
- columnStyle has normal border + padding + width
- cardStyle has normal border + padding (no width)

---

### Task 2: Implement Active/Focused State Styles

**ID:** 03-02
**Type:** auto
**Estimate:** 10 minutes
**Depends on:** 03-01

**Description:**
Create focused state styles (activeColumnStyle, activeCardStyle) that provide clear visual feedback for keyboard navigation.

**Files:**
- `/home/chauthan/projects/practical-go/tui-layout-2/internal/ui/styles.go`

**Action:**
After Task 1's color constants, add:
1. Implement `activeColumnStyle`:
   - Copy columnStyle as base: `columnStyle.Copy()`
   - Override border color: `BorderForeground(lipgloss.Color(activeColumnColor))`
   - Keep same padding/width as columnStyle (only color differs)
2. Implement `activeCardStyle`:
   - Copy cardStyle as base: `cardStyle.Copy()`
   - Override border color: `BorderForeground(lipgloss.Color(activeCardColor))`
   - Add `Bold(true)` to make focused card text stand out
   - Keep same padding as cardStyle

DO NOT modify Task 1 styles. Only add new active styles. DO NOT create text styles yet (Task 3).

**Verify:**
```xml
<verify>
  <automated>cd /home/chauthan/projects/practical-go/tui-layout-2 && go build ./internal/ui</automated>
</verify>
```

**Done:**
- `go build ./internal/ui` completes with no errors
- activeColumnStyle uses purple (#7C3AED) border
- activeCardStyle uses amber (#F59E0B) border + bold text
- Active styles are visually distinct from inactive styles (different colors)

---

### Task 3: Create Text Display Styles

**ID:** 03-03
**Type:** auto
**Estimate:** 10 minutes
**Depends on:** 03-01

**Description:**
Implement text rendering styles for column headers, application title, and help bar to complete the visual design system.

**Files:**
- `/home/chauthan/projects/practical-go/tui-layout-2/internal/ui/styles.go`

**Action:**
After Task 1 and Task 2 styles, add:
1. Implement `titleStyle` (column headers):
   - Use `lipgloss.NewStyle().Bold(true)`
   - Add `Align(lipgloss.Center)` for centered text
   - Add `Width(25)` to match column width
   - NO border (headers appear inside column border)
2. Implement `appTitleStyle` (application title):
   - Use `lipgloss.NewStyle().Bold(true)`
   - Add `Align(lipgloss.Center)` for horizontal centering
   - Set `Width(80)` minimum (expands in Phase 5 responsive layout)
   - Consider `Margin(1, 0)` for spacing above/below title
3. Implement `helpStyle` (footer help text):
   - Use `lipgloss.NewStyle().Faint(true)` for dimmed appearance
   - Add `Align(lipgloss.Center)` for centered text
   - NO width constraint (spans full terminal width)
4. Export ALL styles by capitalizing first letter:
   - Change `var columnStyle` to `var ColumnStyle`
   - Change `var activeColumnStyle` to `var ActiveColumnStyle`
   - Change `var cardStyle` to `var CardStyle`
   - Change `var activeCardStyle` to `var ActiveCardStyle`
   - Change `var titleStyle` to `var TitleStyle`
   - Change `var appTitleStyle` to `var AppTitleStyle`
   - Change `var helpStyle` to `var HelpStyle`

DO NOT modify Task 1 or Task 2 styles (only capitalize names for export).

**Verify:**
```xml
<verify>
  <automated>cd /home/chauthan/projects/practical-go/tui-layout-2 && go build ./internal/ui</automated>
</verify>
```

**Done:**
- `go build ./internal/ui` completes with no errors
- All 7 styles exist and are exported (capitalized)
- TitleStyle is bold + centered
- AppTitleStyle is bold + centered + margin
- HelpStyle is faint/dimmed + centered
- All 7 requirements (STYLE-01 through STYLE-07) satisfied

---

## Success Criteria

Phase 3 is complete when:
1. ✓ All 7 style definitions exist in `/home/chauthan/projects/practical-go/tui-layout-2/internal/ui/styles.go`
2. ✓ `go build ./internal/ui` passes with no errors
3. ✓ Styles are exported (capitalized names) for Phase 4 usage
4. ✓ Color constants defined for maintainability
5. ✓ Active styles use distinct colors (#7C3AED purple, #F59E0B amber)
6. ✓ Inactive styles use grayscale (#245) for subtle appearance
7. ✓ No hardcoded styles outside of styles.go (single source of truth)

---

## Dependencies

**Requires:**
- Phase 1 complete (lipgloss v2.0.0 dependency installed)
- `internal/ui/styles.go` exists (created in Phase 1)

**Enables:**
- Phase 4 (View & Layout Rendering) - styles consumed by View implementation
- Phase 5 (Polish & Responsive Layout) - styles may need dynamic width adjustment

---

## Notes

**Design Decisions:**
- Minimum column width (25 chars) ensures 3 columns fit in 80×24 terminal (75 chars total, 5 chars margin)
- NormalBorder() provides standard TUI appearance (│─┼ characters)
- Purple (#7C3AED) for column focus creates high contrast against gray inactive state
- Amber (#F59E0B) for card focus is distinct from purple, preventing confusion
- Bold text on active card improves accessibility and visual hierarchy
- Faint help text reduces visual noise, keeps focus on board content

**Future Considerations (Phase 5):**
- Column width (25) may become dynamic based on terminal width
- AppTitleStyle width (80) may need expansion for wider terminals
- Consider AdaptiveColor for light/dark terminal theme support

---

**Phase Status:** Ready for execution
**Planned Duration:** ~35 minutes (3 tasks, 10-15 min each)
**Autonomous:** Yes (all tasks auto-type, no checkpoints required)
