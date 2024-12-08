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
	"github.com/qaiswardag/go_backend_api_jwt/internal/utils"
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
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Handler login
func Create(w http.ResponseWriter, r *http.Request) {

	utils.RemoveCookie(w, "session_token", true)
	utils.RemoveCookie(w, "csrf_token", false)

	fileLogger := logger.FileLogger{}

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

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Check if the record already exists
	userByUsername := model.User{}
	if utils.CheckIfRecordExists(db, &userByUsername, "user_name", req.Username, w, fileLogger) {
		return
	}

	// Check if the record already exists
	userByEmail := model.User{}
	if utils.CheckIfRecordExists(db, &userByEmail, "email", req.Email, w, fileLogger) {
		return
	}

	sessionToken := tokengen.GenerateRandomToken(32)
	utils.SetCookie(w, "session_token", sessionToken, true)
	// Store the session_token in the database

	csrfToken := tokengen.GenerateRandomToken(32)
	utils.SetCookie(w, "csrf_token", csrfToken, false)

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	fileLogger.LogToFile("PASSWORD", fmt.Sprintf("Hashed password: %s", string(hashedPassword)))

	// Create a User object
	newUser := model.User{
		UserName:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  string(hashedPassword),
	}

	// Start a new transaction
	tx := db.Begin()
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		fileLogger.LogToFile("DB", fmt.Sprintf("Failed to start transaction. Error: %s", tx.Error.Error()))
		return
	}

	// Save the user to the database
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		fileLogger.LogToFile("AUTH", fmt.Sprintf("Failed to save user to database. Error: %s", err.Error()))
		return
	}

	// Create a Session object
	session := &model.Session{
		UserID:      int(newUser.ID),
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
		fileLogger.LogToFile("AUTH", fmt.Sprintf("Failed to save session to database. Error: %s", err.Error()))
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		fileLogger.LogToFile("AUTH", fmt.Sprintf("Failed to commit transaction. Error: %s", err.Error()))
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}
