package usersessionscontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/appconstants"
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handler login
func Create(w http.ResponseWriter, r *http.Request) {
	utils.RemoveCookie(w, "session_token", true)
	utils.RemoveCookie(w, "csrf_token", false)

	fileLogger := logger.FileLogger{}

	serverIP, errServerIP := utils.GetServerIP()

	if errServerIP != nil {
		log.Println("Failed to get serer ip.")
	}

	// Read the request body
	var req LoginRequest
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

	// Retrieve the user from the database
	var user model.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found."})
		return
	}

	// Compare the hashed password from the user input with the hashed password stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Password is incorrect."})
		return
	}

	sessionToken := tokengen.GenerateRandomToken(32)
	utils.SetCookie(w, "session_token", sessionToken, false)
	// Store the session_token in the database

	csrfToken := tokengen.GenerateRandomToken(32)
	utils.SetCookie(w, "csrf_token", csrfToken, false)

	// Create a Session object
	session := &model.Session{
		UserID:            int(user.ID),
		AccessToken:       sessionToken,
		ServerIP:          serverIP,
		AccessTokenExpiry: time.Now().Add(appconstants.TokenExpiration),
	}

	// Delete all other sessions that match the UserID and ServerIP
	if err := db.Exec("DELETE FROM sessions WHERE user_id = ? AND server_ip = ?", session.UserID, session.ServerIP).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fileLogger.LogToFile("AUTH", "Failed to delete all other sessions that match the UserID and ServerIP: "+err.Error())
	}

	// Save the session to the database
	if err := db.Create(&session).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to create the session record in the database."})

		utils.RemoveCookie(w, "session_token", true)
		utils.RemoveCookie(w, "csrf_token", false)
		return
	}

	response := map[string]interface{}{
		"message": "Successfully signed in.",
		"user":    user,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fileLogger.LogToFile("AUTH", "Error encoding JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error."})
	}
}
