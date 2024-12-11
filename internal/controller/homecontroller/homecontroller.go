package homecontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	fmt.Printf("heloooo%v", true)
	log.Println("Welcome to home.")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to home."})

}
