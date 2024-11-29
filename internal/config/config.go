package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}
