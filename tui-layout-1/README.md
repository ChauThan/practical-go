# TUI Layout

A simple Terminal User Interface (TUI) application with search and result sections.

## Layout Overview

```
┌─────────────────────────────────────────────┐
│              Search Section                 │
│  ┌─────────────────────────────────────┐   │
│  │ [Search Input Field]                 │   │
│  └─────────────────────────────────────┘   │
└─────────────────────────────────────────────┘
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

## Components

### 1. Search Section
- Input field for entering search queries
- Processes user input and triggers searches

### 2. Result Section
- Displays search results in a scrollable list
- Shows items returned by the search query

## Getting Started

```bash
# Run the application
go run .
```

## Technology Stack

- Go 1.23.6
- Bubble Tea framework for TUI
