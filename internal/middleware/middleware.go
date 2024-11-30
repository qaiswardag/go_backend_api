package middleware

import (
	"net/http"
)

func GlobalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Powered-By", "Go Server")

		// Pass control to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
