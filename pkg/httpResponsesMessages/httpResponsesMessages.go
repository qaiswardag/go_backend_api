package httpResponsesMessages

func GetErrorResponse() struct {
	Error string `json:"error"`
} {
	return struct {
		Error string `json:"error"`
	}{
		Error: "Method Not Allowed",
	}
}
