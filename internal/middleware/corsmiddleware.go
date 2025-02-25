package middleware

/*
   |--------------------------------------------------------------------------
   | Cross-Origin Resource Sharing (CORS) Configuration
   |--------------------------------------------------------------------------
   |
   | Here you may configure your settings for cross-origin resource sharing
   | or "CORS". This determines what cross-origin operations may execute
   | in web browsers.
   |
   | To learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
   |
*/

import (
	"net/http"

	"github.com/qaiswardag/go_backend_api/internal/config"
	"github.com/qaiswardag/go_backend_api/internal/logger"
)

func Cors(next http.Handler) http.Handler {
	fileLogger := logger.FileLogger{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		allowedOrigins := config.GetEnvironmentVariable("CORS_ALLOW_ORIGIN")

		// This is important for enabling cross-origin requests, especially from a frontend on a different domain
		// Set the response content type to JSON with UTF-8 encoding
		// Allows requests from the specified origin (localhost:7777) to access the resource
		// Only requests coming from allowed origins have access the backend
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)

		// Cache the response for 60 seconds
		// This helps reduce server load by caching the response for a short period of time
		// When a client (like a browser) receives this response:
		// It will store the resource in its cache and use the cached version
		// for the next 60 seconds without making another request to the server.
		// After 60 seconds, the cache expires, and the client will re-fetch the resource if needed.
		// or "max-age=60" instead of "no-cache, no-store"
		w.Header().Set("Cache-Control", "no-cache, no-store")

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

		// Handle preflight request
		// GET requests don't trigger a preflight OPTIONS request, so the handler is called only once
		// POST requests first trigger a preflight OPTIONS request, so the handler is called only twice

		fileLogger.LogToFile("CORS", "Handle CORS Preflight Request before processing the request.")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Preflight request handled."}`))
			return
		}

		fileLogger.LogToFile("CORS", "Finished handling CORS Preflight Request.")

		// Pass control to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
