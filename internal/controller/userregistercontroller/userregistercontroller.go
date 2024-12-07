package userregistercontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
	"github.com/qaiswardag/go_backend_api_jwt/internal/security/tokengen"
	"github.com/qaiswardag/go_backend_api_jwt/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	Username string `json:"username"`
	Password string `json:"password"`
}

// Handler login
func Create(w http.ResponseWriter, r *http.Request) {

	serverIP, err := utils.GetServerIP()

	if err != nil {
		log.Println("Failed to get serer ip.")
	}
	// Read the request body
	var req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure the body is closed after reading
	defer r.Body.Close()

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

	//
	//
	//
	//
	//
	//
	// Access the username and password
	logger.LogToFile("INPUT", fmt.Sprintf("Received username: %s", req.Username))
	logger.LogToFile("INPUT", fmt.Sprintf("Received email: %s", req.Email))
	logger.LogToFile("INPUT", fmt.Sprintf("Received password: %s", req.Password))

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	logger.LogToFile("PASSWORD", fmt.Sprintf("Hashed password: %s", string(hashedPassword)))

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Check if the username already exists
	var existingUser model.User
	if err := db.Where("user_name = ?", req.Username).First(&existingUser).Error; err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Username already taken"})
		logger.LogToFile("AUTH", fmt.Sprintf("Username already taken. Received username: %s", req.Username))
		return
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		logger.LogToFile("AUTH", fmt.Sprintf("Failed to query user. Error: %s", err.Error()))
		return
	}

	// Check if the email already exists
	var existingEmailUser model.User
	if err := db.Where("email = ?", req.Email).First(&existingEmailUser).Error; err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email already taken"})
		logger.LogToFile("AUTH", fmt.Sprintf("Email already taken. Received email: %s", req.Email))
		return
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		logger.LogToFile("AUTH", fmt.Sprintf("Failed to query email. Error: %s", err.Error()))
		return
	}

	// Create a User object
	user := &model.User{
		UserName:  req.Username,
		Email:     req.Email,
		FirstName: "john",
		LastName:  "doe",
		Password:  string(hashedPassword),
	}

	// Start a new transaction
	tx := db.Begin()
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		logger.LogToFile("AUTH", fmt.Sprintf("Failed to start transaction. Error: %s", tx.Error.Error()))
		return
	}

	// Save the user to the database
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		logger.LogToFile("AUTH", fmt.Sprintf("Failed to save user to database. Error: %s", err.Error()))
		return
	}

	// Create a Session object
	session := &model.Session{
		UserID:      int(user.ID),
		AccessToken: sessionToken,
		ServerIP:    serverIP,
		// Set expiry to 7 days
		AccessTokenExpiry: time.Now().Add(7 * 24 * time.Hour),
	}

	// Save the session to the database
	if err := tx.Create(session).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		logger.LogToFile("AUTH", fmt.Sprintf("Failed to save session to database. Error: %s", err.Error()))
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		logger.LogToFile("AUTH", fmt.Sprintf("Failed to commit transaction. Error: %s", err.Error()))
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}
