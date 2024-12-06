package usersessionscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
	"github.com/qaiswardag/go_backend_api_jwt/internal/security/tokengen"
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
	Password string `json:"password"`
}

func Create(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure the body is closed after reading
	defer r.Body.Close()

	// Access the username and password
	fmt.Printf("Received password: %s\n", req.Password)

	if req.Password != "1234" {
		fmt.Println("kom heeer..")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// sessionToken := tokengen.GenerateRandomToken(32)
	sessionToken := req.Password
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
		fmt.Printf("Error encoding JSON response: %v\n", err)
	}
}
