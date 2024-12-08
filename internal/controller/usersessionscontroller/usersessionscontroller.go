package usersessionscontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
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
	// utils.RemoveCookie(w, "session_token", true)
	// utils.RemoveCookie(w, "csrf_token", false)

	// fileLogger := logger.FileLogger{}

	// serverIP, errServerIP := utils.GetServerIP()

	// if errServerIP != nil {
	// 	log.Println("Failed to get serer ip.")
	// }

	// // Read the request body
	// var req LoginRequest
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&req); err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }

	// // Ensure the body is closed after reading
	// defer r.Body.Close()

	// db, err := database.InitDB()
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// sessionToken := tokengen.GenerateRandomToken(32)
	// utils.SetCookie(w, "session_token", sessionToken, true)
	// // Store the session_token in the database

	// csrfToken := tokengen.GenerateRandomToken(32)
	// utils.SetCookie(w, "csrf_token", csrfToken, false)

	// // Hash the password using bcrypt
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	http.Error(w, "Error hashing password", http.StatusInternalServerError)
	// 	return
	// }

	// // Compare the hashed password from the user input with the hashed password stored in the database

	// // Create a Session object
	// session := &model.Session{
	// 	UserID:            int(sessionUser.ID),
	// 	AccessToken:       sessionToken,
	// 	ServerIP:          serverIP,
	// 	AccessTokenExpiry: time.Now().Add(appconstants.TokenExpiration),
	// }

	// fileLogger.LogToFile("AUTH", fmt.Sprintf("AUTH", "Auth: Successfully logged in"))

}

// Handler update password
func Update(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(model.UserObject()); err != nil {
		log.Printf("Error encoding JSON response: %v\n", err)
	}
}
