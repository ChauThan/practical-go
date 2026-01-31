package services

import (
	"go-bookstore-api/internal/models"
	"go-bookstore-api/internal/repositories"
)

// BookService defines the interface for book business logic.
type BookService interface {
	CreateBook(book *models.Book) error
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

// bookService implements BookService.
type bookService struct {
	repo repositories.BookRepository
}

// NewBookService creates a new BookService with the given repository.
func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
}

// CreateBook creates a new book.
func (s *bookService) CreateBook(book *models.Book) error {
	return s.repo.Create(book)
}

// GetAllBooks retrieves all books.
func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

// GetBookByID retrieves a book by its ID.
func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
	return s.repo.FindByID(id)
}

// UpdateBook updates an existing book.
func (s *bookService) UpdateBook(book *models.Book) error {
	return s.repo.Update(book)
}

// DeleteBook deletes a book by its ID.
func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
