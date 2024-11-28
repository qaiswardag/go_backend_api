package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func HandleSession(w http.ResponseWriter, r *http.Request) {
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

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is important for enabling cross-origin requests, especially from a frontend on a different domain
		// Allows requests from the specified origin (localhost:7777) to access the resource
		// Only requests coming from http://localhost:7777 are allowed to access the backend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:7777")

		// Set to true means that the frontend is allowed to send cookies (or session tokens)
		// If false, the frontend will not send any cookies or authorization headers when making requests to the backend
		// Specifies whether the browser should include credentials (cookies, HTTP authentication, etc.) in the request
		//This is required if the server needs authentication (e.g., via cookies or session tokens)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Specifies which HTTP methods (GET, POST, PUT, DELETE, OPTIONS) the client is allowed to use for the request
		// This is typically part of the preflight response to tell the client what methods are supported by the server
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Specifies which headers can be included in the actual request
		// For example, `Authorization` header is included here, which tells the frontend it is allowed to send an Authorization token
		// The frontend can send an Authorization token in the header without being blocked by the CORS policy
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept-Version")

		// Log the request method and URL path

		// Handle preflight request
		// GET requests don't trigger a preflight OPTIONS request, so the handler is called only once
		// Post requests first trigger a preflight OPTIONS request, so the handler is called only twice
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/login", HandleLogin)
	mux.HandleFunc("/validate-session", HandleSession)

	handler := Middleware(mux)

	server := http.Server{
		Addr:    ":5555",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
