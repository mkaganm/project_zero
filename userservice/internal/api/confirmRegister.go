package api

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"userservice/internal/data/postgreDB"
	"userservice/internal/utils"
)

type ConfirmRegisterRequest struct {
	UserId         uint64 `json:"user_id" validate:"required"`
	ValidationCode string `json:"validation_code" validate:"required"`
}

// ConfirmRegister service
func ConfirmRegister(c *fiber.Ctx) error {

	request := c.Request()

	var req ConfirmRegisterRequest
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

	ver, err := postgreDB.GetVerificationCodeWithUserId(req.UserId)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "verification code not found",
		}

		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	isValidCode := utils.CheckHashPassword(req.ValidationCode, ver.VerificationCodeHash)
	if !isValidCode {
		err = errors.New("invalid validation code for this user")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "invalid validation code for this user",
		}

		return c.Status(fiber.StatusForbidden).JSON(resp)

	}

	user, err := postgreDB.GetUserWithId(req.UserId)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "user not found",
		}

		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	user.IsVerified = true
	err = postgreDB.UpdateUser(&user)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "failed to update user",
		}

		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	_ = postgreDB.DeleteVerificationWithId(ver.Id)

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
		LoginAttemptCount: user.LoginAttemptCount,
		IsVerified:        user.IsVerified,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
