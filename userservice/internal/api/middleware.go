package api

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"userservice/internal/clients/loggerservice"
	"userservice/internal/utils"
)

// LoggingMiddleware is a middleware that logs all requests
func LoggingMiddleware(c *fiber.Ctx) error {

	startTime := time.Now()

	err := c.Next()

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	logData := loggerservice.Log{
		Source:         c.IP(),
		Method:         c.Path(),
		Request:        string(c.Request().Body()),
		RequestHeader:  string(c.Request().Header.Header()),
		Response:       string(c.Response().Body()),
		ResponseHeader: string(c.Response().Header.Header()),
		Duration:       duration.String(),
		Status:         c.Response().StatusCode(),
	}
	loggerservice.SendLog(logData)

	return err
}

// CookieAuth is a middleware that checks if the user is authenticated
func CookieAuth(c *fiber.Ctx) error {

	cookie := c.Cookies("session")
	print(cookie)

	reqBody := make(map[string]interface{})
	err := c.BodyParser(&reqBody)
	utils.LogErr("Error parsing request body", err)

	err = c.Next()
	return err
}
