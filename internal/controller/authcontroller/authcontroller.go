package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/logger"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
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

	// Retrieve the session user from the context
	sessionUser, ok := r.Context().Value("sessionUserKey").(model.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to retrieve session user from context"})
		fileLogger.LogToFile("AUTH", "Failed to retrieve session user from context")
		return
	}

	// Retrieve the user from the context
	user, ok := r.Context().Value("userKey").(model.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to retrieve user from context"})
		fileLogger.LogToFile("AUTH", "Failed to retrieve user from context")
		return
	}

	// Use the user information as needed
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		fileLogger.LogToFile("AUTH", "Error encoding JSON response")
		return
	}
	// Use the user information as needed
	if err := json.NewEncoder(w).Encode(sessionUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
		fileLogger.LogToFile("AUTH", "Error encoding JSON response")
		return
	}

	// // Log user and sessionUser information with field names
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

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Successfully been authenticated."})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registered successfully"})
	fileLogger.LogToFile("AUTH", "Successfully been authenticated.")

}
