package qrouter

import (
	"fmt"
	"net/http"
)

func AddRoute() {
	fmt.Println("add route ran..")
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request URL:", r.URL)
	fmt.Println("Request Method:", r.Method)
	AddRoute()
}
