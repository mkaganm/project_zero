package consumer

import (
	"encoding/json"
	"mailerservice/internal/config"
	"mailerservice/internal/mail"
	"mailerservice/internal/utils"
)

type MailMessage struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

// ConsumeMailerQueue is a function to consume message from RabbitMQ for mailerservice
func ConsumeMailerQueue() {

	// Connect to RabbitMQ
	conn := Connect()
	defer Close(conn)
	// Create a channel
	ch := CreateChannel(conn)
	defer CloseChannel(ch)

	// Declare a queue
	msqs, err := ch.Consume(
		"mailer_queue", // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	utils.LogErr("Failed to register a consumer", err)

	// Create a channel to receive forever
	forever := make(chan bool)
	// Consume messages
	go func() {
		for d := range msqs {
			println("Received a message: ", string(d.Body))

			var mailMessage MailMessage
			err := json.Unmarshal(d.Body, &mailMessage)
			utils.LogErr("Error while unmarshalling request: ", err)

			err = mail.SendMail(mail.Mail{
				From:    config.EnvConfigs.MailerSenderAddress,
				To:      mailMessage.To,
				Subject: mailMessage.Subject,
				Body:    mailMessage.Body,
			})
			utils.LogErr("Error when sending mail in consumer", err)

		}
	}()
	// Block the channel
	<-forever
}
