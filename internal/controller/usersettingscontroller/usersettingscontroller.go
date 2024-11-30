package usersettingscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	w.WriteHeader(http.StatusOK)
	fmt.Println("Hit user settings route!")
	if err := json.NewEncoder(w).Encode(model.UserObject()); err != nil {
		fmt.Printf("Error encoding JSON response: %v\n", err)
	}
}
