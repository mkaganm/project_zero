package consumer

import (
	"fmt"
	"github.com/streadway/amqp"
	"mailerservice/internal/config"
	"mailerservice/internal/utils"
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
	utils.LogErr("Failed to close connection to RabbitMQ", err)
}

// CloseChannel channel
func CloseChannel(ch *amqp.Channel) {
	err := ch.Close()
	utils.LogErr("Failed to close channel", err)
}
