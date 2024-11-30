package homecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/pkg/httpresp"
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

func Show(w http.ResponseWriter, r *http.Request) http.Handler {
	mux := http.NewServeMux()
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(httpresp.GetErrorNotFoundMessage()); err != nil {
		fmt.Printf("Error encoding JSON response: %v\n", err)
	}

	return mux
}
