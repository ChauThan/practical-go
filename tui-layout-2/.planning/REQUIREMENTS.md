# Requirements: Kanban Board TUI Demo

**Defined:** 2026-02-28
**Core Value:** Demonstrate Elm Architecture pattern in Go for building terminal UIs with clear separation of Model, Update, and View.

## v1 Requirements

Requirements for initial release. Each maps to roadmap phases.

### Project Setup

- [x] **PROJ-01**: Go module initialized with name `tui-layout-2`
- [x] **PROJ-02**: bubbletea and lipgloss dependencies installed and available
- [x] **PROJ-03**: Stub files created (main.go, model.go, view.go, styles.go)
- [x] **PROJ-04**: Minimal bubbletea app runs and exits cleanly on 'q' key

### Model & State

- [ ] **MODEL-01**: Card struct defined with Title string field
- [ ] **MODEL-02**: Column struct defined with Title string and Cards []Card fields
- [ ] **MODEL-03**: Model struct defined with columns []Column, focusedCol int, focusedCard int, width int, height int
- [ ] **MODEL-04**: Init() function returns nil (no initial commands)
- [ ] **MODEL-05**: Update handles tea.WindowSizeMsg by updating width and height
- [ ] **MODEL-06**: Update handles left/right arrow keys by changing focusedCol with bounds checking
- [ ] **MODEL-07**: Update handles up/down arrow keys by changing focusedCard with bounds checking
- [ ] **MODEL-08**: Update handles 'q' and ctrl+c keys by returning tea.Quit
- [ ] **MODEL-09**: NewModel() constructor populates 3 columns with static mock task data

### Visual Styles

- [ ] **STYLE-01**: columnStyle provides border, padding, and minimum width for columns
- [ ] **STYLE-02**: activeColumnStyle highlights focused column with distinct border color (e.g., #7C3AED)
- [ ] **STYLE-03**: cardStyle provides subtle border and padding for task cards
- [ ] **STYLE-04**: activeCardStyle highlights focused card with distinct border color (e.g., #F59E0B)
- [ ] **STYLE-05**: titleStyle renders bold, centered column headers
- [ ] **STYLE-06**: appTitleStyle renders bold, centered application title
- [ ] **STYLE-07**: helpStyle renders dimmed text for footer help bar

### View & Layout

- [ ] **VIEW-01**: View(m Model) string function renders complete TUI layout
- [ ] **VIEW-02**: Each column rendered with appropriate style (active vs inactive)
- [ ] **VIEW-03**: Column titles rendered using titleStyle
- [ ] **VIEW-04**: Cards rendered with appropriate style (focused vs unfocused)
- [ ] **VIEW-05**: Cards stacked vertically using lipgloss.JoinVertical
- [ ] **VIEW-06**: Three columns arranged horizontally using lipgloss.JoinHorizontal
- [ ] **VIEW-07**: Application title rendered above board using appTitleStyle
- [ ] **VIEW-08**: Help bar rendered below board using helpStyle

### Responsive Layout

- [ ] **RESP-01**: Column widths calculated dynamically from terminal width (width / 3)
- [ ] **RESP-02**: Columns stretch to fill available terminal width
- [ ] **RESP-03**: Layout respects minimum column width guard for small terminals
- [ ] **RESP-04**: Column height fills available vertical space
- [ ] **RESP-05**: App title centered horizontally
- [ ] **RESP-06**: Layout reflows correctly at 80×24 terminal size without corruption
- [ ] **RESP-07**: Layout reflows correctly at 120×40 terminal size
- [ ] **RESP-08**: Layout reflows correctly at 200×50 terminal size

## v2 Requirements

Deferred to future release. Tracked but not in current roadmap.

(None identified)

## Out of Scope

Explicitly excluded. Documented to prevent scope creep.

| Feature | Reason |
|---------|--------|
| Real kanban logic (drag & drop, task CRUD) | Educational demo focused on TUI architecture, not application features |
| Backend persistence (database, file storage) | Adds complexity unrelated to TUI demonstration |
| Data manipulation (create, update, delete tasks) | Static mock data sufficient for UI demonstration |
| Real-time updates (channels, event streaming) | Out of scope for Elm Architecture demo |
| Mouse interaction | Keyboard navigation demonstrates event handling adequately |
| External API integration | Not relevant for TUI demonstration |
| Authentication/authorization | No user management in demo scope |

## Traceability

Which phases cover which requirements. Updated during roadmap creation.

| Requirement | Phase | Status |
|-------------|-------|--------|
| PROJ-01 | Phase 1 | ✅ Complete |
| PROJ-02 | Phase 1 | ✅ Complete |
| PROJ-03 | Phase 1 | ✅ Complete |
| PROJ-04 | Phase 1 | ✅ Complete |
| MODEL-01 | Phase 2 | Pending |
| MODEL-02 | Phase 2 | Pending |
| MODEL-03 | Phase 2 | Pending |
| MODEL-04 | Phase 2 | Pending |
| MODEL-05 | Phase 2 | Pending |
| MODEL-06 | Phase 2 | Pending |
| MODEL-07 | Phase 2 | Pending |
| MODEL-08 | Phase 2 | Pending |
| MODEL-09 | Phase 2 | Pending |
| STYLE-01 | Phase 3 | Pending |
| STYLE-02 | Phase 3 | Pending |
| STYLE-03 | Phase 3 | Pending |
| STYLE-04 | Phase 3 | Pending |
| STYLE-05 | Phase 3 | Pending |
| STYLE-06 | Phase 3 | Pending |
| STYLE-07 | Phase 3 | Pending |
| VIEW-01 | Phase 4 | Pending |
| VIEW-02 | Phase 4 | Pending |
| VIEW-03 | Phase 4 | Pending |
| VIEW-04 | Phase 4 | Pending |
| VIEW-05 | Phase 4 | Pending |
| VIEW-06 | Phase 4 | Pending |
| VIEW-07 | Phase 4 | Pending |
| VIEW-08 | Phase 4 | Pending |
| RESP-01 | Phase 5 | Pending |
| RESP-02 | Phase 5 | Pending |
| RESP-03 | Phase 5 | Pending |
| RESP-04 | Phase 5 | Pending |
| RESP-05 | Phase 5 | Pending |
| RESP-06 | Phase 5 | Pending |
| RESP-07 | Phase 5 | Pending |
| RESP-08 | Phase 5 | Pending |

**Coverage:**
- v1 requirements: 35 total
- Mapped to phases: 35
- Complete: 4 (PROJ-01 through PROJ-04)
- Unmapped: 0 ✓

---
*Requirements defined: 2026-02-28*
*Last updated: 2026-02-28 after initial definition*
