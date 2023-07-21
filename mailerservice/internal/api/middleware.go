package api

import "github.com/gofiber/fiber/v2"

// ErrorHandler is a middleware to handle errors
func ErrorHandler(c *fiber.Ctx, err error) error {
	if err != nil {
		// Check the error type and set the appropriate status code
		statusCode := fiber.StatusInternalServerError
		errorMessage := "internal server error"

		switch e := err.(type) {
		case *fiber.Error:
			statusCode = e.Code
			errorMessage = e.Message
		}

		return c.Status(statusCode).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: errorMessage,
		})
	}

	return nil
}
