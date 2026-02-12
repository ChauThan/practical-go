package ui

const (
	minBoxWidth     = 60
	minTermWidth    = 80
	minTermHeight   = 24
	searchBoxHeight = 2
	hMargin         = 2
	vSectionGap     = 1
	selectorLimit   = 156
	titleHeight     = 1 // Height of title lines (outside boxes)
)

// Layout holds the computed layout dimensions for the TUI.
type Layout struct {
	WindowWidth       int
	WindowHeight      int
	BoxWidth          int
	ResultsListHeight int
	TooSmall          bool
}

// RecomputeLayout recalculates layout dimensions based on current window size.
// It updates the provided Model's layout fields and adjusts the sizes of
// the search input and results list components.
func RecomputeLayout(m *Model) {
	if m.windowWidth == 0 || m.windowHeight == 0 {
		return
	}

	// Determine if terminal is too small
	m.tooSmall = m.windowWidth < minTermWidth || m.windowHeight < minTermHeight

	// Calculate box width: min of terminal width minus margins and minBoxWidth,
	// clamped to minBoxWidth if terminal is very small.
	maxBoxWidth := m.windowWidth - 2*hMargin
	if maxBoxWidth < minBoxWidth {
		m.boxWidth = minBoxWidth
	} else {
		m.boxWidth = maxBoxWidth
	}

	// Layout with titles outside boxes:
	// Search section: title(1) + titleMargin(1) + searchBoxHeight
	// Search box is auto-sized from content: borders(2) + padding(2) + content(1) = 5
	// Search section total: title(1) + titleMargin(1) + searchBox(5) = 7 lines
	// Then add 1 line margin before results section in View()
	// For a 24-line terminal: 7 + 1(margin) = 8, remaining for results section = 24 - 8 = 16 lines
	// Results section in View() is: marginTop(1) + title(1) + titleMargin(1) + resultsBox(X)
	// So: 1 + 1 + 1 + boxHeight = 16
	// boxHeight = 16 - 3 = 13 lines
	actualSearchSectionHeight := 7 // 1(title) + 1(titleMargin) + 5(search box)

	// Calculate the actual box Height(X) to set:
	// m.resultsListHeight = windowHeight - 7 (search section) - 1 (margin between)
	// results section needs: marginTop(1) + title(1) + titleMargin(1) + boxHeight(X)
	// boxHeight = m.resultsListHeight - 3 - 1 (reduce by 1 more to show title)
	m.resultsListHeight = m.windowHeight - actualSearchSectionHeight - 1
	m.resultsListHeight = m.resultsListHeight - 4 // subtract marginTop(1) + title(1) + titleMargin(1) + 1(reduction)

	if m.resultsListHeight < 4 {
		m.resultsListHeight = 4 // minimum box height: borders(2) + padding(2)
	}

	// Inner list content width/height for the list component
	// Need to account for box borders and padding
	// focusedBox/unfocusedBox use Padding(1) and RoundedBorder (1 char on each side)
	innerWidth := m.boxWidth - 4 // 2 for borders, 2 for padding
	if innerWidth < 1 {
		innerWidth = 1
	}

	// Adjust results list height for box chrome:
	// m.resultsListHeight is the total box height including borders(2) + padding(2) + content
	// So inner list content height = resultsListHeight - borders(2) - padding(2) = resultsListHeight - 4
	innerListHeight := m.resultsListHeight - 4 // 2(borders) + 2(padding)
	if innerListHeight < 1 {
		innerListHeight = 1
	}

	// Apply sizing to the list
	m.resultsList.SetSize(innerWidth, innerListHeight)

	// Update search input width to fit inside box
	inputWidth := innerWidth
	if inputWidth > selectorLimit {
		inputWidth = selectorLimit
	}
	m.searchInput.Width = inputWidth
	m.searchInput.PlaceholderStyle = NewStyleProvider().NewStyle().Width(inputWidth)
}
