package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/internal/pkg/middleware"
	"github.com/qaiswardag/go_backend_api_jwt/internal/pkg/support"
	"github.com/qaiswardag/go_backend_api_jwt/pkg/httpResponseMessages"
)

type Handler struct{}

func UserObject() map[string]interface{} {
	// Response data
	return map[string]interface{}{
		"user": map[string]string{
			"firstName": "John",
			"lastName":  "Doe",
		},
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/login" && r.Method == http.MethodPost {
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

		// response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(UserObject()); err != nil {
			fmt.Printf("Error encoding JSON response: %v\n", err)
		}

		return
	}
}

func AuthValidationMiddleware(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/validate-session" {
		// Attempt to retrieve the cookie
		cookie, err := r.Cookie("session_token")

		if err != nil {
			fmt.Printf("Error: %v", cookie)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpResponseMessages.GetErrorResponse())
			return
		}

		// Check if the cookie value is empty
		if cookie.Value == "" {
			fmt.Printf("Cookie %s is empty: %+v\n", cookie.Name, cookie)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpResponseMessages.GetErrorResponse())
			return
		}

		// Compare the session token with the stored session token in the database
		if cookie.Name != "session_token" && cookie.Value != "1234" {
			fmt.Println("herrrr..")
			// response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpResponseMessages.GetErrorResponse())
		}

		// Compare the session token with the stored session token in the database
		if cookie.Name == "session_token" && cookie.Value == "1234" {

			// response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(UserObject()); err != nil {
				fmt.Printf("Error encoding JSON response: %v\n", err)
			}

			// Log the cookie name and value
			fmt.Printf("Token Name: %s, Token Value: %s\n\n", cookie.Name, cookie.Value)
			return
		}

	}
}

func Middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware1")

		next.ServeHTTP(w, r)
	})
}

func Middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware2")

		next.ServeHTTP(w, r)
	})
}

func Middleware3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware3")

		next.ServeHTTP(w, r)
	})
}

func MiddlewareMain(handler http.Handler) http.Handler {
	handler = Middleware3(handler)
	handler = Middleware2(handler)
	handler = Middleware1(handler)
	handler = middleware.Cors(handler)
	return handler
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/login", HandleLogin)
	mux.HandleFunc("/validate-session", AuthValidationMiddleware)

	handler := MiddlewareMain(mux)

	server := http.Server{
		Addr:    ":5555",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
