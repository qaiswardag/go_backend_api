package main

import (
	"fmt"

	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/model"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Create tables
	database.DropTables(db)

	// Drop all tables
	database.CreateTables(db)

	// Create 10 fake users
	for i := 1; i <= 10; i++ {
		user := model.User{
			UserName:  fmt.Sprintf("user%d", i),
			FirstName: fmt.Sprintf("FirstName%d", i),
			LastName:  fmt.Sprintf("LastName%d", i),
			Email:     fmt.Sprintf("user%d@example.com", i),
		}
		db.Create(&user)
	}
	// Create 20 fake jobs
	for i := 1; i <= 20; i++ {
		job := model.Job{Title: fmt.Sprintf("job%d", i)}
		db.Create(&job)
	}
}
