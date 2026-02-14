# Help Bar

```
┌─────────────────────────────────────────────┐
│              Search Section                 │
│  ┌─────────────────────────────────────┐   │
│  │ [Search Input Field]                 │   │
│  └─────────────────────────────────────┘   │
└─────────────────────────────────────────────┘
┌─────────────────────────────────────────────┐
│              Result Section                 │
│  ┌─────────────────────────────────────-┐   │
│  │ • Result item 1                       │   │
│  │ • Result item 2                       │   │
│  └─────────────────────────────────────-┘   │
└─────────────────────────────────────────────┘
──────────────────────────────────────────────
│ ↑/↓: Navigate | Enter: Select | q: Quit    │
```

## Description

The Help Bar is the bottommost section of the TUI layout that displays keyboard shortcuts and context-sensitive hints to guide users on available actions.

## UI Specifications

### Section Dimensions
- **Height**: Fixed at 1 line for hints (plus 1 line separator if used)
- **Width**: Full terminal width (100%)
- **Position**: Bottommost section, after result section
- **Separator Line**: Optional separator line (─) between result section and help bar

### Visual Elements
- **Separator Line**: Single-line horizontal separator (─ characters) spanning full width (optional)
- **Hint Items**: Displayed as key-value pairs separated by delimiters
- **Delimiter**: " | " (pipe with spaces) between hint items
- **Key Format**: Keyboard shortcuts displayed in a distinct style (e.g., `↑/↓`, `Enter`, `q`)
- **Key Styling**: May use color, underline, or brackets (e.g., `[]`) to highlight keys
- **Description Text**: Plain text explaining what the shortcut does
- **Text Alignment**: Left-aligned or center-aligned across hints
- **Text Color**: Default terminal text color (or muted color for non-critical hints)

### Hint Items
- **Format**: `{key}: {description}`
- **Spacing**: 1 space after key, 1 space before/after delimiter
- **Key Display Style Options**:
  - Plain text: `↑/↓: Navigate`
  - With brackets: `[↑/↓]: Navigate`
  - Colored/underlined: highlight key for emphasis
- **Max Width**: Truncate hints from right if they exceed available space (optional)

### Context-Sensitive Hints
- **Dynamic Content**: Hints change based on currently focused section
- **Search Section Focus**: Show search-related hints (e.g., `Enter: Search`, `Ctrl+C: Clear`)
- **Result Section Focus**: Show navigation hints (e.g., `↑/↓: Navigate`, `Enter: Select`, `j/k: Scroll`)
- **Global Hints**: Always display application-wide hints (e.g., `q: Quit`, `Ctrl+C: Exit`)

### Common Hint Examples
- Navigation: `↑/↓: Navigate`, `j/k: Scroll`, `Home/End: Jump to start/end`
- Actions: `Enter: Select`, `Escape: Cancel`, `Space: Toggle`
- Search: `/`: Search`, `Ctrl+F`: Find next`, `Ctrl+R`: Refresh`
- Global: `q: Quit`, `?`: Show help`, `Ctrl+C: Exit`

### Focus State
- **Default State**: No focus indicator (help bar is not focusable)
- **Active Indicator**: Optional highlight or border when hints are relevant to current action
- **Priority**: Always visible overlay on top of content (never scrolled)

### Responsive Behavior
- **Terminal Resize**: Hints re-wrap or truncate to fit new width
- **Small Terminals**: Hide less critical hints when width is limited
- **Hint Priority**: Global hints (quit, exit) have lowest priority for hiding
- **Ellipsis**: Show "..." when hints are hidden

### Visual Variations

#### Style 1: Compact (single line)
```
↑/↓: Navigate | Enter: Select | q: Quit
```

#### Style 2: With Separator
```
─────────────────────────────
↑/↓: Navigate | Enter: Select | q: Quit
```

#### Style 3: With brackets
```
[↑/↓] Navigate | [Enter] Select | [q] Quit
```

#### Style 4: Multi-group (sections)
```
[Navigate] ↑/↓ | [Action] Enter | [Global] q
```

### Separator Line
- **Style**: Single-line horizontal line (─)
- **Position**: 1 line above hint text
- **Width**: Full terminal width
- **Visibility**: Optional - can be enabled/disabled via configuration
