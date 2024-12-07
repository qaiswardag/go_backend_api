package main

import "github.com/qaiswardag/go_backend_api_jwt/database"

type Job struct {
	Code   string
	Price  uint
	New    uint
	NewNew uint
}

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Job{})
}
