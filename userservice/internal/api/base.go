package api

import (
	"time"
	"userservice/internal/clients/mailer"
	"userservice/internal/data/entity"
	"userservice/internal/data/repository"
	"userservice/internal/utils"
)

const errMsg = "invalid request body"

type ErrorResponse struct {
	Status       string `json:"status"`
	Error        error  `json:"error"`
	ErrorMessage string `json:"error_message"`
}

type Response struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type UserResponse struct {
	UserId            uint64    `json:"user_id"`
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phone_number"`
	IsBlocked         bool      `json:"is_blocked"`
	LoginAttemptCount int       `json:"login_attempt_count"`
	IsVerified        bool      `json:"is_verified"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CookieData struct {
	UserId       uint64 `json:"user_id"`
	SessionToken string `json:"session_token"`
}

// SendMailCode sends a verification code to the given email address
func SendMailCode(user entity.User) (err error) {

	err = repository.DeleteVerificationWithUserId(user.Id)
	utils.LogErr("delete verification code", err)

	code := utils.GenerateVerificationCode()
	cadeHash, _ := utils.HashPassword(code)

	verificationTable := entity.Verification{
		UserID:               user.Id,
		VerificationCodeHash: cadeHash,
	}

	_, err = repository.InsertVerificationCode(verificationTable)
	if err != nil {
		return err
	}

	mailReq := mailer.SendMailRequest{
		To:      []string{user.Email},
		Subject: "PROJECT_ZERO : Verification Code",
		Body:    "Your verification code is: " + code + "\n\n PROJECT_ZERO",
	}

	mailer.SendMail(mailReq)
	return err
}
