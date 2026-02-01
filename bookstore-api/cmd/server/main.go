package main

import (
	"log"
	"net/http"

	"bookstore-api/internal/config"
	"bookstore-api/internal/database"
	"bookstore-api/internal/handlers"
	"bookstore-api/internal/middleware"
	"bookstore-api/internal/repositories"
	"bookstore-api/internal/services"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")

	// Initialize layers
	bookRepo := repositories.NewGormBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)
	authHandler := handlers.NewAuthHandler(cfg.JWTSecret)

	// Setup router
	mux := http.NewServeMux()

	// Auth middleware
	authMiddleware := middleware.Auth(cfg.JWTSecret)

	// Auth routes (public - for testing)
	mux.HandleFunc("POST /auth/token", authHandler.GenerateToken)

	// Public routes (no auth required)
	mux.HandleFunc("GET /books", bookHandler.GetAllBooks)
	mux.HandleFunc("GET /books/{id}", bookHandler.GetBookByID)

	// Protected routes (auth required)
	mux.Handle("POST /books", authMiddleware(http.HandlerFunc(bookHandler.CreateBook)))
	mux.Handle("PUT /books/{id}", authMiddleware(http.HandlerFunc(bookHandler.UpdateBook)))
	mux.Handle("DELETE /books/{id}", authMiddleware(http.HandlerFunc(bookHandler.DeleteBook)))

	// Wrap with logging middleware
	handler := middleware.Logging(mux)

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
