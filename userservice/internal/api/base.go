package api

import (
	"time"
	"userservice/internal/data/postgreDB"
	"userservice/internal/producer"
	"userservice/internal/utils"
)

const errMsg = "invalid request body"
const unAuthMsg = "unauthorized session"

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

// SendMailCode sends a verification code to the given email address
func SendMailCode(user postgreDB.User) (err error) {

	err = postgreDB.DeleteVerificationWithUserId(user.Id)
	utils.LogErr("delete verification code", err)

	code := utils.GenerateVerificationCode()
	cadeHash, _ := utils.HashPassword(code)

	verificationTable := postgreDB.Verification{
		UserID:               user.Id,
		VerificationCodeHash: cadeHash,
	}

	_, err = postgreDB.InsertVerificationCode(verificationTable)
	if err != nil {
		utils.LogErr("Error when insert verification code to db.", err)
		return err
	}

	producer.PublishMailerMessage(producer.MailMessage{
		To:      []string{user.Email},
		Subject: "PROJECT_ZERO : Verification Code",
		Body:    "Your verification code is: " + code + "\n\n PROJECT_ZERO",
	})

	return err
}
