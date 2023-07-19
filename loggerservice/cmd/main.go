package main

import (
	"github.com/gofiber/fiber/v2"
	"loggerservice/pkg/config"
	"loggerservice/pkg/controller"
	"loggerservice/pkg/data/mongo"
	"loggerservice/pkg/utils"
)

func main() {

	// Initialize environment variables
	config.InitEnvConfigs()
	// Initialize database source name
	mongo.InitDSN()

	app := fiber.New(fiber.Config{
		//ReadTimeout:   time.Second * 15,
		//WriteTimeout:  time.Second * 15,
		Concurrency:  10,
		ServerHeader: "logger_service_v1",
		AppName:      "logger_service_v1",
	})

	// Register routes
	controller.RegisterRoutes(app)

	// Listen on port
	err := app.Listen(config.EnvConfigs.LocalServerPort)
	utils.FatalErr("Error while serving the api", err)

}
