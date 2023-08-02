package api

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/user")

	routes.Use(LoggingMiddleware)
	routes.Post("/register", Register)
	routes.Post("/login", Login)
	routes.Patch("/change-password", CookieAuth, ChancePassword) // CookieAuth is a auth middleware
	routes.Post("/confirm-register", ConfirmRegister)
	routes.Post("/send-verification-code", SendVerificationCode)
	routes.Post("/forgot-password", ForgotPassword)

}
