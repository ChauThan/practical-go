# Search Section

```
┌─────────────────────────────────────────────┐
│              Search Section                 │
│  ┌─────────────────────────────────────┐   │
│  │ [Search Input Field]                 │   │
│  └─────────────────────────────────────┘   │
└─────────────────────────────────────────────┘
```

## Description

The Search Section is the top section of the TUI layout where users can input their search queries.

## UI Specifications

### Section Dimensions
- **Height**: Fixed at 5 lines
- **Width**: Full terminal width (100%)
- **Position**: Top section, starting at line 0

### Visual Elements
- **Outer Border**: Single-line border around the entire search section
- **Header Text**: "Search Section" centered at the top of the section
- **Header Position**: Line 1 of the section
- **Header Alignment**: Center horizontally
- **Inner Border**: Single-line border around the input field
- **Input Field Padding**: 1 character on each side
- **Input Field Position**: Line 3 of the section
- **Input Field Background**: Default terminal background
- **Text Color**: Default terminal text color

### Input Field
- **Input Width**: Section width minus 4 characters (2 for padding, 2 for borders)
- **Cursor Display**: Visible when focused, highlighting the current character position
- **Placeholder Text**: "Type to search..." displayed when empty (optional)
- **Text Alignment**: Left-aligned
- **Maximum Length**: Unlimited (or based on available space)

### Focus State
- **Default Focus**: Search section receives focus on application startup
- **Focus Indicator**: Input field shows cursor when focused
- **Border Enhancement**: Double-line or highlighted border when focused (optional)
