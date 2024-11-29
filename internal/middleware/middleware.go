package middleware

import (
	"net/http"
)

func GlobalMiddleware(handler http.Handler) http.Handler {
	handler = Auth(handler)
	handler = Cors(handler)
	return handler
}
