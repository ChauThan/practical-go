package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-bookstore-api/internal/models"
	"go-bookstore-api/internal/services"
)

// BookHandler handles HTTP requests for books.
type BookHandler struct {
	service services.BookService
}

// NewBookHandler creates a new BookHandler with the given service.
func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook handles POST /books
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	respondJSON(w, http.StatusCreated, book)
}

// GetAllBooks handles GET /books
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch books")
		return
	}

	respondJSON(w, http.StatusOK, books)
}

// GetBookByID handles GET /books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Book not found")
		return
	}

	respondJSON(w, http.StatusOK, book)
}

// UpdateBook handles PUT /books/{id}
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	book.ID = id
	if err := h.service.UpdateBook(&book); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update book")
		return
	}

	respondJSON(w, http.StatusOK, book)
}

// DeleteBook handles DELETE /books/{id}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	respondJSON(w, http.StatusNoContent, nil)
}

// extractIDFromPath extracts the book ID from the request path parameter.
func extractIDFromPath(r *http.Request) (uint, error) {
	idStr := r.PathValue("id")
	if idStr == "" {
		return 0, http.ErrNotSupported
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// respondJSON writes a JSON response with the given status code and data.
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// respondError writes an error response with the given status code and message.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
