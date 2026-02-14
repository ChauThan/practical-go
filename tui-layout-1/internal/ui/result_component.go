package ui

import "strings"

func renderResultComponent(width int, height int) string {
	if height < 5 {
		height = 5
	}

	items := []string{
		"• Result item 1",
		"• Result item 2",
		"• Result item 3",
		"• ...",
	}

	contentLines := make([]string, 0, height)
	contentLines = append(contentLines, titledTopBorder(width, "Result Section"))
	contentLines = append(contentLines, innerTopLine(width))

	innerRows := height - 4
	for index := 0; index < innerRows; index++ {
		lineText := ""
		if index < len(items) {
			lineText = items[index]
		}
		contentLines = append(contentLines, innerContentLine(width, lineText))
	}

	contentLines = append(contentLines, innerBottomLine(width))
	contentLines = append(contentLines, bottomBorder(width))

	return strings.Join(contentLines, "\n")
}
