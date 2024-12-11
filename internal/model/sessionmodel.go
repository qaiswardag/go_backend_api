package model

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID             int       `gorm:"not null"`
	SessionToken       string    `gorm:"not null"`
	SessionTokenExpiry time.Time `gorm:"not null"`
	ServerIP           string    `gorm:"not null"`
}
