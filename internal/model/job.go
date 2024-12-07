package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
}
