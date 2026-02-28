# Kanban Board TUI Demo

## What This Is

A terminal user interface (TUI) demonstration app built in Go using bubbletea and lipgloss libraries. The purpose is educational — demonstrating how to structure and build a TUI using Elm Architecture, without implementing real kanban logic, backend, persistence, or data manipulation. All content is static mock data.

## Core Value

**Demonstrate Elm Architecture pattern in Go for building terminal UIs with clear separation of Model, Update, and View.**

## Requirements

### Validated

(None yet — ship to validate)

### Active

- [ ] Phase 1: Project scaffold (Go module, dependencies, placeholder app)
- [ ] Phase 2: Model & state management (Card, Column, Model structs with navigation)
- [ ] Phase 3: Visual styles (lipgloss styles for all UI components)
- [ ] Phase 4: View & layout rendering (3-column kanban board with dynamic highlighting)
- [ ] Phase 5: Polish & responsive layout (terminal size adaptation, proper reflow)

### Out of Scope

- **Real kanban logic** — This is a demo only, no actual task management
- **Backend persistence** — No database or file storage
- **Data manipulation** — Static mock data only, no CRUD operations
- **Real-time updates** — No channels or event streaming
- **Mouse interaction** — Keyboard navigation only

## Context

**Technical Environment:**
- Go language (latest stable)
- Terminal UI libraries: bubbletea ( Elm Architecture framework), lipgloss (styling)
- Standard terminal sizes: 80×24, 120×40, 200×50

**Prior Work:**
- Existing TUI layout experience from feature/tui-layout-1 branch
- Prior implementation of search sections, result sections with keyboard navigation, selection highlighting, and scrollbar functionality

**Project Type:**
- Educational demonstration project
- Focus on architectural patterns over feature completeness

## Constraints

- **Stack**: Go + bubbletea + lipgloss (Charm libraries) — These are non-negotiable for learning Elm Architecture in terminal UIs
- **Architecture**: Must follow Elm Architecture pattern (Model / Update / View separation) — This is the core educational objective
- **Data**: Static mock data only — No backend, persistence, or data manipulation complexity
- **Scope**: 5 phases as defined in PRD — Keep focused on TUI fundamentals
- **Terminal Support**: Must render correctly at 80×24 minimum — Ensures broad terminal compatibility

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Elm Architecture pattern | Industry-tested pattern for predictable state management in UIs, widely used in web (Elm) and applicable to TUIs | — Pending |
| Static mock data | Focus on TUI structure and rendering without data management complexity | — Pending |
| 5-phase approach | Incremental building: scaffold → model → styles → view → polish | — Pending |

---
*Last updated: 2026-02-28 after project initialization*
