package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/internal/pkg/support"
	"github.com/qaiswardag/go_backend_api_jwt/pkg/httpResponsesMessages"
)

type Handler struct{}

// func authorize(r *http.Request) {
// 	st, err := r.Cookie("session_token")

// 	if err != nil {
// 		fmt.Println("No session token.")
// 	}
// 	fmt.Println("Session token is:", st)
// }

func login(r *http.Request, w http.ResponseWriter) {
	if r.URL.Path == "/løgin" && r.Method == http.MethodPost {
		fmt.Println("came to login")
		sessionToken := support.GenerateToken(32)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: false,
		})
		// Store session_token token in database

		csrfToken := support.GenerateToken(32)
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: false,
		})
		// Store csrf_token token in database

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(httpResponsesMessages.GetSuccessResponse())
		return
	}
	fmt.Println("did not come to login")
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept-Version")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	login(r, w)

	// Handle other methods
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(httpResponsesMessages.GetErrorResponse())
}

func main() {
	handler := Handler{}

	server := http.Server{
		Addr:    ":7070",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
