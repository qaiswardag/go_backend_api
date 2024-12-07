package main

import (
	"github.com/qaiswardag/go_backend_api_jwt/database"
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
}
