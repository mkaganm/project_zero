package producer

import (
	"fmt"
	"github.com/streadway/amqp"
	"userservice/internal/config"
	"userservice/internal/utils"
)

var RabbitDSN *string

func init() {
	dsn := initDSN()
	RabbitDSN = &dsn
}

// Create dsn a new RabbitMQ
func initDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.EnvConfigs.RabbitMQUser,
		config.EnvConfigs.RabbitMQPass,
		config.EnvConfigs.RabbitMQHost,
		config.EnvConfigs.RabbitMQPort)
}

// Connect to RabbitMQ
func Connect() *amqp.Connection {

	conn, err := amqp.Dial(*RabbitDSN)
	utils.FatalErr("Failed to connect to RabbitMQ", err)

	return conn
}

// CreateChannel a new channel
func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	utils.FatalErr("Failed to open a channel", err)

	return ch
}

// Close connection to RabbitMQ
func Close(conn *amqp.Connection) {
	err := conn.Close()
	utils.FatalErr("Failed to close connection to RabbitMQ", err)
}

// CloseChannel channel
func CloseChannel(ch *amqp.Channel) {
	err := ch.Close()
	utils.FatalErr("Failed to close channel", err)
}

// PublishMailerMessage is a function to publish message to RabbitMQ for mailerservice
func PublishMailerMessage() {
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

	err = ch.Publish(
		"",
		"mailer_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello from userservice"),
		})
	utils.LogErr("Failed to publish a message", err)
}
