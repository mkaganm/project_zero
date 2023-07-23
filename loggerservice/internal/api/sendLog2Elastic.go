package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"loggerservice/internal/data/elastic"
)

// SendLog2Elastic sends log to Elasticsearch
func SendLog2Elastic(c *fiber.Ctx) error {

	request := c.Request()
	var data interface{}
	_ = json.Unmarshal(request.Body(), &data)

	go func() {
		elastic.SendLog2Elastic(data)
	}()

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"status": "success",
	})
}
