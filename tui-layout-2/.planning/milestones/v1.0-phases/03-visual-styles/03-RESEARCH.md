# Phase 3: Visual Styles - Research

**Phase:** 3
**Researched:** 2026-02-28
**Tech Stack:** lipgloss v2.0.0

---

## Research Objective

Answer: "What do I need to know to implement lipgloss styles for a kanban board TUI?"

---

## Lipgloss v2 API Patterns

### Style Definition

Lipgloss uses a functional builder pattern for styles:

```go
import "charm.land/lipgloss/v2"

// Basic style
var style = lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    Padding(1).
    Width(20)

// Reusable style function
func columnStyle() lipgloss.Style {
    return lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        Padding(0, 1)
}
```

### Color Specification

Colors in lipgloss v2 use `lipgloss.Color()`:

```go
// Named colors
lipgloss.Color("blue")
lipgloss.Color("white")

// Hex colors
lipgloss.Color("#7C3AED")  // Purple
lipgloss.Color("#F59E0B")  // Amber
lipgloss.Color("#3B82F6")  // Blue

// Adaptive colors (light/dark terminal support)
lipgloss.AdaptiveColor{Light: "black", Dark: "white"}
lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}
```

### Focus State Patterns

For kanban board focus indicators:

1. **Column focus:** Use border color to distinguish active column
2. **Card focus:** Use border color + bold text to distinguish active card
3. **Inactive state:** Subtle/grayscale colors for unfocused elements

```go
// Inactive column style
var columnStyle = lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderColor(lipgloss.Color("245"))  // Gray

// Active column style
var activeColumnStyle = lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderColor(lipgloss.Color("#7C3AED"))  // Purple
```

### Text Styling

```go
// Bold text
lipgloss.NewStyle().Bold(true)

// Centered text
lipgloss.NewStyle().Align(lipgloss.Center)

// Dim/faint text (for help text)
lipgloss.NewStyle().Faint(true)

// Width constraints
lipgloss.NewStyle().Width(80)
```

---

## Style Requirements Mapping

| Requirement | Lipgloss API | Notes |
|-------------|--------------|-------|
| STYLE-01: columnStyle | Border(), Padding(), Width() | Normal border, subtle padding |
| STYLE-02: activeColumnStyle | BorderColor("#7C3AED") | Distinct purple border |
| STYLE-03: cardStyle | Border(), Padding() | Subtle border + padding |
| STYLE-04: activeCardStyle | BorderColor("#F59E0B"), Bold() | Amber border + bold |
| STYLE-05: titleStyle | Bold(), Center() | Column headers |
| STYLE-06: appTitleStyle | Bold(), Center(), Width() | App title, centered |
| STYLE-07: helpStyle | Faint(true) | Dimmed help text |

---

## Best Practices

1. **Export styles as variables or functions** - Use `var StyleName lipgloss.Style` or `func styleName() lipgloss.Style`
2. **Use constants for colors** - Define colors as package-level constants for consistency
3. **Consider terminal width** - Set minimum widths to prevent layout collapse
4. **Test with light/dark backgrounds** - Use AdaptiveColor for broader compatibility
5. **Keep styles separate from logic** - All style definitions in styles.go, no hardcoded styles in view code

---

## Implementation Considerations

### Style Storage Pattern

**Option A: Package-level variables**
```go
var columnStyle = lipgloss.NewStyle()...
var activeColumnStyle = lipgloss.NewStyle()...
```

**Option B: Style functions**
```go
func ColumnStyle() lipgloss.Style {
    return lipgloss.NewStyle()...
}
```

**Recommendation:** Option A (variables) is simpler and more idiomatic for this use case. Styles don't need runtime configuration.

### Border Styles

Lipgloss provides border styles:
- `lipgloss.NormalBorder()` - Standard `│─┼` borders
- `lipgloss.RoundedBorder()` - Rounded corners
- `lipgloss.DoubleBorder()` - Double-line borders
- `lipgloss.HiddenBorder()` - No border

**Recommendation:** Use `NormalBorder()` for columns and cards. Consider `RoundedBorder()` for app title.

### Padding vs Margins

- **Padding:** Internal space inside border
- **Margins:** External space outside border (not used in lipgloss v2)

**Recommendation:** Use Padding(0, 1) for horizontal spacing in columns, Padding(0) for cards to save vertical space.

---

## Terminal Compatibility

**Color support:**
- Most modern terminals support 256-color palette
- Hex colors (#RGB, #RRGGBB) work in terminals with truecolor support
- Fallback: Use 256-color palette codes (e.g., "245" for gray)

**Width constraints:**
- Set minimum width: `Width(25)` per column
- Total board width: 3 columns × 25 = 75 minimum (fits in 80×24 terminal)
- Consider `MaxWidth()` for very wide terminals

---

## Dependencies

No new dependencies required. Using existing lipgloss v2.0.0.

---

## Validation Strategy

1. **Compile check:** `go build ./internal/ui` must pass
2. **Style usage:** All styles referenced in Phase 4 View implementation
3. **Visual distinction:** activeColumnStyle vs columnStyle must be visually different
4. **Color contrast:** Ensure focus colors are visible on common terminal backgrounds

---

## Open Questions

None - lipgloss API is straightforward for this phase.

---

## RESEARCH COMPLETE

**Status:** Ready for planning

**Key findings:**
- Use lipgloss.NewStyle() builder pattern
- Define styles as package-level variables
- Use hex colors (#7C3AED, #F59E0B) for focus states
- NormalBorder() for standard look
- Export styles for use in Phase 4 View implementation

---

*Phase: 03-visual-styles*
*Research completed: 2026-02-28*
