package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/appconstants"
	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
)

// RequireSessionMiddleware is a middleware that checks if the user is authenticated
func RequireSessionMiddleware(next http.Handler) http.Handler {
	fileLogger := logger.FileLogger{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Retrieve the session token from the cookie
		cookie, err := r.Cookie("session_token")

		// Check if the cookie is not found
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Session Cookis is empty."})
			fileLogger.LogToFile("AUTH", "Session Cookis is empty.")
			return
		}

		// Check if the cookie value is empty
		if cookie.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Session Cookie is empty"})
			fileLogger.LogToFile("AUTH", "User authorization failed. Session Cookie is empty.")
			return
		}

		// Connect to the database
		db, err := database.InitDB()
		if err != nil {
			panic("failed to connect database")
		}

		// Retrieve the session user from the database
		authenticatedSession := model.Session{}
		if err := db.Where("access_token = ?", cookie.Value).First(&authenticatedSession).Error; err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Access Token not found."})
			fileLogger.LogToFile("AUTH", "Access Token not found.")
			return
		}

		// Check if the session token matches the stored session token
		if cookie.Name != "session_token" || cookie.Value != authenticatedSession.AccessToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User authorization failed. The session cookie does not match the stored session token."})
			fileLogger.LogToFile("AUTH", "User authorization failed. Session Cookie does not match the stored session token.")
			return
		}

		// Retrieve the user from the database
		authenticatedUser := model.User{}
		if err := db.Where("id = ?", authenticatedSession.UserID).First(&authenticatedUser).Error; err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User not found."})
			fileLogger.LogToFile("AUTH", "User not found.")
			return
		}

		// Save
		ctx := context.WithValue(r.Context(), appconstants.ContextKeyAuthenticatedSession, authenticatedSession)
		ctx = context.WithValue(ctx, appconstants.ContextKeyAuthenticatedUser, authenticatedUser)

		r = r.WithContext(ctx)

		fileLogger.LogToFile("AUTH", "Middleware auth. Successfully been authenticated.")

		// Pass control to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
