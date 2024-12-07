package main

import (
	"github.com/qaiswardag/go_backend_api_jwt/database"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName string `gorm:"unique"`
}
type Job struct {
	gorm.Model
	Title string `gorm:"not null"`
}

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	// Drop all tables
	db.Exec("DROP TABLE IF EXISTS jobs, products")

	// AutoMigrate will create the Job and Product tables
	db.AutoMigrate(&Users{}, &Job{})
}
