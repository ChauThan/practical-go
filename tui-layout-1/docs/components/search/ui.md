# Search Section

```
┌─────────────────────────────────────────────┐
│              Search Section (1)             │
│  Type to search...                          │
└─────────────────────────────────────────────┘
```

## Description

The Search Section is the top section of the TUI layout where users can input their search queries. The input field has no inner border, providing a clean and minimalist interface.

## UI Specifications

### Section Dimensions
- **Height**: Fixed at 3 lines
- **Width**: Full terminal width (100%)
- **Position**: Top section, starting at line 0

### Visual Elements
- **Outer Border**: Single-line border around the entire search section
- **Header Text**: "Search Section (1)" centered at the top of the section
- **Header Position**: Line 1 of the section
- **Header Alignment**: Center horizontally
- **Header Color**: Cyan (`#5fd787` or color 86) when section is focused, gray (`#606060` or color 241) when not focused
- **Input Field Position**: Line 2 of the section
- **Input Field Background**: Default terminal background
- **Text Color**: Default terminal text color

### Input Field
- **Visual Style**: No inner border around the input field
- **Input Width**: Section width minus 4 characters (2 padding on each side)
- **Padding**: 1 character of padding on each side
- **Cursor Display**: Visible when input field is focused, highlighting the current character position
- **Placeholder Text**: "Type to search..." displayed when empty and not focused
- **Text Alignment**: Left-aligned
- **Maximum Length**: Unlimited (based on available space)

### Focus States

#### Section Focus
- When the search section is focused (but input field is not):
  - Header text is cyan and bold
  - Help bar displays: "i: Focus input field"
  - User can press "i" to transition focus to the input field

#### Input Field Focus
- When the input field is focused:
  - Cursor is visible at the current character position
  - Help bar displays: "q: Quit input field"
  - User can type text into the input field
  - User can press "q" to transition focus back to the search section

### Behavior
- **Default Focus**: Search section receives focus on application startup (not the input field directly)
- **Focus Transition**: Press "i" to enter input field mode, press "q" to exit
- **State Persistence**: Input field content is preserved when switching between section focus and input field focus
