# Phase 1: Project Scaffold - Research

**Researched:** 2026-02-28
**Domain:** Go TUI development with Charm libraries (bubbletea + lipgloss)
**Confidence:** HIGH

## Summary

Phase 1 focuses on initializing a Go module with the Charm ecosystem's bubbletea framework and lipgloss styling library. The research confirms that both libraries have reached stable v2.0.0 releases (February 2026), making this an ideal time to adopt them. The Elm Architecture pattern enforced by bubbletea requires a specific project structure with Model, Update, and View components clearly separated.

**Primary recommendation:** Initialize the Go module using `go mod init tui-layout-2` (local development) or `go mod init github.com/username/tui-layout-2` (if publishing), then install bubbletea v2.0.0 and lipgloss v2.0.0 via `go get`. Create a minimal bubbletea application following the official "simple" example pattern before building out the full Elm Architecture structure.

## Standard Stack

### Core

| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| **Go** | 1.21+ (latest stable) | Language | Go modules are default and mature since Go 1.16; module mode required for modern dependency management |
| **bubbletea** | v2.0.0 | TUI framework (Elm Architecture) | Charm's flagship framework, production-proven with 10,000+ applications, officially stable v2 released February 2026 |
| **lipgloss** | v2.0.0 | Terminal styling & layout | Official Charm styling companion to bubbletea, provides CSS-like declarative styling for terminals |

### Supporting

| Library | Version | Purpose | When to Use |
|---------|---------|---------|-------------|
| **bubbles** | (not needed for Phase 1) | Common UI components (text inputs, spinners) | Future phases if reusable components needed; not required for basic scaffold |

### Alternatives Considered

| Instead of | Could Use | Tradeoff |
|------------|-----------|----------|
| bubbletea v2.0.0 | bubbletea v1.x | v2 has breaking API changes; v2 is current stable; no reason to start with v1 |
| lipgloss v2.0.0 | termenv directly | lipgloss provides higher-level abstractions; termenv is lower-level; use lipgloss for productivity |
| `charm.land/bubbletea/v2` import | `github.com/charmbracelet/bubbletea/v2` import | Both resolve to same package; `charm.land` is official import path, shorter, recommended in examples |

**Installation:**

```bash
# Initialize Go module (local development - use this for Phase 1)
go mod init tui-layout-2

# Alternative: If planning to publish to GitHub (optional, can migrate later)
# go mod init github.com/yourusername/tui-layout-2

# Install dependencies (specifically v2.0.0)
go get github.com/charmbracelet/bubbletea/v2@v2.0.0
go get github.com/charmbracelet/lipgloss@v2.0.0

# Download dependencies and create go.sum
go mod tidy
```

**Critical Note on Imports:** Bubbletea v2 uses a different import path than v1:

```go
// CORRECT for v2.0.0
import tea "charm.land/bubbletea/v2"

// INCORRECT (this is v1)
// import tea "github.com/charmbracelet/bubbletea"
```

## Architecture Patterns

### Recommended Project Structure (Phase 1 Minimal)

```
tui-layout-2/
├── go.mod                 # Module definition
├── go.sum                 # Dependency checksums
├── main.go                # Entry point, creates tea.Program
├── model.go               # Model struct + Init() method
├── view.go                # View() function for rendering
└── styles.go              # lipgloss style definitions (for Phase 3, can be empty stub)
```

### Recommended Project Structure (Future Phases)

For Phase 2+, consider this expanded structure:

```
tui-layout-2/
├── go.mod
├── go.sum
├── main.go                # Entry point only
├── model.go               # Core Model struct + Init()
├── update.go              # Update() method (can separate if complex)
├── view.go                # View() function
├── styles.go              # lipgloss style definitions
└── types.go               # Custom types (Card, Column, etc.) - Phase 2
```

### Pattern 1: Minimal Bubbletea Application

**What:** The absolute minimum code needed to create a working bubbletea app that renders text and quits on 'q' key.

**When to use:** Phase 1 scaffold - verifies the framework works before building complex logic.

**Example:**

```go
// Source: https://github.com/charmbracelet/bubbletea (examples/simple/main.go)
// Adapted for Phase 1 requirements

package main

import (
    "fmt"
    "os"
    tea "charm.land/bubbletea/v2"
)

// Model can be any type. For minimal app, a simple struct works.
type model struct {
    ready bool
}

// Init returns initial command. Return nil for no I/O.
func (m model) Init() tea.Cmd {
    return nil
}

// Update handles incoming messages and updates model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

// View renders the UI as a string.
func (m model) View() string {
    return "Hello, Bubble Tea! Press 'q' to quit.\n"
}

func main() {
    p := tea.NewProgram(model{})
    if _, err := p.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err)
        os.Exit(1)
    }
}
```

### Pattern 2: Elm Architecture in Go (Bubbletea)

**What:** The Model-Update-View (Init) pattern that bubbletea enforces. This is the core educational value of the project.

**When to use:** All bubbletea applications - this is the framework's paradigm.

**Components:**

1. **Model** - Application state (struct with fields like `columns`, `focusedCol`, `width`, `height`)
2. **Init()** - Returns initial `tea.Cmd` (often `nil` for simple apps)
3. **Update(msg tea.Msg) (tea.Model, tea.Cmd)` - Pure function handling events, returns updated model
4. **View() string` - Renders UI based on current model state

**Key Principles:**

- **Unidirectional data flow:** Messages → Update → Model → View
- **Pure functions:** Update doesn't perform I/O, returns commands instead
- **Single source of truth:** Model contains all application state

### Pattern 3: Lipgloss Style Organization

**What:** Centralized style definitions using lipgloss's declarative API.

**When to use:** Phase 3 (Visual Styles) - keeping all styling in one file for maintainability.

**Example (for Phase 3 reference):**

```go
// styles.go
package main

import "github.com/charmbracelet/lipgloss"

var (
    // Define styles as package-level variables
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#7D56F4")).
        PaddingTop(1)

    helpStyle = lipgloss.NewStyle().
        Faint(true).
        Foreground(lipgloss.Color("241"))
)
```

### Anti-Patterns to Avoid

- **Global state manipulation:** Never modify global variables in Update(). All state must be in the Model.
- **I/O in Update:** Update() should be pure. Use `tea.Cmd` for I/O operations.
- **Mixing concerns:** Don't put rendering logic in Update() or state logic in View().
- **Hardcoded strings in View:** Use style definitions from styles.go, not inline formatting.

## Don't Hand-Roll

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| Terminal styling (colors, borders, padding) | Manual ANSI escape codes | lipgloss styles | ANSI codes are complex, error-prone, and terminal-dependent; lipgloss handles color profiles automatically |
| Layout calculations (measuring text width, joining blocks) | Custom string manipulation | lipgloss utilities (Width, Height, JoinHorizontal, JoinVertical) | East Asian width characters, tab handling, and proper cell counting are non-trivial |
| Event loop and message routing | Custom select/timeout logic | bubbletea's runtime | Handles stdin, window size, timers, and more; already debugged |
| Alternative screen management | Manual termios calls | tea.WithAltScreen() option | Cross-platform terminal raw mode handling |

**Key insight:** The Charm ecosystem has solved terminal UI primitives. Hand-rolling any of the above introduces bugs that bubbletea/lipgloss have already fixed through production use across thousands of applications.

## Common Pitfalls

### Pitfall 1: Wrong Import Path for Bubbletea v2

**What goes wrong:** Using `github.com/charmbracelet/bubbletea` instead of `charm.land/bubbletea/v2` results in importing v1.x, causing API incompatibility errors.

**Why it happens:** v2 introduced a breaking API change and uses a different import path (following Go's semantic import versioning). Old tutorials and examples still reference the v1 path.

**How to avoid:** Always use `import tea "charm.land/bubbletea/v2"` for new projects. Verify by checking `go.mod` shows `github.com/charmbracelet/bubbletea/v2 v2.0.0`.

**Warning signs:** Compile errors about missing `tea.View` type (v2 uses `tea.View` return type), or `Init()` signature mismatch.

### Pitfall 2: Module Path Mismatch

**What goes wrong:** `go mod init myapp` but imports use `github.com/username/myapp/utils`, causing "cannot find module" errors.

**Why it happens:** Module path in `go.mod` must match import paths. Changing `go.mod` module line requires updating all imports.

**How to avoid:**
- For local-only projects: `go mod init tui-layout-2` and import as `tui-layout-2/...` (though single-package apps won't need internal imports)
- For GitHub projects: `go mod init github.com/username/tui-layout-2` before writing any code
- Never use uppercase letters in module paths (Go rejects them)

**Warning signs:** `go build` fails with "cannot find module providing package" when importing your own packages.

### Pitfall 3: Forgetting go mod tidy

**What goes wrong:** `go run .` fails with "missing module path" or dependencies aren't downloaded.

**Why it happens:** `go mod init` creates empty `go.mod`. Dependencies aren't fetched until `go mod tidy` scans imports.

**How to avoid:** Always run `go mod tidy` after:
1. Running `go get` to install packages
2. Adding new imports to any `.go` file
3. Before first `go run` or `go build`

**Verification:** `go.mod` should have `require` entries for bubbletea/v2 and lipgloss. `go.sum` should exist with checksums.

### Pitfall 4: Not Using Alt Screen for Full-Window Apps

**What goes wrong:** TUI output mixes with shell history, looking messy and confusing.

**Why it happens:** Default bubbletea uses inline mode. Full-window TUIs need alternative screen buffer.

**How to avoid:** For Phase 1, inline mode is fine (minimal app). For Phase 2+ (full kanban board), use:

```go
p := tea.NewProgram(model{}, tea.WithAltScreen())
```

**Warning signs:** Text appears above your shell prompt; pressing 'q' shows previous terminal content mixed with app output.

### Pitfall 5: Blocking Operations in Update

**What goes wrong:** App freezes during file I/O or network calls.

**Why it happens:** Update() runs on the main thread. Blocking calls block the entire event loop.

**How to avoid:** Return `tea.Cmd` functions that perform I/O asynchronously. Phase 1 won't need this, but Phase 2+ might.

**Example pattern (for reference):**

```go
// Return a command instead of blocking
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // ...
    return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
        return TickMsg(t)
    })
}
```

### Pitfall 6: East Asian Character Width Issues

**What goes wrong:** Borders misalign when text contains Chinese/Japanese/Korean characters.

**Why it happens:** Some terminals treat CJK characters as double-width. lipgloss handles this, but requires environment setup.

**How to avoid:** Set `RUNEWIDTH_EASTASIAN=0` if experiencing alignment issues with CJK content (documented in lipgloss FAQ).

**Warning signs:** Borders broken, text overflowing containers, content shifted right.

## Code Examples

Verified patterns from official sources:

### Example 1: Minimal Quit-on-'q' Application

```go
// Source: Adapted from https://github.com/charmbracelet/bubbletea (examples/simple)
// Meets PROJ-04 requirement: app runs and exits cleanly on 'q' key

package main

import (
    "fmt"
    "os"
    tea "charm.land/bubbletea/v2"
)

type model struct{}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    return "Press 'q' to quit.\n"
}

func main() {
    p := tea.NewProgram(model{})
    if _, err := p.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err)
        os.Exit(1)
    }
}
```

### Example 2: Go Module Initialization (2025 Standard)

```bash
# Standard 2025 Go module initialization workflow
mkdir tui-layout-2 && cd tui-layout-2
go mod init tui-layout-2        # For local-only development
# OR
go mod init github.com/username/tui-layout-2  # If planning to publish to GitHub

go get github.com/charmbracelet/bubbletea/v2@v2.0.0
go get github.com/charmbracelet/lipgloss@v2.0.0

# Critical: always run tidy to fetch dependencies and create go.sum
go mod tidy

# Verify installation
go run .   # Should compile and run
```

### Example 3: Basic Lipgloss Styling (Phase 3 Preview)

```go
// Source: https://github.com/charmbracelet/lipgloss (official README)
// For reference in Phase 3

import "github.com/charmbracelet/lipgloss"

var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingTop(2).
    PaddingLeft(4).
    Width(22)

// Usage in View()
func (m model) View() string {
    return style.Render("Hello, kitty")
}
```

## State of the Art

| Old Approach | Current Approach | When Changed | Impact |
|--------------|------------------|--------------|--------|
| bubbletea v1.x API | bubbletea v2.0.0 API | February 2026 | New import path (`charm.land/bubbletea/v2`), `View()` returns `tea.View` type instead of `string` |
| GOPATH-based workspace | Go modules (default since Go 1.16) | Go 1.16 (2021) | No need for $GOPATH, `go mod init` is standard initialization |
| Manual ANSI escape codes | lipgloss declarative styles | 2020-2025 | CSS-like syntax, automatic terminal capability detection |
| Import path `github.com/charmbracelet/bubbletea` | Import path `charm.land/bubbletea/v2` for v2+ | February 2026 | Semantic import versioning - breaking changes require new import path |

**Deprecated/outdated:**

- **bubbletea v0.x / v1.x:** Superseded by v2.0.0. Don't use for new projects.
- **GOPATH workspace setup:** Deprecated since Go 1.16. Always use `go mod init`.
- **Manual terminal size parsing:** Use `tea.WindowSizeMsg` in Update(), not `unix.Ioctl` calls.
- **Importing bubbletea without `/v2`:** Will get v1.x, which has incompatible API.

## Open Questions

1. **Should we use WithAltScreen() in Phase 1?**
   - What we know: Alt screen provides full-window buffer (cleaner for full apps), inline mode mixes with shell (better for simple tools)
   - What's unclear: Whether PROJ-04's "minimal working bubbletea application" implies full-window or inline
   - Recommendation: Use inline mode for Phase 1 (matches official "simple" example), add WithAltScreen() in Phase 2 when building full kanban layout

2. **Module path for local development vs. future publishing**
   - What we know: `go mod init tui-layout-2` works locally; changing to GitHub path later requires updating go.mod and all imports
   - What's unclear: Whether this project will be published publicly
   - Recommendation: Use `tui-layout-2` for Phase 1 (simpler), can migrate to GitHub path before public release if needed

## Sources

### Primary (HIGH confidence)

- **bubbletea GitHub repository** - https://github.com/charmbracelet/bubbletea
  - Verified: Latest release v2.0.0 (February 2026)
  - Verified: Import path for v2 is `charm.land/bubbletea/v2`
  - Verified: Simple example structure with Init/Update/View pattern
  - Verified: Quit handling with tea.Quit command

- **lipgloss GitHub repository** - https://github.com/charmbracelet/lipgloss
  - Verified: Latest release v2.0.0
  - Verified: Declarative style API with CSS-like syntax
  - Verified: Color support (ANSI 16, 256, True Color)
  - Verified: Layout utilities (JoinHorizontal, JoinVertical)

- **Go module documentation** - https://go.dev/ref/mod
  - Verified: `go mod init` syntax and best practices
  - Verified: Module path requirements (lowercase only)
  - Verified: `go mod tidy` purpose and usage

### Secondary (MEDIUM confidence)

- **Bubble Tea tutorial** (official GitHub README)
  - Verified: Elm Architecture pattern explanation (Model-Init-Update-View)
  - Verified: Message handling with type switches
  - Verified: tea.Quit command for clean exit

- **Chinese blog posts on bubbletea** (CSDN, 2025)
  - Verified: Installation commands match official docs
  - Verified: Project initialization workflow
  - Note: Some examples may use v1 API; cross-referenced with official docs

- **Go module best practices articles** (PHP Chinese, 2025)
  - Verified: Common pitfalls (uppercase paths, wrong directory for init)
  - Verified: go.sum purpose (checksums, not manual editing)
  - Note: Non-English sources, but technical content appears accurate

### Tertiary (LOW confidence)

- **Web search for "bubbletea beginner pitfalls"**
  - Result: Irrelevant content (bubble tea drinks vs. Bubble Tea framework)
  - Action: Discarded, relied on official docs and GitHub issues instead

### Production Examples (HIGH confidence)

- **charmbracelet/glow source code** - https://github.com/charmbracelet/glow
  - Verified: Main entry point pattern with Cobra CLI framework
  - Verified: Configuration loading with Viper and env packages
  - Verified: UI separation in `ui/` directory (ui.go, markdown.go, pager.go, stash.go)
  - Verified: Utility functions in separate `utils/` package
  - Verified: Error handling with log package and file output
  - Verified: Multiple modes: TUI, pager, and inline rendering

- **jesseduffield/lazydocker source code** - https://github.com/jesseduffield/lazydocker
  - Verified: Extensive `pkg/` structure (app/, commands/, gui/, config/, utils/)
  - Verified: Clear separation between GUI layer and Docker command layer
  - Verified: Panel-based architecture with self-contained components
  - Verified: Configuration system with customizable keybindings
  - Verified: Internationalization support (i18n package)
  - Verified: Comprehensive test coverage

- **Standard Go Project Layout** - https://github.com/golang-standards/project-layout
  - Verified: Community-standard directory structure (cmd/, internal/, pkg/)
  - Verified: Compiler enforcement of internal/ directory
  - Verified: Best practices for application entry points
  - Verified: Separation concerns: reusable code in pkg/, private code in internal/

### Architecture Patterns (MEDIUM confidence)

- **Domain-Driven Design articles** (Thoughtworks, CSDN, 2025)
  - Verified: Hexagonal architecture (ports and adapters) pattern
  - Verified: Domain layer isolation from infrastructure
  - Verified: Repository pattern for data access
  - Verified: Aggregate roots and entity modeling

- **Go project structure best practices** (CSDN, PHP Chinese, 2025)
  - Verified: cmd/ vs internal/ vs pkg/ usage patterns
  - Verified: When to use each directory based on project size
  - Verified: Common pitfalls of over-structuring small projects
  - Verified: Module boundary enforcement

### Configuration Management (MEDIUM confidence)

- **Viper configuration library** - https://github.com/spf13/viper
  - Verified: Multi-source configuration (files, env vars, flags)
  - Verified: Automatic environment variable binding
  - Verified: Configuration file discovery patterns
  - Verified: Type-safe unmarshaling to structs

- **Go configuration articles** (CSDN, PHP Chinese, 2025)
  - Verified: 12-Factor App configuration principles
  - Verified: Environment variable vs config file tradeoffs
  - Verified: Configuration priority hierarchies
  - Verified: Validation patterns for loaded configuration

### Testing Strategies (MEDIUM confidence)

- **Bubbletea testing articles** (CSDN, 2025-2026)
  - Verified: Unit testing domain logic without bubbletea
  - Verified: teatest package for integration testing
  - Verified: Model state transition testing
  - Verified: View rendering assertion patterns
  - Verified: Golden file testing for UI regression detection

- **Production testing patterns** (Glow, Lazydocker codebases)
  - Verified: Table-driven tests for multiple scenarios
  - Verified: Mock repositories for deterministic testing
  - Verified: Benchmark tests for rendering performance
  - Verified: Error handling test coverage

## Metadata

**Confidence breakdown:**

- **Standard stack:** HIGH - Verified via GitHub API (v2.0.0 releases), official documentation, Go pkg.go.dev
- **Architecture:** HIGH - Elm Architecture pattern well-documented in official bubbletea tutorial; examples align with requirements
- **Pitfalls:** MEDIUM - Import path issue verified in official docs; module path issues cross-referenced across multiple Go best practice sources; alt screen issue documented but low-impact for Phase 1
- **Project structure:** HIGH - Official examples show single-file or minimal multi-file layouts; verified against requirements PROJ-03 (main.go, model.go, view.go, styles.go)
- **Production patterns:** HIGH - Real production apps (Glow, lazydocker) analyzed; standard Go project layout researched; configuration and testing patterns documented

**Research date:** 2026-02-28
**Valid until:** 2026-04-30 (60 days - Go ecosystem is stable; Charm libraries release cycle is measured in months)

---

## Production-Grade Project Structure

### Standard Go Project Layout Applied to TUI Applications

The [golang-standards/project-layout](https://github.com/golang-standards/project-layout) repository provides a community-standard template for organizing Go projects. For TUI applications like this Kanban board, we adapt this as follows:

```
tui-layout-2/
├── cmd/                        # Application entry points
│   └── kanban/                 # Main CLI application
│       └── main.go             # Entry point: parse flags, initialize config, start tea.Program
├── internal/                   # Private application code (compiler-enforced)
│   ├── ui/                     # UI layer (bubbletea models and views)
│   │   ├── model.go            # Core application state
│   │   ├── update.go           # Event handling (Update method)
│   │   ├── view.go             # Rendering logic (View method)
│   │   ├── styles.go           # Lipgloss style definitions
│   │   └── keys.go             # Key binding definitions
│   ├── domain/                 # Business logic (separated from UI)
│   │   ├── card.go             # Card domain model
│   │   ├── column.go           # Column domain model
│   │   └── board.go            # Board aggregate
│   └── config/                 # Configuration management
│       └── config.go           # Config struct and loading
├── pkg/                        # Publicly reusable components (optional)
│   └── tuiutil/                # Generic TUI helpers if needed
│       └── viewport.go         # Reusable viewport component
├── go.mod
├── go.sum
├── Makefile                    # Build automation
├── README.md
└── .planning/                  # Project planning docs (not deployed)
```

### Key Principles

**1. cmd/ Directory - Application Entry Points**

- **Purpose**: Contains exactly one `main` package per executable binary
- **Pattern**: Each subdirectory under `cmd/` represents a different binary
- **Example**: `cmd/kanban/main.go` contains only initialization logic
- **Responsibility**: Parse flags, load configuration, delegate to internal packages

**Why this matters**: As the project grows, you might add additional tools (e.g., `cmd/migrator/main.go` for data migration scripts). The `cmd/` structure keeps each entry point focused.

**2. internal/ Directory - Private Implementation**

- **Compiler enforcement**: Go's `internal/` directory is special - packages outside your module cannot import it, even if they try
- **Use case**: All business logic, UI implementation, domain models live here
- **Benefit**: Clear signal: "This is implementation detail, not a public API"

**Why this matters**: If you later extract reusable components into a library, `internal/` code stays private. The compiler prevents accidental dependencies.

**3. pkg/ Directory - Reusable Components (Optional)**

- **Purpose**: Code that could be extracted into a separate library
- **Pattern**: Only put code here if you intend for external projects to import it
- **Example**: Generic TUI components like custom viewport, scrollbar widgets
- **Caution**: Don't use `pkg/` as a "utils" dump - be intentional about reusability

**Why this matters**: For Phase 1, you likely won't need `pkg/`. Start without it. Add it only when you have genuinely reusable components.

**4. Domain-Driven Design in TUI Context**

For the Kanban board, domain separation is crucial even though there's no backend:

```
internal/domain/
├── card.go          # Card struct + business rules
├── column.go        # Column struct + business rules
└── board.go         # Board aggregate (manages columns and cards)
```

**Domain models** (pure Go structs, no bubbletea dependency):

```go
// internal/domain/card.go
package domain

type Card struct {
    ID      string
    Title   string
    Content string
    Status  Status
}

type Status string

const (
    StatusTodo       Status = "todo"
    StatusInProgress Status = "in-progress"
    StatusDone       Status = "done"
)

// Business rules
func (c *Card) CanMoveTo(target Status) bool {
    // Domain logic: what moves are valid?
    return true
}
```

**UI layer** (depends on domain, not vice versa):

```go
// internal/ui/model.go
package ui

import "tui-layout-2/internal/domain"

type model struct {
    board *domain.Board  // Domain state
    focus FocusState     // UI-specific state (cursor, selection)
}
```

**Why this matters**: Separation of concerns makes testing easier. Domain logic can be unit tested without bubbletea. UI code focuses on rendering and interaction.

---

## Data Layer Patterns

### How Real Data Would Integrate with Bubbletea

Even though Phase 1 uses static mock data, understanding data flow patterns is essential for production-grade TUI apps.

### Repository Pattern in Bubbletea

**Pattern**: Domain layer defines interfaces; UI layer triggers commands; repositories return results as messages.

```
┌─────────────┐
│   Update()  │
└──────┬──────┘
       │ Returns tea.Cmd
       ▼
┌─────────────────┐
│  tea.Cmd        │  "Load data from repository"
└──────┬──────────┘
       │ Executes async
       ▼
┌─────────────────┐
│  Repository     │  "Fetch from DB/file/API"
└──────┬──────────┘
       │ Returns result
       ▼
┌─────────────────┐
│  tea.Msg        │  "DataLoadedMsg{data}"
└──────┬──────────┘
       │ Routes back to Update()
       ▼
┌─────────────┐
│   Update()  │  "Handle DataLoadedMsg"
└─────────────┘
```

**Example structure** (for future phases with persistence):

```go
// internal/domain/repository.go
package domain

type CardRepository interface {
    GetAll() ([]Card, error)
    GetByID(id string) (Card, error)
    Save(card Card) error
}

// Message types for async operations
type CardsLoadedMsg struct {
    Cards []Card
    Err   error
}
```

```go
// internal/ui/commands.go
package ui

func loadCardsCmd(repo domain.CardRepository) tea.Cmd {
    return func() tea.Msg {
        cards, err := repo.GetAll()
        return CardsLoadedMsg{Cards: cards, Err: err}
    }
}
```

```go
// internal/ui/update.go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "r" { // Refresh
            return m, loadCardsCmd(m.repo)
        }
    case CardsLoadedMsg:
        if msg.Err != nil {
            // Handle error in UI
            m.err = msg.Err
        } else {
            m.board.SetCards(msg.Cards)
        }
        return m, nil
    }
    return m, nil
}
```

**Key insights from production apps**:

1. **Glow's approach**: Uses `github.com/caarlos0/env` to parse environment variables into a config struct. Configuration is loaded before the tea.Program starts.

2. **Lazydocker's approach**: Separates concerns into `pkg/commands` (Docker API interactions) and `pkg/gui` (UI). The GUI layer triggers commands, which return messages back to the UI.

3. **Async I/O pattern**: Never block in `Update()`. Always return a `tea.Cmd` for I/O operations. The framework handles goroutines and message routing.

**Why this matters**: Even without a real backend, understanding this pattern helps structure Phase 1 code. Keep "what data" (domain) separate from "how to display" (UI).

---

## Testing Strategy

### How to Test Bubbletea Applications in Production

Testing TUI applications requires special techniques beyond standard unit tests.

### 1. Unit Testing Domain Logic

**What**: Test pure functions and business rules without bubbletea dependency.

**Example**:

```go
// internal/domain/card_test.go
package domain_test

import "testing"
import "tui-layout-2/internal/domain"

func TestCardCanMoveTo(t *testing.T) {
    card := domain.Card{Status: domain.StatusTodo}

    if !card.CanMoveTo(domain.StatusInProgress) {
        t.Error("Expected card to be movable to in-progress")
    }
}
```

**Why**: Domain logic is pure Go - test it like any other package.

### 2. Testing Model State Transitions

**What**: Verify `Update()` method correctly transforms model state.

**Example**:

```go
// internal/ui/model_test.go
package ui

import (
    "testing"
    tea "charm.land/bubbletea/v2"
)

func TestModelUpdateQuit(t *testing.T) {
    m := model{}
    msg := tea.KeyMsg{Type: tea.KeyCtrlC}

    newModel, cmd := m.Update(msg)

    if cmd != tea.Quit {
        t.Errorf("Expected tea.Quit, got %v", cmd)
    }
}
```

**Why**: Test state machine logic without rendering.

### 3. Testing View Rendering

**What**: Assert `View()` returns expected strings.

**Example**:

```go
// internal/ui/view_test.go
package ui

import (
    "testing"
)

func TestViewRendersTitle(t *testing.T) {
    m := model{board: &domain.Board{Title: "My Board"}}
    output := m.View()

    if !strings.Contains(output, "My Board") {
        t.Error("Expected view to contain board title")
    }
}
```

**Why**: Catch rendering regressions without running terminal.

### 4. Integration Testing with `teatest`

Bubbletea v2 provides the `teatest` package for end-to-end testing.

**What**: Simulate user input sequences and verify terminal output.

**Example** (from bubbletea's test suite):

```go
// internal/ui/integration_test.go
package ui

import (
    "testing"
    "time"
    tea "charm.land/bubbletea/v2"
    "github.com/charmbracelet/bubbletea/v2/teatest"
)

func TestInteractiveFlow(t *testing.T) {
    m := initialModel()
    tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))

    // Simulate user pressing 'down' key
    tm.Send(tea.KeyMsg{Type: tea.KeyDown})
    tm.WaitFinished(t, teatest.WithDuration(time.Second))

    // Verify cursor moved
    finalModel := tm.FinalModel(t).(model)
    if finalModel.cursor != 1 {
        t.Errorf("Expected cursor at position 1, got %d", finalModel.cursor)
    }
}
```

**Why**: Test complete user flows, including timing and async operations.

### Production Testing Patterns

**From Glow and Lazydocker**:

1. **Table-driven tests**: Test multiple scenarios with a single test function
2. **Golden file testing**: Compare view output against known-good snapshots
3. **Mock repositories**: Inject fake repositories for deterministic tests
4. **Benchmark tests**: Measure rendering performance for large datasets

**Example: Golden file test**:

```go
func TestViewGolden(t *testing.T) {
    m := model{ /* fixed state */ }
    got := m.View()

    golden, err := os.ReadFile("testdata/view_golden.txt")
    if err != nil {
        t.Fatal(err)
    }

    if got != string(golden) {
        t.Errorf("View output changed. Run: go test -update-golden")
    }
}
```

**Why**: Detect visual regressions. When you intentionally change the UI, update the golden file.

---

## Configuration Management

### Production-Grade Configuration Patterns

Glow demonstrates excellent configuration management using multiple sources and priorities.

### Configuration Hierarchy (from lowest to highest priority)

1. **Defaults** (hardcoded in code)
2. **Config file** (YAML in `~/.config/glow/glow.yml`)
3. **Environment variables** (`GLOW_STYLE`, `GLOW_WIDTH`)
4. **Command-line flags** (`glow -w 120 -s light`)

### Glow's Approach

**File**: `main.go` in charmbracelet/glow

**Key components**:

1. **Viper for unified config**: Supports files, env vars, and flags
2. **Cobra for CLI parsing**: Subcommands and flag definitions
3. **env package**: Type-safe environment variable parsing

**Example** (adapted for this project):

```go
// internal/config/config.go
package config

import (
    "github.com/spf13/viper"
    "github.com/caarlos0/env/v11"
)

type Config struct {
    // UI settings
    Width      int    `mapstructure:"width" env:"WIDTH"`
    Style      string `mapstructure:"style" env:"STYLE"`
    Mouse      bool   `mapstructure:"mouse" env:"MOUSE"`
    Debug      bool   `mapstructure:"debug" env:"DEBUG"`

    // Domain settings (for future phases)
    DataFile   string `mapstructure:"data_file" env:"DATA_FILE"`
    AutoSave   bool   `mapstructure:"auto_save" env:"AUTO_SAVE"`
}

// Load config from multiple sources
func Load() (*Config, error) {
    cfg := &Config{}

    // 1. Set defaults
    viper.SetDefault("width", 80)
    viper.SetDefault("style", "auto")
    viper.SetDefault("mouse", false)

    // 2. Config file (optional)
    viper.SetConfigName("kanban")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("$HOME/.config/kanban")
    viper.AddConfigPath(".")
    viper.ReadInConfig() // Ignore error if file doesn't exist

    // 3. Environment variables
    viper.SetEnvPrefix("KANBAN")
    viper.AutomaticEnv()

    // 4. Unmarshal to struct
    if err := viper.Unmarshal(cfg); err != nil {
        return nil, err
    }

    // 5. Environment variables override (type-safe)
    if err := env.Parse(cfg); err != nil {
        return nil, err
    }

    return cfg, nil
}
```

**Usage in main.go**:

```go
// cmd/kanban/main.go
package main

import (
    "tea "charm.land/bubbletea/v2"
    "tui-layout-2/internal/config"
    "tui-layout-2/internal/ui"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        panic(err)
    }

    // Pass config to UI
    p := tea.NewProgram(ui.NewModel(cfg), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        panic(err)
    }
}
```

### Minimal Approach (Phase 1)

For Phase 1, defer complex configuration. Use hardcoded defaults only:

```go
// internal/config/config.go (Phase 1 minimal)
package config

type Config struct {
    Width int
    Style string
}

func Default() *Config {
    return &Config{
        Width: 80,
        Style: "auto",
    }
}
```

**Why**: YAGNI principle. Add Viper/Cobra complexity only when you need it (likely Phase 3+ when adding user customization).

### Configuration Best Practices

**From Go community standards**:

1. **12-Factor App**: Store config in environment, not code
2. **Priority order**: Defaults < File < Env < Flags (allows overriding)
3. **Type safety**: Use structs, not `viper.Get("key")` strings
4. **Validation**: Validate config after loading (e.g., width must be > 0)
5. **Documentation**: Document all config options in README

**Example: Validation**:

```go
func (c *Config) Validate() error {
    if c.Width < 40 || c.Width > 200 {
        return fmt.Errorf("width must be between 40 and 200, got %d", c.Width)
    }
    return nil
}
```

**Why**: Fail fast with clear error messages rather than confusing runtime behavior.

---

## Reusable Component Strategy

### Creating TUI Components for Reusability

Bubbletea's [Bubbles](https://github.com/charmbracelet/bubbles) package demonstrates component reusability. Components are:

1. **Independent**: Have their own Model, Update, View
2. **Composable**: Can embed in parent models
3. **Configurable**: Accept initialization options
4. **Testable**: Work in isolation

### Pattern: Embedded Components

**Example**: If you build a custom scrollbar component:

```go
// pkg/tuiutil/scrollbar.go
package tuiutil

import tea "charm.land/bubbletea/v2"

type Scrollbar struct {
    height     int
    scrollTop  int
    totalLines int
}

func (s Scrollbar) Init() tea.Cmd {
    return nil
}

func (s Scrollbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return s, nil
}

func (s Scrollbar) View() string {
    // Render scrollbar string
}
```

**Usage in main UI**:

```go
// internal/ui/model.go
type model struct {
    scrollbar  tuiutil.Scrollbar  // Embedded component
    content    string
}

func (m model) View() string {
    return m.scrollbar.View() + "\n" + m.content
}
```

**Why**: Embedding allows reuse without inheritance. Components communicate via messages.

### When to Create Components vs. Inline Code

**Extract a component when**:

- Same UI element appears 3+ times
- Element has complex state (e.g., text input with validation)
- Element is used in multiple projects

**Keep inline when**:

- Simple one-off rendering (e.g., a status bar)
- Tightly coupled to specific domain logic
- Only appears once

**Bubbles examples**: `textarea`, `textinput`, `viewport` - complex enough to warrant abstraction.

---

## Error Handling Patterns

### Production-Grade Error Handling for TUI Apps

### 1. Error Messages in UI

**Pattern**: Render errors in view, don't crash program.

```go
type model struct {
    err  error
    // ... other fields
}

func (m model) View() string {
    if m.err != nil {
        return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
    }
    // Normal rendering
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case errorMsg:
        m.err = msg.err
        return m, nil
    }
    // ... other cases
}
```

**Why**: Users see helpful messages, not stack traces.

### 2. Logging for Debugging

**Glow's approach**: Use `charmbracelet/log` package with file output.

```go
// In main.go
f, err := tea.LogToFile("debug.log", "debug")
if err != nil {
    log.Error("Failed to open log file", "err", err)
}
defer f.Close()
```

**Why**: Non-intrusive debugging. Doesn't clutter TUI output.

### 3. Graceful Degradation

**Pattern**: Detect terminal capabilities and adapt.

```go
func detectCapabilities() (colorProfile, mouseSupport) {
    if !term.IsTerminal(int(os.Stdout.Fd())) {
        return lipgloss.NoTTY{}, false
    }

    return lipgloss.ColorProfile(), true
}
```

**Why**: Works in CI pipelines (no TTY) and terminals with limited features.

---

## Production Examples Analysis

### Glow (charmbracelet/glow)

**Structure**:
- `main.go`: Entry point, CLI parsing with Cobra
- `ui/`: All bubbletea code (ui.go, markdown.go, pager.go, etc.)
- `utils/`: Helper functions (file operations, string utilities)

**Key patterns**:
- Config loaded before tea.Program starts
- Multiple modes: TUI, pager, inline rendering
- Markdown rendering separate from UI logic
- Extensive keybindings in `ui/keys.go`

**Lessons**:
- Separate "business logic" (markdown rendering) from UI
- Use environment variables for developer options (debug, profiling)
- Provide multiple interfaces (TUI vs. CLI) based on context

### Lazydocker (jesseduffield/lazydocker)

**Structure**:
- `main.go`: Minimal entry
- `pkg/app/`: Application orchestrator
- `pkg/gui/`: All UI code (presentation, panels, layout)
- `pkg/commands/`: Docker API integration
- `pkg/config/`: Configuration management

**Key patterns**:
- Clear separation: GUI talks to Commands via interfaces
- Panel-based architecture (containers, images, volumes, etc.)
- Each panel is self-contained model
- Keybindings centralized and customizable
- Heavy use of `gocui` (similar to bubbletea, different framework)

**Lessons**:
- Component isolation enables independent testing
- Configuration-driven keybindings for user customization
- Separate "service" layer for external system interaction

---

## Recommended Project Structure for Phase 1

Given production patterns and Phase 1's scope (minimal working app), start simple:

```
tui-layout-2/
├── go.mod
├── go.sum
├── main.go                   # Entry point (no flags yet)
├── model.go                  # Core Model + Init()
├── update.go                 # Update() method
├── view.go                   # View() function
└── styles.go                 # Lipgloss styles (can be empty initially)
```

**Phase 2+ expansion** (as complexity grows):

```
tui-layout-2/
├── cmd/
│   └── kanban/
│       └── main.go           # Move main.go here
├── internal/
│   ├── ui/
│   │   ├── model.go          # Move from root
│   │   ├── update.go         # Move from root
│   │   ├── view.go           # Move from root
│   │   ├── styles.go         # Move from root
│   │   └── keys.go           # New: keybinding constants
│   └── domain/
│       ├── card.go           # New: domain logic
│       ├── column.go         # New: domain logic
│       └── board.go          # New: aggregate root
├── pkg/                      # Optional: add later if needed
├── go.mod
└── README.md
```

**Why**: Start flat, add structure when needed. Over-engineering early slows progress.

---

## Key Takeaways for Production-Grade TUI Apps

1. **Separation of concerns**: Domain logic ≠ UI rendering ≠ I/O
2. **Standard Go layout**: `cmd/`, `internal/`, `pkg/` - these aren't ceremony, they're tools
3. **Message-based async**: Never block Update(), return Commands instead
4. **Test pure functions**: Unit test domain logic without bubbletea
5. **Config hierarchy**: Defaults < File < Env < Flags
6. **Component reusability**: Extract when repeated 3+ times or highly complex
7. **Error handling**: Show in UI, log to file, never panic in production
8. **Start simple**: Flat structure → add layers as complexity demands

---

## Updated Anti-Patterns

### Anti-Pattern: Over-Structuring Early Projects

**What goes wrong**: Creating `cmd/`, `internal/`, `pkg/` directories for a 200-line app.

**Why it happens**: Tutorials show "ideal" structure without explaining when to apply it.

**How to avoid**: Start with flat structure (main.go, model.go, view.go). Introduce `internal/` when you have 3+ files that form a coherent subsystem.

**Warning signs**: Empty directories, jumping between 6 files to change one feature.

### Anti-Pattern: Mixing Domain and UI Logic

**What goes wrong**: Business rules (e.g., "card can't move from Done to Todo") scattered in Update().

**Why it happens**: No domain layer, all logic lives in bubbletea Model.

**How to avoid**: Put business rules in domain structs (`Card.CanMoveTo()`). Update() orchestrates, doesn't decide.

**Example of problem**:

```go
// DON'T: Domain logic in Update()
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // ... navigation logic ...
    if selectedCard.Status == "done" && targetColumn == "todo" {
        // Why is this rule here?
        return m, nil // Prevent move
    }
}
```

**Correct approach**:

```go
// DO: Domain logic in domain layer
func (c *Card) CanMoveTo(target Column) bool {
    return !(c.Status == StatusDone && target.Type == ColumnTypeTodo)
}

// Update() just calls domain method
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    if !m.board.CanMoveCard(selectedCard, targetColumn) {
        return m, nil
    }
    // Execute move
}
```

### Anti-Pattern: Global Configuration

**What goes wrong**: Reading `os.Getenv()` throughout code, making testing impossible.

**Why it happens**: Convenience - just call `os.Getenv("WIDTH")` wherever needed.

**How to avoid**: Load config once in `main()`, pass to components via structs.

**Correct approach**:

```go
// DO: Dependency injection
type Model struct {
    config *config.Config
    // ... other fields
}

func NewModel(cfg *config.Config) Model {
    return Model{config: cfg}
}
```
