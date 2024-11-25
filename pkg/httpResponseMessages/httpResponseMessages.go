package httpResponseMessages

type Messages struct {
	ErrorMessage   string `json:"error_message"`
	ErrorNotFound  string `json:"error_not_found"`
	SuccessMessage string `json:"success_message"`
}

func (m *Messages) SetSuccessMessage(message string) {
	m.SuccessMessage = message
}

func (m *Messages) SetErrorMessage(message string) {
	m.ErrorMessage = message
}

func (m *Messages) SetErrorNotFoundMessage(message string) {
	m.ErrorNotFound = message
}

func GetSuccessResponse() Messages {
	m := Messages{}
	m.SetSuccessMessage("Everything good..")
	return m
}

func GetErrorResponse() Messages {
	m := Messages{}
	m.SetErrorMessage("This is an error message..")
	return m
}

func GetErrorNotFoundMessage() Messages {
	m := Messages{}
	m.SetErrorNotFoundMessage("This is an error message. Page not found..")
	return m
}
