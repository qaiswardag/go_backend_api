package main

import (
	"fmt"

	"github.com/qaiswardag/go_backend_api/database"
	"github.com/qaiswardag/go_backend_api/internal/config"
	"github.com/qaiswardag/go_backend_api/internal/model"
	"golang.org/x/crypto/bcrypt"
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

	var passwordUnhashed string = "123456"

	// Hash the password using bcrypt
	hashedPassword, errHashing := bcrypt.GenerateFromPassword([]byte(passwordUnhashed), bcrypt.DefaultCost)
	if errHashing != nil {
		fmt.Println("Hashing error.")
		return
	}
	// Create 10 fake users
	for i := 1; i <= 10; i++ {
		user := model.User{
			UserName: fmt.Sprintf("user%d", i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			// Convert hashed password from byte slice to string
			Password:  string(hashedPassword),
			FirstName: fmt.Sprintf("FirstName%d", i),
			LastName:  fmt.Sprintf("LastName%d", i),
		}
		db.Create(&user)
	}
	// Create 20 fake jobs
	for i := 1; i <= 20; i++ {
		job := model.Job{
			Title: fmt.Sprintf("job%d", i),
			Description: fmt.Sprintln("Consectetuer adipiscing elit. Ac per bibendum quis nec tristique porttitor. Maecenas eros maximus augue, nostra facilisi metus magna. Consequat condimentum mollis luctus molestie turpis et tortor vivamus. Elementum himenaeos potenti tempus nascetur ultrices per. Lacinia tortor eget mus felis magnis luctus. Tellus dis donec erat condimentum per nostra nibh dignissim. Purus sapien finibus mauris vivamus etiam pretium. Hac curae porttitor elementum eget lobortis lobortis. \n\nElementum non sagittis feugiat condimentum dui bibendum ultricies torquent. Sem platea bibendum blandit viverra id urna pellentesque. Phasellus tristique in sodales leo fermentum; cursus dictum. Aptent parturient mus eleifend orci ac. Amet lectus vehicula lacus ac velit. Tortor ex ipsum; fusce hac gravida sagittis porttitor. Ac mollis risus suscipit sodales libero metus magnis.",
				i),
		}
		db.Create(&job)
	}
}
