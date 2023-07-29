package api

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/mailer")

	routes.Use(LoggingMiddleware)
	routes.Post("/send-mail", SendMail)

}
