package qrouter

import (
	"fmt"
	"net/http"
)

func NewRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request URL:", r.URL)
	fmt.Println("Request Method:", r.Method)
}
