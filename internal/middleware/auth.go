package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
)

func RequireSessionMiddleware(next http.Handler) http.Handler {
	fileLogger := logger.FileLogger{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Attempt to retrieve the cookie
		cookie, err := r.Cookie("session_token")

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

		db, err := database.InitDB()
		if err != nil {
			panic("failed to connect database")
		}

		// Retrieve the access token from database
		var sessionUser model.Session
		if err := db.Where("access_token = ?", cookie.Value).First(&sessionUser).Error; err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Access Token not found."})
			fileLogger.LogToFile("AUTH", "Access Token not found.")
			return
		}

		// Compare the session token with the stored session token in the database
		if cookie.Name != "session_token" || cookie.Value != sessionUser.AccessToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User authorization failed. The session cookie does not match the stored session token."})
			fileLogger.LogToFile("AUTH", "User authorization failed. Session Cookie does not match the stored session token.")
			return
		}

		// Retrieve the user information from the database
		var user model.User
		if err := db.Where("id = ?", sessionUser.UserID).First(&user).Error; err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "User not found."})
			fileLogger.LogToFile("AUTH", "User not found.")
			return
		}

		// If the session token matches, authorization is successful
		ctx := context.WithValue(r.Context(), "sessionUserKey", sessionUser)
		ctx = context.WithValue(ctx, "userKey", user)
		r = r.WithContext(ctx)

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "Successfully been authenticated."}); err != nil {
			fileLogger.LogToFile("AUTH", "Error encoding JSON response")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		}

		fileLogger.LogToFile("AUTH", "Middleware auth. Successfully been authenticated.")

		// Pass control to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
