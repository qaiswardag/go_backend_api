package middleware

import (
	"net/http"
)

func GlobalMiddleware(handler http.Handler) http.Handler {
	handler = RequireSessionMiddleware(handler)
	handler = Cors(handler)
	return handler
}
