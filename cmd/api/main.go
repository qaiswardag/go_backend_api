package main

import (
	"log"
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/config"
	"github.com/qaiswardag/go_backend_api_jwt/internal/routes"
)

func main() {
	// Load environment variables file
	config.LoadEnvironmentFile()

	// serverAddr:
	// Go Server Ports Behind a Web Serve:
	// Do not directly expose the Go port (e.g., localhost:8080) to the public internet.

	// Setting Up the Server to Forward Requests to Go Application:
	// Hosting company: configure the server to forwards public requests from (e.g., www.myissue.digitalocean.com) to (e.g., localhost:8080).
	// Ensure that the Go application is running on the specific internal port (e.g., localhost:8080) and listens for these requests.

	// Managing High Traffic with Go:
	// With millions of requests, distributing the load ensures that no single instance is overwhelmed, improving performance and reliability.
	// Ensure that the Go application is deployed across multiple instances, and
	// use load balancing to distribute requests evenly across those instances.
	serverAddr := config.GetEnvironmentVariable("SERVER_ADDR")

	handler := routes.MainRouter()

	server := http.Server{
		Addr:    serverAddr,
		Handler: handler,
	}

	errServer := server.ListenAndServe()

	if errServer != nil {
		log.Fatalf("Server failed to start: %v", errServer)

	}
}
