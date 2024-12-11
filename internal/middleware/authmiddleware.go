package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/appconstants"
	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
	"github.com/qaiswardag/go_backend_api_jwt/internal/utils"
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
		if err := db.Where("session_token = ?", cookie.Value).First(&authenticatedSession).Error; err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Access Token not found."})
			fileLogger.LogToFile("AUTH", "Access Token not found.")
			return
		}

		// Check if the session token is not found
		if cookie.Name != "session_token" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User authorization failed. No session token in the cookie."})
			fileLogger.LogToFile("AUTH", "User authorization failed. No session token in the cookie.")
			return
		}

		// Check if the session token does not match the stored session token in the database
		if cookie.Value != authenticatedSession.SessionToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User authorization failed. The session cookie does not match the stored session token."})
			fileLogger.LogToFile("AUTH", "User authorization failed. The session cookie does not match the stored session token.")
			return
		}

		// Check if the session is older than current time
		if time.Now().After(authenticatedSession.SessionTokenExpiry) {
			utils.RemoveCookie(w, "session_token", true)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User authorization failed. The session token has expired."})
			fileLogger.LogToFile("AUTH", "User authorization failed. The session token has expired.")
			return
		}

		// Get the user from the database based on the session user ID
		authenticatedUser := model.User{}
		if err := db.Where("id = ?", authenticatedSession.UserID).First(&authenticatedUser).Error; err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User not found."})
			fileLogger.LogToFile("AUTH", "User not found.")
			return
		}

		// Extend the session token expiry by 7 days
		authenticatedSession.SessionTokenExpiry = time.Now().Add(appconstants.TokenExpiration)

		// Save the updated session to the database
		if err := db.Save(&authenticatedSession).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Failed to update the session expiry in the database."})
			fileLogger.LogToFile("AUTH", "Failed to update the session expiry in the database.")
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
