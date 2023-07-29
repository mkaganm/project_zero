package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"userservice/internal/data/repository"
)

type SendVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// SendVerificationCode service
func SendVerificationCode(c *fiber.Ctx) error {

	request := c.Request()

	var req SendVerificationRequest
	err := json.Unmarshal(request.Body(), &req)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "invalid request body",
		}

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "invalid request body",
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

	err = SendMailCode(user)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "error while sending verification code",
		}

		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := Response{
		Status: "success",
		Data:   map[string]interface{}{},
	}

	resp.Data["user"] = UserResponse{
		UserId:            user.Id,
		Username:          user.Username,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		IsBlocked:         user.IsBlocked,
		IsVerified:        user.IsVerified,
		LoginAttemptCount: user.LoginAttemptCount,
		UpdatedAt:         user.UpdatedAt,
		CreatedAt:         user.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
