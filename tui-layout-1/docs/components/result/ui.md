# Result Section

```
┌─────────────────────────────────────────────┐
│              Result Section                 │
│  ┌─────────────────────────────────────┐   │
│  │ • Result item 1                      │   │
│  │ • Result item 2                      │   │
│  │ • Result item 3                      │   │
│  │   ...                                │   │
│  └─────────────────────────────────────┘   │
└─────────────────────────────────────────────┘
```

## Description

The Result Section is the bottom section of the TUI layout that displays search results in a scrollable list format.

## UI Specifications

### Section Dimensions
- **Height**: Remaining terminal height after Search Section (variable)
- **Width**: Full terminal width (100%)
- **Position**: Bottom section, starting after Search Section (line 5)

### Visual Elements
- **Outer Border**: Single-line border around the entire result section
- **Header Text**: "Result Section" centered at the top of the section
- **Header Position**: Line 1 of the section
- **Header Alignment**: Center horizontally
- **Inner Border**: Single-line border around the result list
- **Result List Padding**: 1 character on each side
- **Result List Position**: Starting at line 3 of the section
- **Result List Background**: Default terminal background
- **Text Color**: Default terminal text color

### Result List
- **List Width**: Section width minus 4 characters (2 for padding, 2 for borders)
- **List Height**: Section height minus 4 lines (header + top border + bottom border)
- **Item Separator**: Each item on its own line
- **Bullet Character**: "• " (bullet point followed by space) prefix for each item
- **Top Item Position**: Line 3 of the section (first visible item)
- **Empty State**: "No results found" message displayed when list is empty
- **Empty State Alignment**: Center horizontally and vertically

### Item Styling
- **Item Text**: Left-aligned with bullet point prefix
- **Max Text Length**: Truncate with "..." if exceeds available width
- **Highlighting**: Selected item highlighted with inverse colors or different background (optional)
- **Selection Indicator": ">" or different bullet style for currently selected item (optional)

### Scrollbar
- **Scrollbar Position**: Right edge of the inner border (optional)
- **Visual Style**: Simple character (│, █, or similar)
- **Thumb Size**: Proportional to visible items vs. total items
- **Thumb Position**: Based on current scroll position

### Focus State
- **Focus Indicator": Title or border highlighted when focused (optional)
- **Navigation Indicator": Selected row with different styling when focused (optional)
