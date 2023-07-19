package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckHashPassword checks if a password matches a hash
func CheckHashPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateVerificationCode is a function that generates a verification code
func GenerateVerificationCode() string {

	rand.Seed(time.Now().UnixNano())

	code := ""
	for i := 0; i < 6; i++ {

		num := rand.Intn(10)
		code += fmt.Sprintf("%d", num)
	}

	return code
}
