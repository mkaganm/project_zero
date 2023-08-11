package producer

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"userservice/internal/messages"
	"userservice/internal/utils"
)

// MailMessage is a struct that represents a request to send an email
type MailMessage struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

// PublishMailerMessage is a function to publish message to RabbitMQ for mailerservice
func PublishMailerMessage(data MailMessage) {
	// Connect to RabbitMQ
	conn := messages.Connect()
	defer messages.Close(conn)
	// Create a channel
	ch := messages.CreateChannel(conn)
	defer messages.CloseChannel(ch)

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
