package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
	"userservice/internal/config"
)

var JWT *jwt.Token

func init() {
	initJWT()
}

func initJWT() {
	JWT = jwt.New(jwt.SigningMethodHS256)
	claims := JWT.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
}

// GenerateToken generates a token
func GenerateToken(secret string) (string, error) {
	token, err := JWT.SignedString([]byte(secret))
	return token, err
}

// CheckToken checks if a token is valid
func CheckToken(token string) bool {

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvConfigs.Secret), nil
	})

	if err != nil {
		return false
	}

	return true
}

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
