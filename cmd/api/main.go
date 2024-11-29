package main

import (
	"fmt"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/routes"
)

func main() {

	handler := routes.SetupRoutes()

	server := http.Server{
		Addr:    ":5555",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
