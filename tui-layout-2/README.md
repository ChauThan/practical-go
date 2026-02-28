# Kanban Board TUI Demo

A terminal user interface (TUI) demonstration app built in Go using [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss) libraries.

## Purpose

This is an educational project demonstrating the **Elm Architecture pattern** (Model-Update-View) for building terminal UIs in Go. It does NOT implement real kanban logic, backend, or persistence — all content is static mock data.

## Prerequisites

- Go 1.24 or later

## Running

```bash
# Run the application
go run ./cmd/kanban

# Or build and run
go build -o kanban ./cmd/kanban
./kanban
```

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
