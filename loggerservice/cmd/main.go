package main

import (
	"github.com/gofiber/fiber/v2"
	"loggerservice/internal/api"
	"loggerservice/internal/config"
	"loggerservice/internal/data/mongo"
	"loggerservice/internal/utils"
)

func main() {

	// Initialize environment variables
	config.InitEnvConfigs()
	// Initialize mongo connection
	mongo.InitMongoDSN()

	app := fiber.New(fiber.Config{
		//ReadTimeout:   time.Second * 15,
		//WriteTimeout:  time.Second * 15,
		Concurrency:  10,
		ServerHeader: "logger_service_v1",
		AppName:      "logger_service_v1",
	})

	// Register routes
	api.RegisterRoutes(app)

	// Listen on port
	err := app.Listen(config.EnvConfigs.LocalServerPort)
	utils.FatalErr("Error while serving the api", err)

}
