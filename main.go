package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("w:", w)
	// fmt.Println("r:", r)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept-Version")

	fmt.Println(" http.MethodOptions:", http.MethodOptions)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Welcome",
	}

	errorResponse := struct {
		Error string `json:"error"`
	}{
		Error: "Method Not Allowed",
	}

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {
	handler := Handler{}

	server := http.Server{
		Addr:    ":7070",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
