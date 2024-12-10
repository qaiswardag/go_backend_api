package utils

import (
	"net/http"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/internal/appconstants"
)

// RemoveCookie removes a cookie by setting its expiration date to a time in the past.
func RemoveCookie(w http.ResponseWriter, name string, httpOnly bool) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: httpOnly,
		Path:     "/",
	})
}

// SetCookie sets a cookie with the given parameters.
func SetCookie(w http.ResponseWriter, name, value string, httpOnly bool) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(appconstants.TokenExpiration),
		HttpOnly: httpOnly,
		Path:     "/",
	})
}
