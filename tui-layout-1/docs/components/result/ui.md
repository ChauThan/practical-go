# Result Section

```
┌─────────────────────────────────────────────┐
│              Result Section                 │
│  • Result item 1                      │     │
│  • Result item 2                      │[█]  │
│  • Result item 3                      │     │
│  ...                               │     │
│                                      │     │
│  Result n                       │[▓▓]│     │
```

## Description

The Result Section is the bottom section of the TUI layout that displays search results in a scrollable list format with keyboard navigation and selection highlighting.

## UI Specifications

### Section Dimensions
- **Height**: Remaining terminal height after Search Section (variable)
- **Width**: Full terminal width (100%)
- **Position**: Bottom section, starting after Search Section (line 5)

### Visual Elements
- **Outer Border**: **Removed** - No border around the result section
- **Header Text**: "Result Section" with the section number (e.g., "Result Section (2)")
- **Header Position**: Line 1 of the section
- **Header Alignment**: Center horizontally
- **Header Color**: Color `86` (cyan) when focused, Color `241` (gray) when unfocused
- **Header Style**: Bold when focused, normal when unfocused
- **Inner Border**: **Removed** - No border around the result list
- **Result List Padding**: 1 character on left side
- **Result List Position**: Starting at line 3 of the section
- **Result List Background**: Default terminal background
- **Text Color**: Default terminal text color

### Result List
- **List Width**: Section width minus 6 characters (2 for left padding, 4 for scrollbar)
- **List Height**: Section height minus 4 lines (header + top line + bottom line + bottom border)
- **Item Separator**: Each item on its own line
- **Bullet Character**: "• " (bullet point followed by space) prefix for each item
- **Top Item Position**: Line 3 of the section (first visible item)
- **Empty State**: "No results found" message displayed when list is empty
- **Empty State Alignment**: Center horizontally and vertically

### Item Styling
- **Item Text**: Left-aligned with bullet point prefix
- **Max Text Length**: Truncate with "..." if exceeds available width
- **Selection Highlighting**: Selected item uses **inverse colors** (text/background swap) OR distinct background color (Color `236`)
- **Selection Persistence**: Maintained when switching focus between sections
- **Visual Contrast**: Selection must meet WCAG AA minimum 4.5:1 contrast ratio

### Scrollbar
- **Scrollbar Position**: Right edge of the result section (within inner border)
- **Scrollbar Width**: 1 character
- **Track Character**: `│` (vertical line)
- **Thumb Character**: `▓` or `█` (filled block)
- **Thumb Size**: Calculated proportionally: `int(float64(visibleCount) / float64(totalItems) * visibleHeight)`
- **Thumb Position**: Based on current scroll offset: `int(float64(scrollOffset) / float64(totalItems - visibleCount) * freeSpace)`
- **Visibility**: Always visible when content exceeds display area
- **Behavior**: Updates in real-time during scroll navigation

### Keyboard Navigation
- **Navigation Trigger**: Active when `ResultFocused` state is set
- **Up Arrow**: Move selection to previous item
- **Down Arrow**: Move selection to next item
- **Auto-Scroll**: Automatically scroll to keep selected item visible
- **Scroll Down Threshold**: When `selectedIdx >= scrollOffset + visibleCount - 1`
- **Scroll Up Threshold**: When `selectedIdx < scrollOffset`
- **Boundary Handling**: Stop at first/last item (no wrap in MVP)

### Focus State
- **Focus Indicator**: Title color changes (cyan when focused, gray when unfocused)
- **Title Style**: Bold when focused, normal when unfocused
- **Navigation Active**: Arrow keys only respond when result section is focused
- **Focus Switching**: Use `2` key to switch focus to result section

## Behavior Details

### Scrolling Logic

The result section implements auto-scrolling to keep the selected item visible:

```go
// When moving down
if selectedIdx >= scrollOffset + visibleCount - 1 {
    scrollOffset = selectedIdx - visibleCount + 1
}

// When moving up
if selectedIdx < scrollOffset {
    scrollOffset = selectedIdx
}
```

**Key Points**:
- `scrollOffset` represents the index of the first visible item
- `visibleCount` is the number of items that fit in the display area
- Scroll only occurs when selection would go out of view
- Selection is always visible after any navigation action

### Scrollbar Calculation

**Thumb Size Formula**:
```
thumbSize = max(1, int(float64(visibleCount) / float64(totalItems) * availableHeight))
```

**Thumb Position Formula**:
```
freeSpace = availableHeight - thumbSize
thumbPosition = int(float64(scrollOffset) / float64(totalItems - visibleCount) * freeSpace)
```

**Edge Cases**:
- Single item: Thumb size = 1
- Exact fit: Thumb size = availableHeight
- No scrollbar needed (items fit): Hide scrollbar or show full-width thumb

### State Management

The Model tracks the following state for the result section:

```go
type Model struct {
    // ... existing fields
    resultItems   []string // All result items
    scrollOffset  int      // First visible item index
    selectedIdx   int      // Currently selected item index
    visibleCount  int      // Number of items visible in view (recalculated on resize)
}
```

**State Initialization** (in `NewModel()`):
```go
scrollOffset  = 0
selectedIdx   = 0
resultItems   = []string{"Result item 1", "Result item 2", ...}
visibleCount  = calculated based on height
```

## UI Visual Examples

### Scenario 1: Fewer Items Than Display Area (No Scroll)

```
┌─────────────────────────────────────────────┐
│             Result Section (2)             │
│  • Result item 1                      │     │
│  • Result item 2                    [▓▓▓▓]│
│  • Result item 3                      │     │
│  [empty space...]                     │     │
│                                      │     │
└─────────────────────────────────────────────┘
```
- Scrollbar thumb spans full height (all items visible)
- Selection on item 3 with highlighting

### Scenario 2: More Items Than Display Area (Scroll Required)

```
┌─────────────────────────────────────────────┐
│             Result Section (2)             │
│  • Result item 5                      │     │
│  • Result item 6                      │    █│
│  • Result item 7                      │     │
│  • Result item 8                      │     │
│                                      │     │
└─────────────────────────────────────────────┘
```
- Only items 5-8 visible (scrollOffset = 4)
- Scrollbar thumb small (indicating more content below)
- Selection on item 6 with highlighting

### Scenario 3: Bottom of List Selected

```
┌─────────────────────────────────────────────┐
│             Result Section (2)             │
│  • Result item 47                     │     │
│  • Result item 48                     │    █│
│  • Result item 49                     │     │
│  • Result item 50                     │     │
│                                      │     │
└─────────────────────────────────────────────┘
```
- Last batch of items visible (scrollOffset = 46)
- Scrollbar thumb at bottom position
- Selection on item 48 with highlighting

### Focus State Comparison

**Unfocused Result Section** (Search section focused):
```
┌─────────────────────────────────────────────┐
│             Result Section (2)             │  ← Gray, not bold
│  • Result item 1                      │     │
│  • Result item 2                      │    █│
│  • Result item 3                      │     │
│  ...                                │     │
└─────────────────────────────────────────────┘
```

**Focused Result Section** (active navigation):
```
┌─────────────────────────────────────────────┐
│             Result Section (2)             │  ← Cyan, bold
│  • Result item 1                      │     │
│  • Result item 2                      │    █│
│  • Result item 3                      │     │
│  ...                                │     │
└─────────────────────────────────────────────┘
```

## Implementation Notes

### Rendering Optimization

To maintain performance with large result sets:
- Only render visible items (not entire `resultItems` array)
- Slice visible items: `resultItems[scrollOffset : scrollOffset+visibleCount]`
- Cache styled item strings when content hasn't changed

### Resize Handling

When terminal is resized:
```go
case tea.WindowSizeMsg:
    m.width = typed.Width
    m.height = typed.Height
    // Recalculate viewport
    m.visibleCount = calculateVisibleCount(m.height - topPadding - bottomPadding)
    // Adjust scroll position if new size would make selection invisible
    if m.selectedIdx >= m.scrollOffset + m.visibleCount {
        m.scrollOffset = m.selectedIdx - m.visibleCount + 1
    }
```

### Edge Cases to Handle

1. **Empty Results**:
   - Display "No results found" centered
   - Hide scrollbar
   - No selection state

2. **Single Item**:
   - Scrollbar thumb size = 1 character
   - Selection always on that item
   - Scrolling disabled

3. **Exact Fit** (items = visibleCount):
   - Scrollbar thumb fills entire available height
   - Scrolling disabled

4. **Very Long Item Text**:
   - Truncate with "..." at end
   - Preserve item prefix (bullet point)
   - Truncation should not affect selection area

## Accessibility Considerations

- **High Contrast**: Selection highlighting uses inverse colors or distinctive background
- **Color Independence**: Focus state indication uses both color AND bold styling
- **Keyboard Only**: All interactions available via keyboard (no mouse required)
- **Screen Reader Compatible**: Clear semantic structure, no ASCII art in data region
- **WCAG AA Compliance**: Contrast ratio >= 4.5:1 for selection vs. background

## Testing Checklist

- [ ] Result section renders without outer border
- [ ] Header displays correctly with focus state changes
- [ ] Scrollbar appears when results exceed display capacity
- [ ] Scrollbar thumb size is proportionally accurate
- [ ] Scrollbar thumb position reflects current scroll offset
- [ ] Arrow keys navigate through all result items
- [ ] Selection moves in correct direction (up/up, down/down)
- [ ] Auto-scroll keeps selection in view
- [ ] Selection highlighting is clearly visible
- [ ] Selection state persists when switching focus
- [ ] Window resize adjusts viewport correctly
- [ ] Empty state displays properly
- [ ] Single item scenario handled correctly
- [ ] Large dataset (100+ items) performance is acceptable
