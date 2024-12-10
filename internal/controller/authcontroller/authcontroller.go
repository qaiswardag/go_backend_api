package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/appconstants"
	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
	"github.com/qaiswardag/go_backend_api_jwt/internal/utils"
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

// Get current user with basic information from auth middleware "RequireSessionMiddleware" saved in the context
func Show(w http.ResponseWriter, r *http.Request) {
	fileLogger := logger.FileLogger{}

	authenticatedUser, okAuthenticatedUser := r.Context().Value(appconstants.ContextKeyAuthenticatedUser).(model.User)

	// Check if the session user is not found in the context
	if !okAuthenticatedUser {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to retrieve authenticated user from context."})
		fileLogger.LogToFile("AUTH", "Failed to retrieve authenticated user from context.")
		return
	}

	response := map[string]interface{}{
		"user": authenticatedUser,
	}

	w.WriteHeader(http.StatusOK)
	fileLogger.LogToFile("AUTH", "Successfully retrieved the authenticated user from the context.")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		fileLogger.LogToFile("AUTH", "Error encoding JSON response")
		return
	}
}

// Destroy the session "sign out the user" by deleting the session from the databaseand
// and remove the session token and CSRF token from the client
func Destroy(w http.ResponseWriter, r *http.Request) {
	fileLogger := logger.FileLogger{}

	utils.RemoveCookie(w, "session_token", true)
	utils.RemoveCookie(w, "csrf_token", false)

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database.")
	}

	// Retrieve the session user from the context
	authenticatedSession, okAuthenticatedSession := r.Context().Value(appconstants.ContextKeyAuthenticatedSession).(model.Session)

	// Check if the session user is not found in the context
	if !okAuthenticatedSession {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to retrieve session user from context."})
		fileLogger.LogToFile("AUTH", "Failed to retrieve session user from context.")
		return
	}

	// Delete the session from the database
	if err := db.Exec("DELETE FROM sessions WHERE user_id = ? AND server_ip = ?", authenticatedSession.UserID, authenticatedSession.ServerIP).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fileLogger.LogToFile("AUTH", "Failed to delete all other sessions that match the UserID and ServerIP: "+err.Error())
	}

	// Send a response to the client with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Successfully signed out."})

}
