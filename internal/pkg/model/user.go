package model

func UserObject() map[string]interface{} {
	// Response data
	return map[string]interface{}{
		"user": map[string]string{
			"firstName": "John",
			"lastName":  "Doe",
		},
	}
}
