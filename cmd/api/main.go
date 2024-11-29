package main

import (
	"fmt"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/config"
	"github.com/qaiswardag/go_backend_api_jwt/internal/routes"
)

func main() {
	// Load environment variables file
	config.LoadEnv()
	serverAddr := config.GetCORSOrigin("SERVER_ADDR")

	handler := routes.SetupRoutes()

	server := http.Server{
		Addr:    serverAddr,
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
