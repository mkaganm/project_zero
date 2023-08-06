package api

import (
	"github.com/gofiber/fiber/v2"
	"mailerservice/internal/clients/logger"
	"time"
)

// LoggingMiddleware is a middleware that logs all requests
func LoggingMiddleware(c *fiber.Ctx) error {

	startTime := time.Now()

	err := c.Next()

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	logData := logger.Log{
		Source:         c.IP(),
		Method:         c.Path(),
		Request:        string(c.Request().Body()),
		RequestHeader:  string(c.Request().Header.Header()),
		Response:       string(c.Response().Body()),
		ResponseHeader: string(c.Response().Header.Header()),
		Duration:       duration.String(),
		Status:         c.Response().StatusCode(),
	}
	logger.SendMongoLog(logData)

	return err
}
