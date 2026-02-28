# Phase 1: Project Scaffold - Plan

**Created:** 2026-02-28
**Status:** Ready for Execution
**Wave Count:** 2

---

## Frontmatter

```yaml
wave: 1
depends_on: []
files_modified:
  - go.mod
  - go.sum
  - cmd/kanban/main.go
  - internal/ui/model.go
  - internal/ui/update.go
  - internal/ui/view.go
  - internal/ui/styles.go
  - internal/domain/types.go
  - README.md
autonomous: true
```

---

## Goal

Initialize a production-grade Go project structure with bubbletea and lipgloss dependencies, creating a minimal working TUI application that demonstrates Elm Architecture fundamentals (Model-Init-Update-View separation).

**Success means:** The application compiles, runs, displays "hello" text, and exits cleanly when 'q' is pressed — providing a solid foundation for incremental development in Phases 2-5.

---

## Requirements Coverage

This phase implements the following requirements from REQUIREMENTS.md:

- **PROJ-01**: Go module initialized with name `tui-layout-2`
- **PROJ-02**: bubbletea and lipgloss dependencies installed and available
- **PROJ-03**: Stub files created (main.go, model.go, view.go, styles.go)
- **PROJ-04**: Minimal bubbletea app runs and exits cleanly on 'q' key

---

## Must-Haves (Goal-Backward Verification)

To achieve the phase goal, these outcomes MUST exist:

1. **Go module initialized** — `go.mod` exists with module name `tui-layout-2` and requires bubbletea v2.0.0 and lipgloss v2.0.0
2. **Dependencies fetched** — `go.sum` exists with checksums for all dependencies; `go mod tidy` runs without errors
3. **Production-grade directory structure** — `cmd/kanban/`, `internal/ui/`, `internal/domain/` directories exist (not flat structure)
4. **Elm Architecture stub files** — All four required files (main.go, model.go, update.go, view.go, styles.go, types.go) exist and compile
5. **Minimal TUI runs** — `go run ./cmd/kanban` starts the application, renders placeholder text, and responds to 'q' key to quit
6. **Clean quit mechanism** — Application exits cleanly with `tea.Quit` command on 'q' and ctrl+c keys

---

## Task Breakdown

### Wave 1: Project Initialization (Can start immediately)

#### Task 1.1: Initialize Go Module and Install Dependencies

**Description:** Initialize the Go module with the correct name and install bubbletea v2.0.0 and lipgloss v2.0.0 dependencies. Use the production-grade import path for bubbletea v2.

**Steps:**
1. Run `go mod init tui-layout-2` to initialize the module
2. Run `go get github.com/charmbracelet/bubbletea/v2@v2.0.0` to install bubbletea v2
3. Run `go get github.com/charmbracelet/lipgloss@v2.0.0` to install lipgloss v2
4. Run `go mod tidy` to fetch dependencies and create go.sum
5. Verify `go.mod` contains `require` entries for both dependencies with v2.0.0 versions

**Acceptance Criteria:**
- [ ] `go.mod` file exists with `module tui-layout-2`
- [ ] `go.mod` contains `require github.com/charmbracelet/bubbletea/v2 v2.0.0`
- [ ] `go.mod` contains `require github.com/charmbracelet/lipgloss v2.0.0`
- [ ] `go.sum` file exists with checksums
- [ ] `go mod tidy` runs without errors
- [ ] `go list -m all` shows both dependencies

**Verification:** Run `cat go.mod | grep require` and `ls go.sum`

**Time Estimate:** 5 minutes

---

#### Task 1.2: Create Production-Grade Directory Structure

**Description:** Create the production-grade directory structure following Go standard project layout (cmd/, internal/). This is NOT a flat structure — we're building for maintainability from Day 1.

**Steps:**
1. Create `cmd/kanban/` directory for application entry point
2. Create `internal/ui/` directory for Bubbletea UI layer
3. Create `internal/domain/` directory for business logic types (Card, Column)
4. Verify directory structure exists with `ls -R`

**Acceptance Criteria:**
- [ ] `cmd/kanban/` directory exists
- [ ] `internal/ui/` directory exists
- [ ] `internal/domain/` directory exists
- [ ] Directories are empty (no files yet)

**Verification:** Run `tree -d -L 3` or `find . -type d | grep -E "(cmd|internal)"`

**Time Estimate:** 2 minutes

**Dependencies:** None (can run in parallel with Task 1.1)

---

### Wave 2: Elm Architecture Implementation (After Wave 1 completes)

#### Task 1.3: Create Domain Types Stub

**Description:** Create the `internal/domain/types.go` file with minimal Card and Column struct definitions. These are pure Go structs with no Bubbletea dependencies — this is domain layer separation.

**File:** `/home/chauthan/projects/practical-go/tui-layout-2/internal/domain/types.go`

**Implementation:**
```go
package domain

// Card represents a task card in the kanban board
type Card struct {
    Title string
}

// Column represents a kanban column containing cards
type Column struct {
    Title  string
    Cards  []Card
}
```

**Acceptance Criteria:**
- [ ] `internal/domain/types.go` file exists
- [ ] Package is `domain`
- [ ] Card struct exists with Title field
- [ ] Column struct exists with Title and Cards fields
- [ ] File compiles without errors (`go build ./internal/domain`)

**Verification:** Run `go build ./internal/domain`

**Time Estimate:** 3 minutes

**Dependencies:** Task 1.2 (directory structure must exist)

---

#### Task 1.4: Create Bubbletea Model Stub

**Description:** Create the `internal/ui/model.go` file with the Model struct implementing bubbletea's Model interface (Init, Update, View methods). Use the correct v2 API signatures.

**File:** `/home/chauthan/projects/practical-go/tui-layout-2/internal/ui/model.go`

**Implementation:**
```go
package ui

import (
    tea "charm.land/bubbletea/v2"
    "tui-layout-2/internal/domain"
)

// Model holds the application state
type Model struct {
    columns    []domain.Column
    focusedCol int
    focusedCard int
    width      int
    height     int
    ready      bool
}

// Init returns the initial command
func (m Model) Init() tea.Cmd {
    return nil
}

// Update handles incoming messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.ready = true
    }
    return m, nil
}

// View renders the UI
func (m Model) View() string {
    if !m.ready {
        return "Initializing..."
    }
    return "Press 'q' to quit.\n"
}
```

**Acceptance Criteria:**
- [ ] `internal/ui/model.go` file exists
- [ ] Package is `ui`
- [ ] Model struct exists with columns, focusedCol, focusedCard, width, height, ready fields
- [ ] Init() method returns `tea.Cmd` (nil)
- [ ] Update() method has correct v2 signature: `(tea.Msg) (tea.Model, tea.Cmd)`
- [ ] Update() handles tea.KeyMsg for 'q' and 'ctrl+c' with tea.Quit
- [ ] Update() handles tea.WindowSizeMsg
- [ ] View() method returns string
- [ ] Import path uses `charm.land/bubbletea/v2` (NOT github.com/charmbracelet/bubbletea)
- [ ] File compiles without errors

**Verification:** Run `go build ./internal/ui`

**Time Estimate:** 5 minutes

**Dependencies:** Task 1.1 (bubbletea installed), Task 1.3 (domain types exist)

---

#### Task 1.5: Create Styles Stub File

**Description:** Create the `internal/ui/styles.go` file with placeholder lipgloss style definitions. This file will be expanded in Phase 3, but we create the stub now for compilation.

**File:** `/home/chauthan/projects/practical-go/tui-layout-2/internal/ui/styles.go`

**Implementation:**
```go
package ui

import "github.com/charmbracelet/lipgloss"

// Styles will be expanded in Phase 3
var (
    // Placeholder styles - Phase 3 will implement full styling
    titleStyle = lipgloss.NewStyle().Bold(true)
)
```

**Acceptance Criteria:**
- [ ] `internal/ui/styles.go` file exists
- [ ] Package is `ui`
- [ ] At least one lipgloss style variable is defined
- [ ] File compiles without errors

**Verification:** Run `go build ./internal/ui`

**Time Estimate:** 2 minutes

**Dependencies:** Task 1.1 (lipgloss installed)

---

#### Task 1.6: Create Main Entry Point

**Description:** Create the `cmd/kanban/main.go` file that initializes the Bubbletea program with the Model. This is the application entry point — it should be minimal, just creating and running the program.

**File:** `/home/chauthan/projects/practical-go/tui-layout-2/cmd/kanban/main.go`

**Implementation:**
```go
package main

import (
    "fmt"
    "os"
    tea "charm.land/bubbletea/v2"
    "tui-layout-2/internal/ui"
)

func main() {
    // Create initial model with mock data
    model := ui.Model{
        columns: []ui.Column{  // Note: Using ui.Column for now, will update in Phase 2
            {Title: "To Do", Cards: []ui.Card{}},
            {Title: "In Progress", Cards: []ui.Card{}},
            {Title: "Done", Cards: []ui.Card{}},
        },
        focusedCol:   0,
        focusedCard:  0,
    }

    // Create and run the program
    p := tea.NewProgram(model)
    if _, err := p.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
```

**Acceptance Criteria:**
- [ ] `cmd/kanban/main.go` file exists
- [ ] Package is `main`
- [ ] Imports bubbletea with `charm.land/bubbletea/v2` path
- [ ] Creates tea.NewProgram with Model instance
- [ ] Calls p.Run() and handles errors
- [ ] Application exits with status 1 on error
- [ ] File compiles without errors

**Verification:** Run `go build ./cmd/kanban`

**Time Estimate:** 5 minutes

**Dependencies:** Task 1.4 (Model exists), Task 1.1 (bubbletea installed)

---

#### Task 1.7: Create README Documentation

**Description:** Create a basic README.md explaining the project, its purpose (Elm Architecture demonstration), and how to run it.

**File:** `/home/chauthan/projects/practical-go/tui-layout-2/README.md`

**Implementation:**
```markdown
# Kanban Board TUI Demo

A terminal user interface (TUI) demonstration app built in Go using [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss) libraries.

## Purpose

This is an educational project demonstrating the **Elm Architecture pattern** (Model-Update-View) for building terminal UIs in Go. It does NOT implement real kanban logic, backend, or persistence — all content is static mock data.

## Prerequisites

- Go 1.21 or later

## Running

\`\`\`bash
# Run the application
go run ./cmd/kanban

# Or build and run
go build -o kanban ./cmd/kanban
./kanban
\`\`\`

## Controls

- `q` or `ctrl+c`: Quit the application

## Project Structure

This project follows the [standard Go project layout](https://github.com/golang-standards/project-layout):

- `cmd/kanban/`: Application entry point
- `internal/ui/`: Bubbletea UI layer (Model, Update, View)
- `internal/domain/`: Business logic types

## Phase Status

- Phase 1: Project Scaffold ✅
- Phase 2: Model & State Management (Pending)
- Phase 3: Visual Styles (Pending)
- Phase 4: View & Layout Rendering (Pending)
- Phase 5: Polish & Responsive Layout (Pending)

## License

MIT
```

**Acceptance Criteria:**
- [ ] `README.md` file exists in project root
- [ ] Explains project purpose (Elm Architecture demo)
- [ ] Includes running instructions (`go run ./cmd/kanban`)
- [ ] Lists controls ('q' to quit)
- [ ] Documents project structure
- [ ] Shows phase status

**Verification:** Read the file and check all sections are present

**Time Estimate:** 5 minutes

**Dependencies:** None (can run in parallel with other tasks)

---

#### Task 1.8: Verify End-to-End Execution

**Description:** Run the complete application to verify it compiles, starts, renders text, and quits cleanly on 'q' key. This is the integration test for Phase 1.

**Steps:**
1. Run `go mod tidy` to ensure dependencies are correct
2. Run `go build ./cmd/kanban` to verify compilation
3. Run `go run ./cmd/kanban` to start the application
4. Verify "Initializing..." or "Press 'q' to quit." text appears
5. Press 'q' key and verify application exits cleanly
6. Run `go run ./cmd/kanban` again and press ctrl+c to verify alternative quit mechanism

**Acceptance Criteria:**
- [ ] `go mod tidy` runs without errors
- [ ] `go build ./cmd/kanban` produces executable without errors
- [ ] `go run ./cmd/kanban` starts the TUI application
- [ ] Application displays placeholder text (not blank screen)
- [ ] 'q' key causes clean exit (returns to shell prompt)
- [ ] ctrl+c key causes clean exit
- [ ] No error messages or panic output
- [ ] Terminal returns to normal state after exit

**Verification:** Manual execution test following the steps above

**Time Estimate:** 5 minutes

**Dependencies:** ALL previous tasks (1.1 through 1.7) must complete

---

## Execution Order Summary

**Wave 1 (Parallel - No Dependencies):**
- Task 1.1: Initialize Go module and dependencies (5 min)
- Task 1.2: Create directory structure (2 min)
- Task 1.7: Create README (5 min)

**Wave 2 (Sequential - Depends on Wave 1):**
- Task 1.3: Create domain types (3 min) — depends on 1.2
- Task 1.4: Create Model (5 min) — depends on 1.1, 1.3
- Task 1.5: Create styles stub (2 min) — depends on 1.1
- Task 1.6: Create main.go (5 min) — depends on 1.1, 1.4
- Task 1.8: End-to-end verification (5 min) — depends on ALL

**Total Time Estimate:** ~32 minutes (10 minutes parallel + 22 minutes sequential)

---

## Verification Criteria

Phase 1 is complete when ALL of the following are true:

### Compilation Checks
- [ ] `go mod tidy` runs without errors
- [ ] `go build ./cmd/kanban` succeeds (produces executable)
- [ ] `go build ./internal/ui` succeeds
- [ ] `go build ./internal/domain` succeeds
- [ ] No compilation warnings or errors

### Runtime Checks
- [ ] `go run ./cmd/kanban` starts the TUI application
- [ ] Application displays text (not blank/frozen)
- [ ] 'q' key causes clean exit
- [ ] ctrl+c causes clean exit
- [ ] No error messages on stdout or stderr
- [ ] Terminal state is normal after exit (no raw mode issues)

### File Structure Checks
- [ ] `go.mod` exists with module name `tui-layout-2`
- [ ] `go.sum` exists with dependency checksums
- [ ] `cmd/kanban/main.go` exists and compiles
- [ ] `internal/ui/model.go` exists and compiles
- [ ] `internal/ui/update.go` logic exists (in model.go for Phase 1)
- [ ] `internal/ui/view.go` logic exists (in model.go for Phase 1)
- [ ] `internal/ui/styles.go` exists and compiles
- [ ] `internal/domain/types.go` exists and compiles
- [ ] `README.md` exists with project documentation

### Dependency Checks
- [ ] `go.mod` contains `github.com/charmbracelet/bubbletea/v2 v2.0.0`
- [ ] `go.mod` contains `github.com/charmbracelet/lipgloss v2.0.0`
- [ ] Import paths use `charm.land/bubbletea/v2` (NOT old v1 path)
- [ ] `go list -m all` shows both dependencies

### Architecture Checks
- [ ] Production-grade structure used (cmd/, internal/)
- [ ] NOT a flat structure (all files not in root)
- [ ] Domain types separated from UI layer (internal/domain/)
- [ ] Elm Architecture pattern visible (Model, Init, Update, View methods)

---

## Risks and Mitigations

### Risk 1: Wrong Bubbletea Import Path (HIGH IMPACT)

**Problem:** Using `github.com/charmbracelet/bubbletea` instead of `charm.land/bubbletea/v2` imports v1 API, causing compilation failures.

**Mitigation:**
- RESEARCH.md explicitly documents this as Pitfall 1
- Verification criteria check import paths
- Use `grep -r "github.com/charmbracelet/bubbletea"` to catch old imports

**Detection:** Compilation errors about `tea.View` type or Init() signature mismatch

---

### Risk 2: Missing go.mod tidy (MEDIUM IMPACT)

**Problem:** Dependencies added but not fetched, causing "module not found" errors.

**Mitigation:**
- Task 1.1 explicitly includes `go mod tidy` step
- Task 1.8 verification includes `go mod tidy` check
- Always run after adding imports

**Detection:** `go run` fails with "missing module path" errors

---

### Risk 3: Flat Structure Instead of Production-Grade (MEDIUM IMPACT)

**Problem:** Creating all files in root directory instead of cmd/internal structure, violating production requirements.

**Mitigation:**
- Planning context explicitly requires production-grade structure
- Task 1.2 creates directory structure upfront
- Verification criteria check for cmd/ and internal/ directories

**Detection:** File structure missing `cmd/` or `internal/` directories

---

### Risk 4: Package Name Mismatches (LOW IMPACT)

**Problem:** Wrong package names in files (e.g., `package main` in internal/ui/model.go).

**Mitigation:**
- Task specifications explicitly state package names
- Verification includes `go build ./internal/...` checks

**Detection:** `go build` fails with "cannot find package" errors

---

## Notes for Executor

1. **Import Path Critical:** Always use `charm.land/bubbletea/v2` for bubbletea imports. The v1 path (`github.com/charmbracelet/bubbletea`) will NOT work.

2. **Production Structure:** Do NOT create a flat structure with all files in the root. Follow the `cmd/kanban/` and `internal/` pattern specified in Task 1.2.

3. **Elm Architecture:** Even though this is a minimal phase, maintain the separation of concerns. Don't put rendering logic in Update() or state logic in View().

4. **Go Module Path:** Use `tui-layout-2` as the module name (not `github.com/...`). This is local development.

5. **Manual Testing Required:** Task 1.8 requires actually running the TUI and pressing 'q'. Don't skip this — it confirms the app works end-to-end.

6. **Next Phase Preparation:** Phase 2 will expand the Model significantly (real columns/cards, keyboard navigation). Keep the structure clean to make Phase 2 easier.

---

## Phase Completion Checklist

Use this checklist before marking Phase 1 complete:

- [ ] All 8 tasks (1.1 through 1.8) are marked as done
- [ ] All compilation verification criteria pass
- [ ] All runtime verification criteria pass
- [ ] All file structure verification criteria pass
- [ ] All dependency verification criteria pass
- [ ] All architecture verification criteria pass
- [ ] README.md is updated with Phase 1 marked as complete
- [ ] No errors in `go run ./cmd/kanban`
- [ ] Application successfully quits on 'q' key
- [ ] Production-grade directory structure confirmed (cmd/, internal/)

**When ALL items are checked, Phase 1 is complete and ready for Phase 2.**

---

*Plan created: 2026-02-28*
*Planner: GSD Planner Agent*
*Phase: 01 - Project Scaffold*
*Requirements: PROJ-01, PROJ-02, PROJ-03, PROJ-04*
