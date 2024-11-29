package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/internal/pkg/middleware"
	"github.com/qaiswardag/go_backend_api_jwt/internal/pkg/model"
	"github.com/qaiswardag/go_backend_api_jwt/internal/pkg/support"
	"github.com/qaiswardag/go_backend_api_jwt/pkg/httpResponseMessages"
)

type Handler struct{}

func HandleLoginRoute(w http.ResponseWriter, r *http.Request) {
	// sessionToken := support.GenerateToken(32)
	sessionToken := "1234"
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})
	// Store the session token in the database.

	csrfToken := support.GenerateToken(32)
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})
	// Store csrf_token token in database

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(model.UserObject()); err != nil {
		fmt.Printf("Error encoding JSON response: %v\n", err)
	}
}

func HandleUserSettingsRoute(w http.ResponseWriter, r *http.Request) {
	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(model.UserObject()); err != nil {
		fmt.Printf("Error encoding JSON response: %v\n", err)
	}
}

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	handler := middleware.MiddlewareMain(mux)

	// Main route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(httpResponseMessages.GetErrorNotFoundMessage()); err != nil {
			fmt.Printf("Error encoding JSON response: %v\n", err)
		}
	})

	mux.HandleFunc("/login", HandleLoginRoute)
	mux.HandleFunc("/user/settings", HandleUserSettingsRoute)

	return handler
}
