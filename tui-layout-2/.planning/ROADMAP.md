# Roadmap: Kanban Board TUI Demo

**Created:** 2026-02-28
**Total Phases:** 5
**Requirements Coverage:** 35/35 (100%)

## Overview

This roadmap transforms the PRD's 5-phase plan into executable phases with clear goals, requirements mapping, and success criteria. Each phase builds incrementally toward a fully functional TUI kanban board demo.

---

## Phase 1: Project Scaffold

**Goal:** Initialize the Go module structure, install dependencies, and create a minimal working bubbletea application.

**Requirements:**
- PROJ-01, PROJ-02, PROJ-03, PROJ-04

**Status:** ✅ Complete (2026-02-28)

**Success Criteria:**
1. ✓ `go mod tidy` runs without errors
2. ✓ `go run ./cmd/kanban` starts the application
3. ✓ Application renders placeholder text ("Initializing...", "Press 'q' to quit.")
4. ✓ Application exits cleanly when 'q' or ctrl+c key is pressed
5. ✓ All stub files exist and compile (main.go, model.go, styles.go, types.go)

**Outcomes:**
- ✓ Working Go module with bubbletea v2.0.0 and lipgloss v2.0.0 dependencies
- ✓ Production-grade directory structure (cmd/, internal/)
- ✓ Elm Architecture Model with Init/Update/View methods
- ✓ Domain layer (Card, Column) separated from UI layer
- ✓ Minimal runnable TUI with quit functionality
- ✓ README.md with project documentation and controls

---

## Phase 2: Model & State Management

**Goal:** Define the application state structure and implement keyboard navigation logic using Elm Architecture's Update pattern.

**Requirements:**
- MODEL-01 through MODEL-09

**Status:** → In Progress (Plan 1/3 Complete - 2026-02-28)

**Plans:**
- Plan 1: ✓ Complete (Model constructor + Navigation + Debug View)
- Plan 2: Pending
- Plan 3: Pending

**Progress:** 4/9 requirements satisfied (MODEL-01, MODEL-02, MODEL-03, MODEL-09)

**Success Criteria:**
1. ✓ Card, Column, and Model structs compile without errors
2. ✓ Left/right arrow keys change focused column (no out-of-bounds crashes)
3. ✓ Up/down arrow keys change focused card within column (no out-of-bounds crashes)
4. ✓ 'q' and ctrl+c keys quit the application
5. ⏳ Terminal resize updates Model's width and height fields
6. ✓ Mock data displays 3 columns: "To Do", "In Progress", "Done" with tasks

**Success Criteria:**
1. ✓ Card, Column, and Model structs compile without errors
2. ✓ Left/right arrow keys change focused column (no out-of-bounds crashes)
3. ✓ Up/down arrow keys change focused card within column (no out-of-bounds crashes)
4. ✓ 'q' and ctrl+c keys quit the application
5. ✓ Terminal resize updates Model's width and height fields
6. ✓ Mock data displays 3 columns: "To Do", "In Progress", "Done" with tasks

**Outcomes:**
- ✓ Complete state model for TUI
- ✓ Working keyboard navigation (column and card selection)
- ✓ Foundation for View rendering in Phase 4
- ✓ Debug View for focus state verification

---

## Phase 3: Visual Styles

**Goal:** Define all visual styles using lipgloss to create a polished, professional terminal UI appearance.

**Requirements:**
- STYLE-01 through STYLE-07

**Status:** ✅ Complete (2026-02-28)

**Plans:**
- Plan 1: ✓ Complete (Color constants + Base styles + Active styles + Text styles)

**Progress:** 7/7 requirements satisfied (STYLE-01, STYLE-02, STYLE-03, STYLE-04, STYLE-05, STYLE-06, STYLE-07)

**Success Criteria:**
1. ✓ All 7 styles are defined as exported constants or variables
2. ✓ columnStyle and activeColumnStyle are visually distinct
3. ✓ cardStyle and activeCardStyle are visually distinct
4. ✓ titleStyle renders bold, centered text
5. ✓ appTitleStyle renders bold, centered application title
6. ✓ helpStyle renders dimmed text
7. ✓ No hardcoded styles outside of styles.go

**Outcomes:**
- ✓ Consistent visual design system
- ✓ Clear visual feedback for focused states
- ✓ Maintainable style definitions in single file
- ✓ Color constants for purple (#7C3AED) and amber (#F59E0B) active states
- ✓ All 7 styles exported for Phase 4 consumption

---

## Phase 4: View & Layout Rendering

**Goal:** Implement the View function to render the complete 3-column kanban board layout with dynamic highlighting of focused elements.

**Requirements:**
- VIEW-01 through VIEW-08

**Success Criteria:**
1. `go run .` renders a 3-column kanban layout in terminal
2. Active column has distinct border color (visual focus indicator)
3. Active card has distinct border color (visual selection indicator)
4. All 3 columns render side-by-side filling terminal width
5. Cards stack vertically within columns
6. Application title appears above board
7. Help bar appears below board with navigation hints
8. Keyboard navigation updates visual highlighting in real-time

**Outcomes:**
- Fully functional visual TUI interface
- Complete Elm Architecture implementation (Model, Update, View)
- Interactive kanban board demonstration

---

## Phase 5: Polish & Responsive Layout

**Goal:** Ensure the layout adapts gracefully to different terminal sizes and handles edge cases without visual corruption.

**Requirements:**
- RESP-01 through RESP-08

**Success Criteria:**
1. Terminal resize triggers layout reflow without application restart
2. No content clipped or borders broken at 80×24 terminal size
3. No content clipped or borders broken at 120×40 terminal size
4. No content clipped or borders broken at 200×50 terminal size
5. Columns are equal width and fill horizontal space
6. Columns fill available vertical space
7. Application title remains centered regardless of terminal width
8. Minimum column width guard prevents layout collapse on small terminals

**Outcomes:**
- Production-ready responsive TUI
- Professional appearance across common terminal sizes
- Robust handling of terminal resize events

---

## Phase Summary

| Phase | Name | Status | Requirements | Success Criteria |
|-------|------|--------|--------------|------------------|
| 1 | Project Scaffold | ✅ Complete | 4/4 | 5/5 |
| 2 | Model & State | → In Progress | 4/9 | 4/6 |
| 3 | Visual Styles | ✅ Complete | 7/7 | 7/7 |
| 4 | View & Layout | ○ Pending | 0/8 | 0/8 |
| 5 | Polish & Responsive | ○ Pending | 0/8 | 0/8 |

**Total:** 15/35 requirements complete | 16/34 success criteria passed

---

## Dependencies Between Phases

- **Phase 1** must complete before any other phase (provides foundation)
- **Phase 2** must complete before **Phase 4** (View depends on Model structure)
- **Phase 3** can be developed in parallel with **Phase 2** (styles are independent)
- **Phase 4** must complete before **Phase 5** (responsiveness requires working layout)

---

## Execution Notes

**Recommended Order:**
1. Phase 1 → foundation
2. Phase 2 + Phase 3 → can be done in parallel (Model and Styles are independent)
3. Phase 4 → requires Phase 2 complete, uses Phase 3 styles
4. Phase 5 → requires Phase 4 complete

**Parallelization Opportunities:**
- Phases 2 and 3 are independent and can be worked simultaneously
- Within Phase 2, different Model aspects can be developed separately
- Within Phase 3, styles can be defined independently

**Verification Strategy:**
- After each phase: run `go run .` and verify success criteria manually
- After Phase 4: full integration test of keyboard navigation and visual updates
- After Phase 5: test at multiple terminal sizes (80×24, 120×40, 200×50)

---
*Roadmap created: 2026-02-28*
*Last updated: 2026-02-28 (Phase 3 marked complete)*
