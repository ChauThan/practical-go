package ui

import "strings"

func renderSearchComponent(width int) string {
	lines := []string{
		titledTopBorder(width, "Search Section"),
		innerTopLine(width),
		innerContentLine(width, "Type to search..."),
		innerBottomLine(width),
		bottomBorder(width),
	}

	return strings.Join(lines, "\n")
}
