package api

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"userservice/internal/data/repository"
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

	user, err := repository.GetUserWithEmail(req.Email)
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

		err = repository.IncrementLoginAttemptCount(&user)
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

	err = repository.ResetLoginAttemptCount(&user)
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

	return c.Status(fiber.StatusOK).JSON(resp)
}
