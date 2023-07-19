package controller

import (
	"github.com/gofiber/fiber/v2"
	"loggerservice/pkg/services"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/logger")
	routes.Post("/success", services.SendSuccessLog)
	routes.Get("/log", services.GetLog)
}
