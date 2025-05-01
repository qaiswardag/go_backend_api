package jsonhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ReadJSON struct{}

// Define a struct that matches your JSON body
type RequestBody struct {
	Email string `json:"email"`
}

func (R *ReadJSON) ValidateAndParseJSON(r *http.Request) {

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	// fmt.Println("isss now", string(bodyBytes))

	//  Unmarshal into a struct
	reqBody := RequestBody{}

	if err := json.Unmarshal(bodyBytes, &reqBody); err != nil {
		fmt.Println("error unmarshaling body:", err)
		return
	}

	fmt.Println("Email is:", reqBody.Email)
}
