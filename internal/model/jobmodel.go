package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	UserID      int    `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
}
