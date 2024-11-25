package httpResponseMessages

type Messages struct {
	ErrorMessage   string `json:"error_message"`
	SuccessMessage string `json:"success_message"`
}

func (m *Messages) SetErrorMessage(message string) {
	m.ErrorMessage = message
}
func (m *Messages) SetSuccessMessage(message string) {
	m.SuccessMessage = message
}

func GetErrorResponse() string {
	m := Messages{}
	m.SetErrorMessage("This is an error message..")
	return m.ErrorMessage
}

func GetSuccessResponse() string {
	m := Messages{}
	m.SetSuccessMessage("Everything good..")
	return m.SuccessMessage
}
