package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

// CheckPassword Custom validator for password
func CheckPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check if password has at least 8 characters
	if len(password) < 8 {
		return false
	}

	// Check if password has at least one uppercase letter
	hasUppercase := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
			break
		}
	}
	if !hasUppercase {
		return false
	}

	// Check if password has at least one lowercase letter
	hasLowercase := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLowercase = true
			break
		}
	}
	if !hasLowercase {
		return false
	}

	// Check if password has at least one digit
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return false
	}

	// Check if password has at least one special character
	hasSpecial := false
	specialChars := `!@#$%^&*()-_=+{}[]|<>,.?/~`
	for _, char := range password {
		if strings.ContainsRune(specialChars, char) {
			hasSpecial = true
			break
		}
	}
	if !hasSpecial {
		return false
	}

	return true
}
