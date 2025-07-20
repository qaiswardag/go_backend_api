package router

import (
	"fmt"
	"net/http"
)

func SayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helooooooooo")
}
