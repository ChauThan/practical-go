package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AuthHandler handles authentication-related requests.
type AuthHandler struct {
	jwtSecret string
}

// NewAuthHandler creates a new AuthHandler with the given JWT secret.
func NewAuthHandler(jwtSecret string) *AuthHandler {
	return &AuthHandler{jwtSecret: jwtSecret}
}

// TokenRequest represents the request body for generating a token.
type TokenRequest struct {
	Username string `json:"username"`
}

// TokenResponse represents the response body containing the token.
type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

// GenerateToken handles POST /auth/token - generates a test JWT token.
func (h *AuthHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Username == "" {
		req.Username = "testuser"
	}

	// Create token with claims
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"sub":  req.Username,
		"name": req.Username,
		"iat":  time.Now().Unix(),
		"exp":  expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := TokenResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt.Format(time.RFC3339),
	}

	respondJSON(w, http.StatusOK, response)
}
