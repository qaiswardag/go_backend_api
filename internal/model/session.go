package model

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID            int       `gorm:"not null"`
	AccessToken       string    `gorm:"not null"`
	ServerIP          string    `gorm:"not null"`
	AccessTokenExpiry time.Time `gorm:"not null"`
}
