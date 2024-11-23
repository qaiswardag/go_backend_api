package httpResponsesMessages

type Messages struct {
	ErrorMessage string `json:"error_message"`
}

func (m *Messages) SetErrorMessage(message string) {
	m.ErrorMessage = message
}

func GetErrorResponse() Messages {
	m := Messages{}
	m.SetErrorMessage("This is an error message")
	return m
}

func GetSuccessResponse() Messages {
	m := Messages{}
	m.SetErrorMessage("Everything good")
	return m
}
