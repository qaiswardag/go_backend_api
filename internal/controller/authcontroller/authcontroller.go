package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/database"
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

func Show(w http.ResponseWriter, r *http.Request) {
	fileLogger := logger.FileLogger{}

	sessionUser, _ := r.Context().Value("sessionUserKey").(model.Session)
	user, _ := r.Context().Value("userKey").(model.User)

	response := map[string]interface{}{
		"sessionUser": sessionUser,
		"user":        user,
	}

	w.WriteHeader(http.StatusOK)
	fileLogger.LogToFile("AUTH", "Successfully found user and sent response.")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		fileLogger.LogToFile("AUTH", "Error encoding JSON response")
		return
	}

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// Log user and sessionUser information with field names
	// userJSON, err := json.MarshalIndent(user, "", "  ")
	// if err != nil {
	// 	fileLogger.LogToFile("USER", "Error marshalling user to JSON")
	// } else {
	// 	fileLogger.LogToFile("USER", fmt.Sprintf("User is: %s", userJSON))
	// }

	// sessionUserJSON, err := json.MarshalIndent(sessionUser, "", "  ")
	// if err != nil {
	// 	fileLogger.LogToFile("USER", "Error marshalling sessionUser to JSON")
	// } else {
	// 	fileLogger.LogToFile("USER", fmt.Sprintf("Session User is: %s", sessionUserJSON))
	// }
}

func Update(w http.ResponseWriter, r *http.Request) {
	fileLogger := logger.FileLogger{}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"message": "ØØØØØJ. IKKKKKKKKE HEEEER."})
	fileLogger.LogToFile("AUTH", "ØØØØØJ. IKKKKKKKKE HEEEER.")
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	fileLogger := logger.FileLogger{}

	utils.RemoveCookie(w, "session_token", true)
	utils.RemoveCookie(w, "csrf_token", false)

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Retrieve the user from the context
	sessionUser, okSessionUser := r.Context().Value("sessionUserKey").(model.Session)

	if !okSessionUser {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to retrieve session user from context"})
		fileLogger.LogToFile("AUTH", "Failed to retrieve session user from context")
		return
	}

	// Delete all other sessions that match the UserID and ServerIP
	if err := db.Exec("DELETE FROM sessions WHERE user_id = ? AND server_ip = ?", sessionUser.UserID, sessionUser.ServerIP).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fileLogger.LogToFile("AUTH", "Failed to delete all other sessions that match the UserID and ServerIP: "+err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error."})

}
