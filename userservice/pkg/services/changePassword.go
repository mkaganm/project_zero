package services

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"userservice/pkg/clients/logger"
	"userservice/pkg/data/repository"
	"userservice/pkg/utils"
)

type changePasswordRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	NewPassword1 string `json:"new_password_1" validate:"required,CheckPassword"`
	NewPassword2 string `json:"new_password_2" validate:"required,CheckPassword"`
}

func ChancePassword(c *fiber.Ctx) error {

	request := c.Request()

	var req changePasswordRequest

	err := json.Unmarshal(request.Body(), &req)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errMsg,
		}

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

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

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	if req.NewPassword1 != req.NewPassword2 {

		err = errors.New("new password does not match")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errMsg,
		}

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

		return c.Status(fiber.StatusBadRequest).JSON(resp)

	}

	user, err := repository.GetUserWithEmail(req.Email)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "user not found",
		}

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	if user.IsBlocked {

		err = errors.New("user is blocked please contact admin")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "user is blocked please contact admin",
		}

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	isValidPassword := utils.CheckHashPassword(req.Password, user.Password)
	if !isValidPassword {

		err = errors.New("invalid password provided for this user account")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "invalid password provided for this user account",
		}

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	if utils.CheckHashPassword(req.NewPassword1, user.Password) {

		err = errors.New("new password cannot be the same as the old password")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "new password cannot be the same as the old password",
		}

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	user.Password, _ = utils.HashPassword(req.NewPassword1)

	err = repository.UpdateUser(&user)
	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "unable to update user password",
		}

		logger.SendLog(logger.Log{
			Source:   utils.CurrentTrace(),
			Request:  req,
			Response: resp,
		})

		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := Response{
		Status: "success",
		Data:   map[string]interface{}{},
	}

	resp.Data["user"] = UserResponse{
		UserId:      user.Id,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	logger.SendLog(logger.Log{
		Source:   utils.CurrentTrace(),
		Request:  req,
		Response: resp,
	})

	return c.Status(fiber.StatusOK).JSON(resp)
}
