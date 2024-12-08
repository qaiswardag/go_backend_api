package database

import (
	"fmt"

	"github.com/qaiswardag/go_backend_api_jwt/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	DBName := config.GetEnvironmentVariable("DB_DATABASE")
	DBUsername := config.GetEnvironmentVariable("DB_USERNAME")
	DBPassword := config.GetEnvironmentVariable("DB_PASSWORD")

	if db != nil {
		return db, nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUsername, DBPassword, DBName)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
