package api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"loggerservice/internal/data/elastic"
	"loggerservice/internal/utils"
)

type ElasticRequest struct {
	Index string      `json:"index" validate:"required"`
	Data  interface{} `json:"data" validate:"required"`
}

// SendLog2Elastic sends log to Elasticsearch
func SendLog2Elastic(c *fiber.Ctx) error {

	request := c.Request()
	var req ElasticRequest

	err := json.Unmarshal(request.Body(), &req)
	if err != nil {
		utils.LogErr("Error unmarshaling request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: err.Error(),
		})
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		utils.LogErr("Error validating request:", err)
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: err.Error(),
		})
	}

	err = elastic.SendLog2Elastic(req.Data, req.Index)
	if err != nil {
		utils.LogErr("Error sending log to Elasticsearch:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"status": "success",
	})
}
