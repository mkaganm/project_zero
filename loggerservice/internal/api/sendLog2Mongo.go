package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"loggerservice/internal/config"
	"loggerservice/internal/data/mongo"
	"loggerservice/internal/utils"
	"time"
)

type LogRequest struct {
	Collection     string      `json:"collection" validate:"required"`
	Source         string      `json:"source"`
	Method         string      `json:"method"`
	Request        interface{} `json:"request"`
	RequestHeader  interface{} `json:"request_header"`
	Response       interface{} `json:"response"`
	ResponseHeader interface{} `json:"response_header"`
	Duration       string      `json:"duration"`
	Status         int         `json:"status"`
}

// SendLog2Mongo sends success log to mongo
func SendLog2Mongo(c *fiber.Ctx) error {

	request := c.Request()

	var req LogRequest
	err := json.Unmarshal(request.Body(), &req)
	utils.LogErr("Error unmarshaling request body:", err)

	validate := validator.New()
	err = validate.Struct(req)
	utils.LogErr("Error validating request body:", err)

	client := mongo.InitDB()
	defer mongo.CloseDB(client)

	db := client.Database(config.EnvConfigs.MongoDb)
	collection := db.Collection(req.Collection)

	LogMongo := mongo.LogMongo{
		Source:         req.Source,
		Method:         req.Method,
		Request:        req.Request,
		RequestHeader:  req.RequestHeader,
		Response:       req.Response,
		ResponseHeader: req.ResponseHeader,
		Duration:       req.Duration,
		Status:         req.Status,
		Timestamp:      time.Now(),
	}

	result, err := collection.InsertOne(c.Context(), LogMongo)
	utils.LogErr("Error inserting log to mongo:", err)

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status": "success",
		"_id":    result.InsertedID,
	})
}
