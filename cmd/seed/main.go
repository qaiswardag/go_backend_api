package main

import (
	"fmt"
	"time"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"gorm.io/gorm"
)

// `fullname` varchar(255) NOT NULL,
// `username` varchar(255) NOT NULL UNIQUE,
// `password` varchar(255) NOT NULL,
// `loginattempts` int(11) NOT NULL,
type User struct {
	gorm.Model
	UserName  string `gorm:"unique"`
	Email     string `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	// Nullable field to track the number of login attempts
	LoginAttempts *int
}
type Job struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
}
type Session struct {
	gorm.Model
	UserID            int       `gorm:"not null"`
	AccessToken       string    `gorm:"not null"`
	ServerIP          string    `gorm:"not null"`
	AccessTokenExpiry time.Time `gorm:"not null"`
}

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	// Drop all tables
	db.Exec("DROP TABLE IF EXISTS users, jobs, sessions")

	// AutoMigrate will create the Job and Product tables
	db.AutoMigrate(&User{}, &Job{}, &Session{})

	// Create 10 fake users
	for i := 1; i <= 10; i++ {
		user := User{
			UserName:  fmt.Sprintf("user%d", i),
			FirstName: fmt.Sprintf("FirstName%d", i),
			LastName:  fmt.Sprintf("LastName%d", i),
			Email:     fmt.Sprintf("user%d@example.com", i),
		}
		db.Create(&user)
	}
	// Create 20 fake jobs
	for i := 1; i <= 20; i++ {
		job := Job{Title: fmt.Sprintf("job%d", i)}
		db.Create(&job)
	}
}
