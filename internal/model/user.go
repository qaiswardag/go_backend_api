package model

func UserObject() map[string]interface{} {
	return map[string]interface{}{
		"user": map[string]string{
			"firstName": "John",
			"lastName":  "Doe",
		},
	}
}
