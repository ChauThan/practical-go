package repositories

import (
	"go-bookstore-api/internal/models"

	"gorm.io/gorm"
)

// BookRepository defines the interface for book data access.
type BookRepository interface {
	Create(book *models.Book) error
	FindAll() ([]models.Book, error)
	FindByID(id uint) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
}

// gormBookRepository implements BookRepository using GORM.
type gormBookRepository struct {
	db *gorm.DB
}

// NewGormBookRepository creates a new BookRepository using GORM.
func NewGormBookRepository(db *gorm.DB) BookRepository {
	return &gormBookRepository{db: db}
}

// Create inserts a new book into the database.
func (r *gormBookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

// FindAll retrieves all books from the database.
func (r *gormBookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

// FindByID retrieves a book by its ID.
func (r *gormBookRepository) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Update modifies an existing book in the database.
func (r *gormBookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

// Delete removes a book from the database by its ID.
func (r *gormBookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}
