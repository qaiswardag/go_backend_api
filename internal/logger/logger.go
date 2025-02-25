package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// FileLogger implements the Logger interface
type FileLogger struct{}

// LogToFile logs a message to the specified log file
func (f FileLogger) LogToFile(title string, message string) {

	// Check if the "storage" directory exists
	storageDir := "storage"
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		// Create the "storage" directory
		err := os.Mkdir(storageDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create storage directory: %v", err)
		}
	}

	// Check if the "logger" directory exists inside "storage"
	loggerDir := "storage/logger"
	if _, err := os.Stat(loggerDir); os.IsNotExist(err) {
		// Create the "logger" directory
		err := os.Mkdir(loggerDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create logger directory: %v", err)
		}
	}

	// Get the current date and time
	currentTime := time.Now().Format("2006-01-02_15-04-05")

	// Create the log file name with the current date and time
	fileName := "storage/logger/logger_" + currentTime + ".log"

	// Open the log file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open logger file: %v", err)
	}
	defer file.Close()

	// Create a file logger and set the output to the file
	fileLogger := log.New(file, "", log.LstdFlags)

	// Format the log entry
	logEntry := fmt.Sprintf("%-20s %s", title+":", message)

	// Log the formatted entry
	fileLogger.Println(logEntry)

}
