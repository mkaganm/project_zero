package api

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"loggerservice/internal/config"
	"loggerservice/internal/data/mongo"
	"loggerservice/internal/utils"
)

type getLogRequest struct {
	Collection string `json:"collection" validate:"required"`
	Id         string `json:"id" validate:"required"`
}

// GetLog is a service to get logs
func GetLog(c *fiber.Ctx) error {
	request := c.Request()
	var req getLogRequest

	err := json.Unmarshal(request.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "invalid request body",
		})
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "invalid request body",
		})
	}

	client := mongo.InitDB()
	defer mongo.CloseDB(client)

	db := client.Database(config.EnvConfigs.MongoDb)
	collection := db.Collection(req.Collection)

	id, err := primitive.ObjectIDFromHex("64a7ca966e0fe788745d2e87")
	utils.LogErr("error while converting id", err)

	filter := bson.M{"_id": id}
	var result bson.M

	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "error while getting log",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
