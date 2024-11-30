package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/pkg/httpresp"
)

func RequireSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Attempt to retrieve the cookie
		cookie, err := r.Cookie("session_token")

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpresp.GetErrorResponse())
			fmt.Println("User authorization failed")
			return
		}

		// Check if the cookie value is empty
		if cookie.Value == "" {
			fmt.Printf("Cookie %s is empty: %+v\n", cookie.Name, cookie)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpresp.GetErrorResponse())
			fmt.Println("User authorization failed")
			return
		}

		// Compare the session token with the stored session token in the database
		if cookie.Name != "session_token" || cookie.Value != "1234" {
			// response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpresp.GetErrorResponse())
			fmt.Println("User authorization failed")
		}

		// Compare the session token with the stored session token in the database
		if cookie.Name == "session_token" && cookie.Value == "1234" {

			// response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			// Log the cookie name and value
			fmt.Printf("Token Name: %s, Token Value: %s\n\n", cookie.Name, cookie.Value)
			return
		}

		next.ServeHTTP(w, r)
	})
}
