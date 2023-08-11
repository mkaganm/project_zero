package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
	"userservice/internal/data/redisDB"
	"userservice/internal/messages/producer"
	"userservice/internal/utils"
)

// LoggingMiddleware is a middleware that logs all requests
func LoggingMiddleware(c *fiber.Ctx) error {

	startTime := time.Now()

	err := c.Next()

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	producer.PublishMongoLogMessage(producer.MongoLogMessage{
		Collection:     "userservice",
		Source:         c.IP(),
		Method:         c.Path(),
		Request:        string(c.Request().Body()),
		RequestHeader:  string(c.Request().Header.Header()),
		Response:       string(c.Response().Body()),
		ResponseHeader: string(c.Response().Header.Header()),
		Duration:       duration.String(),
		Status:         c.Response().StatusCode(),
	})

	return err
}

// CookieAuth is a middleware that checks if the user is authenticated
func CookieAuth(c *fiber.Ctx) error {

	cookie := c.Cookies("session")
	var cookieData redisDB.CookieData
	_ = json.Unmarshal([]byte(cookie), &cookieData)

	reqBody := make(map[string]interface{})
	_ = c.BodyParser(&reqBody)

	if uint64(reqBody["user_id"].(float64)) != cookieData.UserId {

		utils.LogInfo("user_id in cookie and request body does not match")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        nil,
			ErrorMessage: unAuthMsg,
		}

		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	tokenOK := utils.CheckToken(cookieData.SessionToken)
	if !tokenOK {

		utils.LogInfo("session token is not valid")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        nil,
			ErrorMessage: unAuthMsg,
		}

		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	var redisData *redisDB.CookieData
	redisData, err := redisDB.GetCookieData(cookieData.Key)
	if redisData.SessionToken != cookieData.SessionToken || err != nil {

		utils.LogInfo("session token in cookie and redis does not match")

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: unAuthMsg,
		}

		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	err = c.Next()
	return err
}
