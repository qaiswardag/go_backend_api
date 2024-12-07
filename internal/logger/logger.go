package logger

import (
	"log"
	"os"
	"time"
)

// LogToFile logs a message to the specified log file
func LogToFile(title string, message string) {

	// Check if the directory exists
	storageDir := "storage"

	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		// Create the "storage" directory
		err := os.Mkdir(storageDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create storage directory: %v", err)
		}
	}

	// Get the current date and time
	currentTime := time.Now().Format("2006-01-02_15-04-05")

	// Create the log file name with the current date and time
	fileName := "storage/logger_" + currentTime + ".log"

	// Open the log file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open logger file: %v", err)
	}
	defer file.Close()

	// Create a file logger and set the output to the file
	fileLogger := log.New(file, "", log.LstdFlags)

	fileLogger.Println(currentTime + " " + title + ": " + message + "\n")
}
