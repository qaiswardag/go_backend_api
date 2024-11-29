package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetCORSOrigin(key string) string {
	fmt.Println("den er:", os.Getenv(key))
	return os.Getenv(key)
}
