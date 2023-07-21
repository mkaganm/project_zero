package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"mailerservice/internal/config"
	"mailerservice/internal/mail"
)

type sendMailRequest struct {
	To      []string `json:"to" validate:"required"`
	Subject string   `json:"subject" validate:"required"`
	Body    string   `json:"body" validate:"required"`
}

// SendMail sends an email to the user
func SendMail(c *fiber.Ctx) error {

	request := c.Request()

	var req sendMailRequest

	err := json.Unmarshal(request.Body(), &req)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}

	mailInfo := mail.Mail{
		From:    config.EnvConfigs.MailerSenderAddress,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}
	err = mail.SendMail(mailInfo)
	if err != nil {
		return err
	}

	resp := Response{
		Status: "success",
		Data:   map[string]interface{}{},
	}

	resp.Data["infos"] = mailInfo
	return c.Status(fiber.StatusOK).JSON(resp)
}
