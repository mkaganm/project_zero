package controller

import (
	"github.com/gofiber/fiber/v2"
	"userservice/pkg/services"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/user")

	routes.Post("/register", services.Register)
	routes.Post("/login", services.Login)
	routes.Patch("/change-password", services.ChancePassword)
	routes.Post("/confirm-register", services.ConfirmRegister)
	routes.Post("/send-verification-code", services.SendVerificationCode)
	routes.Post("/forgot-password", services.ForgotPassword)

}
