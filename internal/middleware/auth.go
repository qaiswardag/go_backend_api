package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/pkg/httpresp"
)

func RequireSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Attempt to retrieve the cookie
		cookie, err := r.Cookie("session_token")

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpresp.GetErrorResponse())
			log.Panicf("User authorization failed: unable to retrieve session_token cookie. Error: " + err.Error())
			return
		}

		// Check if the cookie value is empty
		if cookie.Value == "" {
			log.Printf("Cookie %s is empty: %+v\n", cookie.Name, cookie)
			logger.LogToFile("AUTH", fmt.Sprintf("Authorization successful: %s: %s", cookie.Name, cookie.Value))

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpresp.GetErrorResponse())
			log.Println("User authorization failed: session_token cookie is empty.")
			return
		}

		// Compare the session token with the stored session token in the database
		if cookie.Name != "session_token" || cookie.Value != "1234" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(httpresp.GetErrorResponse())
			log.Printf("User authorization failed: session_token cookie does not match the stored session token.")
			return
		}

		// Compare the session token with the stored session token in the database
		if cookie.Name == "session_token" && cookie.Value == "1234" {
			logger.LogToFile("AUTH", fmt.Sprintf("Authorization successful: %s: %s", cookie.Name, cookie.Value))
		}

		// Pass control to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
