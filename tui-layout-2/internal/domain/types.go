package domain

// Card represents a task card in the kanban board
type Card struct {
	Title string
}

// Column represents a kanban column containing cards
type Column struct {
	Title  string
	Cards  []Card
}
