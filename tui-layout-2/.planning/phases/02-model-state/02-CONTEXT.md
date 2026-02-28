# Phase 2: Model & State Management - Context

**Gathered:** 2026-02-28
**Status:** Ready for planning

<domain>
## Phase Boundary

Populate the Model with mock kanban data (3 columns with tasks), implement keyboard navigation (arrow keys) for column and card selection, and add boundary checks to prevent out-of-bounds crashes.

This phase establishes the state management foundation for the interactive kanban board. Visual rendering comes in Phase 4.

</domain>

<decisions>
## Implementation Decisions

### Navigation Boundaries
- Column navigation: **Stop at edges** — left/right arrows do nothing when at first/last column
- Card navigation: **Stop at edges** — up/down arrows do nothing when at top/bottom card in a column
- No wrap-around behavior — prevents user confusion about current position

### Keyboard Controls
- **Arrow keys**: Left/right for columns, up/down for cards
- **Vim-style keys**: hjkl supported in addition to arrows (h=left, j=down, k=up, l=right)
- **Quit**: 'q' and ctrl+c (already implemented in Phase 1)

### Focus Indication
- **Both visual highlight and cursor marker**:
  - Focused column: distinct border color (to be implemented in Phase 3)
  - Focused card: distinct border color + cursor marker (e.g., "* " or "▸ ") prefix
  - This dual approach makes focus obvious even before colors are added in Phase 3

### Mock Data Content
- **Realistic work tasks** — actual development work items
  - To Do: "Fix login bug", "Write unit tests", "Update README"
  - In Progress: "Refactor auth module", "Code review PR #42"
  - Done: "Setup CI pipeline", "Deploy v1.0"
- Makes the demo feel like a real application

### Claude's Discretion
- Exact cursor marker character choice (*, ▸, →, etc.)
- Whether to show row numbers alongside cards
- Order of fields in Card/Column structs (if adjustments needed)

</decisions>

<code_context>
## Existing Code Insights

### Reusable Assets
- **domain.Card** and **domain.Column** types — already defined in internal/domain/types.go
- **ui.Model struct** — has fields: columns, focusedCol, focusedCard, width, height, ready
- **Elm Architecture pattern** — Init(), Update(msg), View() already established
- **Quit handling** — tea.Quit command already wired to 'q' and ctrl+c
- **Window resize** — tea.WindowSizeMsg already updates width/height fields

### Established Patterns
- **Domain/UI separation** — types in internal/domain/, UI logic in internal/ui/
- **Model.Update pattern** — switch on msg type, return (updatedModel, tea.Cmd)
- **tea.Quit for clean exit** — established pattern for termination
- **tea.WindowSizeMsg** — standard way to handle terminal resize

### Integration Points
- **NewModel() constructor** — needs to be created to initialize Model with mock data
- **Update() method** — extend existing switch statement to handle arrow keys
- **focusedCol/focusedCard** — currently unused fields, need to be updated by key handlers
- **columns slice** — currently empty, needs initialization with 3 mock columns

</code_context>

<specifics>
## Specific Ideas

- "I want the keyboard behavior to feel natural and predictable"
- "Focus should be obvious even before we add colors in Phase 3"
- "The mock data should make it feel like a real kanban board"

</specifics>

<deferred>
## Deferred Ideas

- Drag-and-drop card movement — separate phase
- Card editing (title changes, descriptions) — separate phase
- Keyboard shortcuts beyond navigation (delete, archive, etc.) — future phases
- Mouse interaction — out of scope for keyboard-first TUI

</deferred>

---

*Phase: 02-model-state*
*Context gathered: 2026-02-28*
