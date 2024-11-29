package middleware

import (
	"fmt"
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CORS before preflight")
		// This is important for enabling cross-origin requests, especially from a frontend on a different domain
		// Allows requests from the specified origin (localhost:7777) to access the resource
		// Only requests coming from http://localhost:7777 are allowed to access the backend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:7777")

		// Set to true means that the frontend is allowed to send cookies (or session tokens)
		// If false, the frontend will not send any cookies or authorization headers when making requests to the backend
		// Specifies whether the browser should include credentials (cookies, HTTP authentication, etc.) in the request
		//This is required if the server needs authentication (e.g., via cookies or session tokens)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Specifies which HTTP methods (GET, POST, PUT, DELETE, OPTIONS) the client is allowed to use for the request
		// This is typically part of the preflight response to tell the client what methods are supported by the server
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Specifies which headers can be included in the actual request
		// For example, `Authorization` header is included here, which tells the frontend it is allowed to send an Authorization token
		// The frontend can send an Authorization token in the header without being blocked by the CORS policy
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept-Version")

		// Log the request method and URL path

		// Handle preflight request
		// GET requests don't trigger a preflight OPTIONS request, so the handler is called only once
		// Post requests first trigger a preflight OPTIONS request, so the handler is called only twice
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		fmt.Println("CORS after preflight")
		next.ServeHTTP(w, r)
	})
}
