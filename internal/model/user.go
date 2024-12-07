package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string `gorm:"unique"`
	Email         string `gorm:"unique;not null"`
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	LoginAttempts *int
}

func UserObject() map[string]interface{} {
	return map[string]interface{}{
		"user": map[string]string{
			"firstName": "John",
			"lastName":  "Doe",
		},
	}
}
