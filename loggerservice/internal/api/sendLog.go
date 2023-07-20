package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"loggerservice/internal/config"
	"loggerservice/internal/data/mongo"
	"time"
)

type SuccessLogRequest struct {
	Collection     string      `json:"collection" validate:"required"`
	Request        interface{} `json:"request"`
	RequestHeader  interface{} `json:"requestHeader"`
	Response       interface{} `json:"response"`
	ResponseHeader interface{} `json:"responseHeader"`
}

// SendSuccessLog sends success log to mongo
func SendSuccessLog(c *fiber.Ctx) error {

	request := c.Request()

	var req SuccessLogRequest

	err := json.Unmarshal(request.Body(), &req)
	if err != nil {
		log.Default().Println("error while unmarshalling request body", err)
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

	log := mongo.SuccessLog{
		Request:   req.Request,
		Response:  req.Response,
		Timestamp: time.Now(),
	}

	result, err := collection.InsertOne(c.Context(), log)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: "failed to insert success log",
		})
	}

	insertedID := result.InsertedID

	response := Response{
		Status: "success",
		Data:   make(map[string]interface{}),
	}

	response.Data["inserted_id"] = insertedID
	response.Data["collection"] = req.Collection

	return c.Status(fiber.StatusOK).JSON(response)
}
