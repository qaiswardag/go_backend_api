package userregistercontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api/database"
	"github.com/qaiswardag/go_backend_api/internal/appconstants"
	"github.com/qaiswardag/go_backend_api/internal/logger"
	"github.com/qaiswardag/go_backend_api/internal/model"
	"github.com/qaiswardag/go_backend_api/internal/security/tokengen"
	"github.com/qaiswardag/go_backend_api/internal/utils"
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

func Create(w http.ResponseWriter, r *http.Request) {
	utils.RemoveCookie(w, "session_token", true)
	utils.RemoveCookie(w, "csrf_token", false)

	fileLogger := logger.FileLogger{}

	serverIP, errServerIP := utils.GetServerIP()

	if errServerIP != nil {
		log.Println("Failed to get serer ip.")
	}

	// Read the request body
	var req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request body"})
		return
	}

	// Ensure the body is closed after reading
	defer r.Body.Close()

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Ensure username is min 2 and maximum 50 characters
	if req.Username == "" || len(req.Username) < appconstants.MinTwoCharacters ||
		len(req.Username) > appconstants.MaxFiftyCharacters {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Username must be between 2 and 50 characters."})
		fileLogger.LogToFile("BAD REQUEST", "Username must be between 2 and 50 characters.")
		return
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
	utils.SetCookie(w, "session_token", sessionToken, false)

	csrfToken := tokengen.GenerateRandomToken(32)
	utils.SetCookie(w, "csrf_token", csrfToken, false)

	// Hash the password using bcrypt
	hashedPassword, errHashing := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHashing != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hashing error."})
		fileLogger.LogToFile("Hashing", "Hashing error.")
		return
	}

	// Create a User object
	newUser := model.User{
		UserName:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		// Convert hashed password from byte slice to string
		Password: string(hashedPassword),
	}

	// Start a new transaction
	tx := db.Begin()
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error."})
		fileLogger.LogToFile("DB", fmt.Sprintf("Failed to start transaction. Error: %s", tx.Error.Error()))
		return
	}

	// Save the user to the database
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error."})
		fileLogger.LogToFile("AUTH", fmt.Sprintf("Failed to save user to database. Error: %s", err.Error()))
		return
	}

	// Create a Session object
	session := &model.Session{
		UserID:             int(newUser.ID),
		SessionToken:       sessionToken,
		SessionTokenExpiry: time.Now().Add(appconstants.TokenExpiration),
		ServerIP:           serverIP,
	}

	// Save the session to the database
	if err := tx.Create(session).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		fileLogger.LogToFile("AUTH", fmt.Sprintf("Failed to save session to database. Error: %s", err.Error()))
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		fileLogger.LogToFile("AUTH", fmt.Sprintf("Failed to commit transaction. Error: %s", err.Error()))
		return
	}

	response := map[string]interface{}{
		"message": "Registered successfully.",
		"user":    newUser,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fileLogger.LogToFile("AUTH", "Error encoding JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error."})
	}
}
