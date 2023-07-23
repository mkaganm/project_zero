package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"loggerservice/internal/config"
	"loggerservice/internal/data/mongo"
	"loggerservice/internal/utils"
	"time"
)

type LogRequest struct {
	Collection     string      `json:"collection" validate:"required"`
	Request        interface{} `json:"request"`
	RequestHeader  interface{} `json:"requestHeader"`
	Response       interface{} `json:"response"`
	ResponseHeader interface{} `json:"responseHeader"`
}

// SendLog2Mongo sends success log to mongo
func SendLog2Mongo(c *fiber.Ctx) error {

	request := c.Request()

	go func() {

		var req LogRequest
		err := json.Unmarshal(request.Body(), &req)
		utils.LogErr("Error unmarshaling request body:", err)

		client := mongo.InitDB()
		defer mongo.CloseDB(client)

		db := client.Database(config.EnvConfigs.MongoDb)
		collection := db.Collection(req.Collection)

		LogMongo := mongo.LogMongo{
			Request:        req.Request,
			RequestHeader:  req.RequestHeader,
			Response:       req.Response,
			ResponseHeader: req.ResponseHeader,
			Timestamp:      time.Now(),
		}

		_, err = collection.InsertOne(c.Context(), LogMongo)
		utils.LogErr("Error inserting log to MongoDB:", err)

	}()

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"status": "success",
	})
}
