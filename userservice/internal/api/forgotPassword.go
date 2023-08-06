package api

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"userservice/internal/data/postgreDB"
	"userservice/internal/utils"
)

type ForgotPasswordRequest struct {
	UserId           uint64 `json:"user_id" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	VerificationCode string `json:"verification_code" validate:"required"`
	NewPassword1     string `json:"new_password_1" validate:"required,CheckPassword"`
	NewPassword2     string `json:"new_password_2" validate:"required,CheckPassword"`
}

// ForgotPassword service
func ForgotPassword(c *fiber.Ctx) error {

	request := c.Request()

	var req ForgotPasswordRequest
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
	_ = validate.RegisterValidation("CheckPassword", utils.CheckPassword)
	err = validate.Struct(req)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errMsg,
		}

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	if req.NewPassword1 != req.NewPassword2 {

		err = errors.New("new passwords are not the same")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errMsg,
		}

		return c.Status(fiber.StatusBadRequest).JSON(resp)

	}

	ver, err := postgreDB.GetVerificationCodeWithUserId(req.UserId)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "no verification code found for this user",
		}

		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	isValidCode := utils.CheckHashPassword(req.VerificationCode, ver.VerificationCodeHash)
	if !isValidCode {

		err = errors.New("wrong verification code")

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

	if utils.CheckHashPassword(req.NewPassword1, user.Password) {

		err = errors.New("new password cannot be the same as the old password")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errMsg,
		}

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	newPassword, _ := utils.HashPassword(req.NewPassword1)

	user.Password = newPassword
	err = postgreDB.UpdateUser(&user)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "failed to update user",
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
		LoginAttemptCount: user.LoginAttemptCount,
		IsVerified:        user.IsVerified,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
