package producer

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"userservice/internal/messages"
	"userservice/internal/utils"
)

type MongoLogMessage struct {
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

type ElasticLogMessage struct {
	Index string      `json:"index" validate:"required"`
	Data  interface{} `json:"data" validate:"required"`
}

// PublishMongoLogMessage is a function to publish message to RabbitMQ for loggerservice
func PublishMongoLogMessage(data MongoLogMessage) {
	// Connect to RabbitMQ
	conn := messages.Connect()
	defer messages.Close(conn)
	// Create a channel
	ch := messages.CreateChannel(conn)
	defer messages.CloseChannel(ch)

	_, err := ch.QueueDeclare(
		"logger_mongo_queue", // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	utils.LogErr("Failed to declare a queue", err)

	jsonData, err := json.Marshal(data)
	utils.LogErr("Error while marshalling request: ", err)

	err = ch.Publish(
		"",
		"logger_mongo_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	utils.LogErr("Failed to publish a message", err)
}

// PublishElasticLogMessage is a function to publish message to RabbitMQ for loggerservice
func PublishElasticLogMessage(data ElasticLogMessage) {
	// Connect to RabbitMQ
	conn := messages.Connect()
	defer messages.Close(conn)
	// Create a channel
	ch := messages.CreateChannel(conn)
	defer messages.CloseChannel(ch)

	_, err := ch.QueueDeclare(
		"logger_elastic_queue", // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	utils.LogErr("Failed to declare a queue", err)

	jsonData, err := json.Marshal(data)
	utils.LogErr("Error while marshalling request: ", err)

	err = ch.Publish(
		"",
		"logger_elastic_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	utils.LogErr("Failed to publish a message", err)
}
