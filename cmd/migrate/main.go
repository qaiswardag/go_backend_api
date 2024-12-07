package main

import (
	"github.com/qaiswardag/go_backend_api_jwt/database"
	"github.com/qaiswardag/go_backend_api_jwt/internal/config"
)

func main() {

	// Load environment variables file
	config.LoadEnvironmentFile()

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Create tables
	database.DropTables(db)

	// Drop all tables
	database.CreateTables(db)
}
