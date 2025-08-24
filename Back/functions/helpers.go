package functions

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var err error

// Hash password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Check password
func CheckHash(password string, database_password string) bool {
	err = bcrypt.CompareHashAndPassword([]byte(database_password), []byte(password))
	return err == nil
}

// Generate UUID
func GenerateUUID() string {
	return uuid.New().String()
}
