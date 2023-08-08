package producer

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"userservice/internal/utils"
)

type LogMessage struct {
	Collection     string      `json:"collection"`
	Source         string      `json:"source"`
	Method         string      `json:"method"`
	Request        interface{} `json:"request"`
	RequestHeader  interface{} `json:"request_header"`
	Response       interface{} `json:"response"`
	ResponseHeader interface{} `json:"response_header"`
	Duration       string      `json:"duration"`
	Status         int         `json:"status"`
}

// PublishLogMessage is a function to publish message to RabbitMQ for mailerservice
func PublishLogMessage(data LogMessage) {
	// Connect to RabbitMQ
	conn := Connect()
	defer Close(conn)
	// Create a channel
	ch := CreateChannel(conn)
	defer CloseChannel(ch)

	_, err := ch.QueueDeclare(
		"mailer_queue", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	utils.LogErr("Failed to declare a queue", err)

	jsonData, err := json.Marshal(data)
	utils.LogErr("Error while marshalling request: ", err)

	err = ch.Publish(
		"",
		"mailer_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	utils.LogErr("Failed to publish a message", err)
}
