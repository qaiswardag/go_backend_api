package usercontroller

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

func Create(w http.ResponseWriter, r *http.Request) {
	// sessionToken := tokengen.GenerateRandomToken(32)
	sessionToken := "1234"
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
