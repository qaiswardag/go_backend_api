package model

import "gorm.io/gorm"

// TODO: use min 2 characters and max 255 characters for fields
type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	// Hide the password from the JSON response by omitting the field using the json:"-"`
	Password string `gorm:"not null" json:"-"`

	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	LoginAttempts *int
	Public        *bool `gorm:""`
	CurrentTeam   *int  `gorm:""`
}
