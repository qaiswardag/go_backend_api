package main

import (
	"github.com/qaiswardag/go_backend_api/database"
	"github.com/qaiswardag/go_backend_api/internal/config"
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
