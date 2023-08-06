package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
	"userservice/internal/config"
	"userservice/internal/data/postgreDB"
	"userservice/internal/data/redisDB"
	"userservice/internal/utils"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login service
func Login(c *fiber.Ctx) error {

	request := c.Request()

	var req loginRequest

	err := json.Unmarshal(request.Body(), &req)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errMsg,
		}

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errMsg,
		}

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	user, err := postgreDB.GetUserWithEmail(req.Email)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "user not found",
		}

		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	if user.IsBlocked {

		err = errors.New("user is blocked")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: err.Error(),
		}

		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	isValidPassword := utils.CheckHashPassword(req.Password, user.Password)
	if !isValidPassword {

		err = postgreDB.IncrementLoginAttemptCount(&user)
		utils.LogErr("Error while increment login attempt count", err)

		err = errors.New("invalid password")
		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "invalid password",
		}

		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	if !user.IsVerified {

		respUser := UserResponse{
			UserId:            user.Id,
			Username:          user.Username,
			FirstName:         user.FirstName,
			LastName:          user.LastName,
			Email:             user.Email,
			PhoneNumber:       user.PhoneNumber,
			IsBlocked:         user.IsBlocked,
			LoginAttemptCount: user.LoginAttemptCount,
			IsVerified:        user.IsVerified,
			CreatedAt:         user.CreatedAt,
			UpdatedAt:         user.UpdatedAt,
		}

		resp := Response{
			Status: "failed",
			Data:   make(map[string]interface{}),
		}
		resp.Data["user"] = respUser

		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	err = postgreDB.ResetLoginAttemptCount(&user)
	utils.LogErr("Error while reset login attempt count", err)

	resp := Response{
		Status: "success",
		Data:   make(map[string]interface{}),
	}

	resp.Data["user"] = UserResponse{
		UserId:            user.Id,
		Username:          user.Username,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		IsBlocked:         user.IsBlocked,
		LoginAttemptCount: user.LoginAttemptCount,
		IsVerified:        user.IsVerified,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}

	// Generate token and give it to user as cookie for authentication
	token, _ := utils.GenerateToken(config.EnvConfigs.Secret)
	cookieData := redisDB.CookieData{
		UserId:       user.Id,
		SessionToken: token,
		Key:          fmt.Sprintf("cookie:%s", uuid.New().String()),
		Timestamp:    time.Now(),
	}
	// Insert cookie data to redis
	err = redisDB.InsertCookieData(&cookieData)
	if err != nil {
		utils.LogErr("Error while inserting cookie data to redis", err)

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "session not created",
		}

		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	// Convert cookie data to json
	cookieJson, _ := json.Marshal(cookieData)
	// Create cookie
	cookie := fiber.Cookie{
		Name:     "session",
		Value:    string(cookieJson),
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(resp)
}
