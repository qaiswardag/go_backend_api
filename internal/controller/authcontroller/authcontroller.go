package authcontroller

import (
	"encoding/json"
	"log"
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

	fileLogger.LogToFile("AUTH", "Get auth.")
	if err := json.NewEncoder(w).Encode(model.UserObject()); err != nil {
		log.Printf("Error encoding JSON response: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
}
