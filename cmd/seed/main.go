package main

import (
	"fmt"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"gorm.io/gorm"
)

type User struct {
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
	db.Exec("DROP TABLE IF EXISTS users, jobs")

	// AutoMigrate will create the Job and Product tables
	db.AutoMigrate(&User{}, &Job{})

	// Create 10 fake users
	for i := 1; i <= 10; i++ {
		user := User{UserName: fmt.Sprintf("user%d", i)}
		db.Create(&user)
	}

	// Create 20 fake jobs
	for i := 1; i <= 20; i++ {
		job := Job{Title: fmt.Sprintf("job%d", i)}
		db.Create(&job)
	}
}
