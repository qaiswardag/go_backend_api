package userregistercontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
	"github.com/qaiswardag/go_backend_api_jwt/internal/security/tokengen"
	"golang.org/x/crypto/bcrypt"
)

/*
   |--------------------------------------------------------------------------
   | Controller Method Naming Convention
   |--------------------------------------------------------------------------
   |
   | Controller methods: index, create, store, show, edit, update, destroy.
   | Please aim for consistency by using these method names in all controllers.
   |
*/

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handler login
func Create(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	var req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure the body is closed after reading
	defer r.Body.Close()

	// Access the username and password
	logger.LogToFile(fmt.Sprintf("Received Email: %s\n", req.Email))
	logger.LogToFile(fmt.Sprintf("Received password: %s\n", req.Password))

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	logger.LogToFile(fmt.Sprintf("Hashed password: %s\n", string(hashedPassword)))

	database.InitDB()

	w.WriteHeader(http.StatusUnauthorized)
	return

	sessionToken := tokengen.GenerateRandomToken(32)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
	// Store the session_token in the database

	csrfToken := tokengen.GenerateRandomToken(32)
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})
	// Store csrf_token token in database

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(model.UserObject()); err != nil {
		log.Printf("Error encoding JSON response: %v\n", err)
	}
}
