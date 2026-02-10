package ui

const (
	minBoxWidth     = 60
	minTermWidth    = 80
	minTermHeight   = 24
	searchBoxHeight = 3
	hMargin         = 2
	vSectionGap     = 1
	selectorLimit   = 156
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

	// Search box always occupies exactly searchBoxHeight lines of CONTENT.
	// Actual rendered height includes: borders(2) + padding(2) + bottomMargin(1)
	// Total search height = searchBoxHeight + 5
	// Results section takes the remaining height.
	actualSearchBoxHeight := searchBoxHeight + 5 // borders + padding + margin
	m.resultsListHeight = m.windowHeight - actualSearchBoxHeight - vSectionGap
	if m.resultsListHeight < 1 {
		m.resultsListHeight = 1
	}

	// Inner list content width/height for the list component
	// Need to account for box borders and padding
	// focusedBox/unfocusedBox use Padding(1) and RoundedBorder (1 char on each side)
	innerWidth := m.boxWidth - 4 // 2 for borders, 2 for padding
	if innerWidth < 1 {
		innerWidth = 1
	}

	// Adjust results list height for box chrome inside results section:
	// Results section reserves resultsListHeight lines for the entire box
	// Actual inner space = resultsListHeight - borders(2) - padding(2) = resultsListHeight - 4
	// But we also need space for title line + margin (2 lines)
	innerListHeight := m.resultsListHeight - 6 // borders(2) + padding(2) + title+marginTop(2)
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
