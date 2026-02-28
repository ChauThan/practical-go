# Phase 2: Model & State Management - Plan

**Created:** 2026-02-28
**Phase:** 2
**Requirements:** MODEL-01 through MODEL-09
**Dependencies:** Phase 1 (Project Scaffold)
**Estimated Waves:** 3

---

## Frontmatter

```yaml
wave: 3
depends_on:
  - phase: 1
    reason: "Requires production-grade directory structure and Elm Architecture foundation from Phase 1"
files_modified:
  - internal/ui/model.go
  - cmd/kanban/main.go
  - internal/domain/types.go (optional field adjustments)
autonomous: true
```

---

## Phase Goal

Populate the Model with mock kanban board data (3 columns with realistic tasks), implement keyboard navigation for column and card selection with boundary checks, and establish focus state management.

**Success Definition:** Users can navigate columns with left/right arrows and cards with up/down arrows, with focus tracked in the Model and no out-of-bounds crashes.

---

## Requirements Coverage

| Requirement | Plan | Status |
|-------------|------|--------|
| MODEL-01 | Plan 1 | Pending |
| MODEL-02 | Plan 1 | Pending |
| MODEL-03 | Plan 1 | Pending |
| MODEL-04 | Phase 1 (existing) | Complete |
| MODEL-05 | Plan 2 | Pending |
| MODEL-06 | Plan 2 | Pending |
| MODEL-07 | Plan 2 | Pending |
| MODEL-08 | Phase 1 (existing) | Complete |
| MODEL-09 | Plan 1 | Pending |

---

## Plan 1: Model Constructor with Mock Data

**Wave:** 1
**Dependencies:** None
**Requirements:** MODEL-01, MODEL-02, MODEL-03, MODEL-09
**Files Modified:** `internal/ui/model.go`, `cmd/kanban/main.go`

### Tasks

#### Task 1: Create NewModel() constructor function
**File:** `internal/ui/model.go`

**Action:** Add a public `NewModel()` function that returns an initialized `Model` struct.

**Implementation Details:**
- Function signature: `func NewModel() Model`
- Initialize `columns` field with 3 `domain.Column` instances
- Initialize `focusedCol` to 0 (first column)
- Initialize `focusedCard` to 0 (first card in column)
- Initialize `width` and `height` to 80 and 24 (default terminal size)
- Initialize `ready` to false

**Acceptance Criteria:**
- [ ] Function compiles without errors
- [ ] Returns Model with all fields initialized
- [ ] focusedCol and focusedCard are non-negative

---

#### Task 2: Populate mock kanban data
**File:** `internal/ui/model.go`

**Action:** Populate the `columns` slice with 3 realistic kanban columns containing task cards.

**Implementation Details:**

**Column 1 - "To Do":**
- Card: "Fix login bug"
- Card: "Write unit tests"
- Card: "Update README"

**Column 2 - "In Progress":**
- Card: "Refactor auth module"
- Card: "Code review PR #42"

**Column 3 - "Done":**
- Card: "Setup CI pipeline"
- Card: "Deploy v1.0"

Use `domain.Column{Title: "...", Cards: []domain.Card{...}}` syntax.

**Acceptance Criteria:**
- [ ] Exactly 3 columns created
- [ ] Column titles are "To Do", "In Progress", "Done"
- [ ] Each column has at least 2 cards
- [ ] All cards have realistic task titles
- [ ] Code compiles without errors

---

#### Task 3: Update main.go to use NewModel()
**File:** `cmd/kanban/main.go`

**Action:** Replace the empty Model declaration with a call to `NewModel()`.

**Current code (lines 11-13):**
```go
// Create initial model
// Note: Model fields will be initialized in Phase 2
var model ui.Model
```

**Replace with:**
```go
// Create initial model with mock data
model := ui.NewModel()
```

**Acceptance Criteria:**
- [ ] main.go compiles without errors
- [ ] `go run ./cmd/kanban` starts application without crashes
- [ ] Model is initialized with mock data on startup

---

### Verification Criteria

1. `go build ./cmd/kanban` succeeds
2. `go run ./cmd/kanban` starts and displays "Press 'q' to quit" (View still uses placeholder text)
3. No runtime panics related to nil slices or uninitialized fields
4. Model can be inspected (if adding debug logging) and contains 3 columns with cards

---

## Plan 2: Keyboard Navigation Implementation

**Wave:** 2
**Dependencies:** Plan 1 (Model must have data to navigate)
**Requirements:** MODEL-05, MODEL-06, MODEL-07
**Files Modified:** `internal/ui/model.go`

### Tasks

#### Task 1: Implement column navigation (left/right arrows)
**File:** `internal/ui/model.go`

**Action:** Extend the `Update()` method's `tea.KeyMsg` switch to handle left/right arrow keys and 'h'/'l' vim keys.

**Implementation Details:**

Add cases to the existing `switch msg.String()` block (after line 28):

```go
case "left", "h":
    if m.focusedCol > 0 {
        m.focusedCol--
        m.focusedCard = 0 // Reset to top card when changing columns
    }
case "right", "l":
    if m.focusedCol < len(m.columns)-1 {
        m.focusedCol++
        m.focusedCard = 0 // Reset to top card when changing columns
    }
```

**Acceptance Criteria:**
- [ ] Left arrow and 'h' move focus to previous column (stop at column 0)
- [ ] Right arrow and 'l' move focus to next column (stop at last column)
- [ ] focusedCard resets to 0 when changing columns
- [ ] No out-of-bounds access when columns slice is empty
- [ ] Compiles without errors

---

#### Task 2: Implement card navigation (up/down arrows)
**File:** `internal/ui/model.go`

**Action:** Extend the `Update()` method to handle up/down arrow keys and 'j'/'k' vim keys.

**Implementation Details:**

Add cases to the existing `switch msg.String()` block:

```go
case "up", "k":
    if m.focusedCard > 0 {
        m.focusedCard--
    }
case "down", "j":
    currentColumn := m.columns[m.focusedCol]
    if m.focusedCard < len(currentColumn.Cards)-1 {
        m.focusedCard++
    }
```

**Acceptance Criteria:**
- [ ] Up arrow and 'k' move focus to previous card (stop at card 0)
- [ ] Down arrow and 'j' move focus to next card (stop at last card)
- [ ] No out-of-bounds access when Cards slice is empty
- [ ] Navigation respects focusedCol (only moves cards within current column)
- [ ] Compiles without errors

---

#### Task 3: Add boundary check guards
**File:** `internal/ui/model.go`

**Action:** Ensure all navigation has defensive bounds checking to prevent panics.

**Implementation Details:**

1. Add a guard clause at the start of card navigation:
```go
// Before up/down cases, add:
if len(m.columns) == 0 {
    return m, nil
}
```

2. Wrap card navigation in column-specific bounds check:
```go
currentCol := m.columns[m.focusedCol]
if len(currentCol.Cards) == 0 {
    return m, nil
}
```

**Acceptance Criteria:**
- [ ] Navigation is safe when columns slice is empty
- [ ] Navigation is safe when a column has no cards
- [ ] No index out of range panics in any navigation scenario
- [ ] Code compiles without errors

---

### Verification Criteria

1. `go build ./cmd/kanban` succeeds
2. Application runs without crashes on arrow key presses
3. Left/right arrows change focusedCol (can verify with debug logging or View output)
4. Up/down arrows change focusedCard
5. Navigation stops at boundaries (no wrap-around)
6. No panic messages in terminal during navigation

---

## Plan 3: Focus State Validation

**Wave:** 3
**Dependencies:** Plan 2 (Navigation must work to validate focus state)
**Requirements:** MODEL-05 (partial), MODEL-06 (partial), MODEL-07 (partial)
**Files Modified:** `internal/ui/model.go`, `internal/ui/view.go` (temporary debug output)

### Tasks

#### Task 1: Add debug View output for focus state
**File:** `internal/ui/model.go`

**Action:** Temporarily modify the `View()` method to display current focus state for verification.

**Implementation Details:**

Replace the existing `View()` method content (lines 40-47) with:

```go
func (m Model) View() tea.View {
    var content string
    if !m.ready {
        content = "Initializing..."
    } else {
        // Debug output to verify navigation works
        focusedCol := "To Do"
        focusedCard := "None"
        if len(m.columns) > 0 && m.focusedCol < len(m.columns) {
            focusedCol = m.columns[m.focusedCol].Title
            if len(m.columns[m.focusedCol].Cards) > 0 && m.focusedCard < len(m.columns[m.focusedCol].Cards) {
                focusedCard = m.columns[m.focusedCol].Cards[m.focusedCard].Title
            }
        }
        content = fmt.Sprintf(
            "Focused Column: %s (index %d)\nFocused Card: %s (index %d)\n\nUse arrow keys or hjkl to navigate. Press 'q' to quit.",
            focusedCol, m.focusedCol, focusedCard, m.focusedCard,
        )
    }
    return tea.NewView(content)
}
```

Add import for "fmt" at the top of the file if not already present.

**Acceptance Criteria:**
- [ ] View displays current focusedCol index and column title
- [ ] View displays current focusedCard index and card title
- [ ] Output updates in real-time as arrow keys are pressed
- [ ] fmt package imported correctly
- [ ] Compiles without errors

---

#### Task 2: Test all navigation scenarios
**Action:** Manual testing of all navigation paths.

**Test Scenarios:**

1. **Column Navigation:**
   - [ ] Press left arrow at column 0 → focusedCol stays at 0
   - [ ] Press right arrow at column 0 → focusedCol becomes 1
   - [ ] Press right arrow at column 2 → focusedCol stays at 2
   - [ ] Press 'h' and 'l' keys → same behavior as arrows

2. **Card Navigation:**
   - [ ] Press up arrow at card 0 → focusedCard stays at 0
   - [ ] Press down arrow at card 0 → focusedCard becomes 1
   - [ ] Press down arrow at last card → focusedCard stays at last card
   - [ ] Press 'j' and 'k' keys → same behavior as arrows

3. **Cross-Column Navigation:**
   - [ ] Navigate to column 2, card 1
   - [ ] Press left arrow → focusedCol becomes 1, focusedCard resets to 0
   - [ ] Navigate to column 1, card 2
   - [ ] Press right arrow → focusedCol becomes 2, focusedCard resets to 0

4. **Boundary Safety:**
   - [ ] Rapid arrow key mashing → no crashes
   - [ ] Navigation when column has 0 cards → no crashes
   - [ ] All combinations work correctly

**Acceptance Criteria:**
- [ ] All test scenarios pass without crashes
- [ ] Focus state displayed correctly in View
- [ ] Navigation feels predictable and natural

---

### Verification Criteria

1. All navigation test scenarios pass
2. No runtime panics or crashes during testing
3. Focus state accurately reflects user input
4. View updates show correct column/card titles and indices
5. Navigation stops at edges as designed

---

## Must-Haves (Goal-Backward Verification)

Derived from phase goal: "Users can navigate columns and cards with keyboard, focus tracked in Model, no crashes."

### Critical Success Factors

1. **Must-Have:** Model initialized with 3 columns containing mock task data
   - Why: Without data, there's nothing to navigate
   - Verified by: Plan 1, Task 2 (mock data population)

2. **Must-Have:** Left/right arrows change focused column with bounds checking
   - Why: Core navigation requirement
   - Verified by: Plan 2, Task 1 (column navigation)

3. **Must-Have:** Up/down arrows change focused card with bounds checking
   - Why: Core navigation requirement
   - Verified by: Plan 2, Task 2 (card navigation)

4. **Must-Have:** No out-of-bounds panics in any navigation scenario
   - Why: Application stability is non-negotiable
   - Verified by: Plan 2, Task 3 (boundary guards) + Plan 3, Task 2 (testing)

5. **Must-Have:** Focus state (focusedCol, focusedCard) accurately tracked
   - Why: Focus tracking enables Phase 4 visual highlighting
   - Verified by: Plan 3, Task 1 (debug View output) + Plan 3, Task 2 (testing)

---

## Execution Notes

### Wave Order
1. **Wave 1:** Plans 1 (Model constructor with mock data)
2. **Wave 2:** Plans 2 (Keyboard navigation implementation)
3. **Wave 3:** Plans 3 (Focus state validation with debug View)

### Integration Points
- **NewModel() called from main.go** — Plan 1, Task 3
- **Update() switch statement extension** — Plan 2, Tasks 1 and 2 add cases to existing switch
- **View() temporary debug output** — Plan 3, Task 1 replaces placeholder content

### Testing Strategy
- After Plan 1: Verify app starts with mock data (no crashes)
- After Plan 2: Test arrow keys don't cause panics
- After Plan 3: Comprehensive navigation testing with debug output

### Rollback Considerations
- If navigation causes panics: Add stricter boundary checks in Plan 2, Task 3
- If mock data structure is wrong: Adjust domain types or data initialization in Plan 1, Task 2
- If View debug output is confusing: Simplify display format in Plan 3, Task 1

---

## Post-Phase Cleanup

After Phase 2 completes, before Phase 4:

1. **Remove debug View output** — The debug View from Plan 3, Task 1 will be replaced in Phase 4 with proper kanban board rendering
2. **Consider adding cursor marker** — Per CONTEXT.md decision, focused cards should get "* " or "▸ " prefix (will be implemented in Phase 4 View)
3. **Document keyboard controls** — Update README.md with navigation hints (arrow keys, hjkl, q to quit)

---

*Plan created: 2026-02-28*
*Ready for execution via /gsd:execute-phase*
