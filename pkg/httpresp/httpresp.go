package httpresp

type Messages struct {
	ErrorMessage   string `json:"error_message,omitempty"`
	ErrorNotFound  string `json:"error_not_found,omitempty"`
	SuccessMessage string `json:"success_message,omitempty"`
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
	return Messages{
		SuccessMessage: "Everything good..",
	}
}

func GetErrorResponse() Messages {
	return Messages{
		ErrorMessage: "This is an error message..",
	}
}

func GetErrorNotFoundMessage() Messages {
	return Messages{
		ErrorNotFound: "This is an error message. Page not found..",
	}
}
