package usersessionscontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/database"
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handler login
func Create(w http.ResponseWriter, r *http.Request) {
	utils.RemoveCookie(w, "session_token", true)
	utils.RemoveCookie(w, "csrf_token", false)

	// fileLogger := logger.FileLogger{}

	// serverIP, errServerIP := utils.GetServerIP()

	// if errServerIP != nil {
	// 	log.Println("Failed to get serer ip.")
	// }

	// Read the request body
	var req LoginRequest
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

	// Create a User object
	sessionUser := model.User{
		UserName:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  string(hashedPassword),
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

}

// Handler update password
func Update(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(model.UserObject()); err != nil {
		log.Printf("Error encoding JSON response: %v\n", err)
	}
}
