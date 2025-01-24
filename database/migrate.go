package database

import (
	"github.com/qaiswardag/go_backend_api/internal/model"
	"gorm.io/gorm"
)

// Drop all tables
func DropTables(db *gorm.DB) {
	db.Exec("DROP TABLE IF EXISTS users, jobs, sessions")
}

// Create tables
func CreateTables(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Job{}, &model.Session{})
}
