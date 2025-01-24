package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/qaiswardag/go_backend_api/internal/logger"
	"gorm.io/gorm"
)

// CheckIfRecordExists checks if a record exists in the specified table based on the given field and value.
func CheckIfRecordExists(db *gorm.DB, model interface{}, fieldName string, fieldValue interface{}, w http.ResponseWriter, logger logger.FileLogger) bool {

	err := db.Where(fmt.Sprintf("%s = ?", fieldName), fieldValue).First(model).Error

	if err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("%s already taken", fieldName)})
		logger.LogToFile("AUTH", fmt.Sprintf("%s already taken. Received value: %v", fieldName, fieldValue))
		return true
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": ""})
		logger.LogToFile("AUTH", fmt.Sprintf("Failed to query user. Error: %s", err.Error()))
		return true
	}
	return false
}
