package main

import (
	"mailerservice/internal/messages/consumer"
)

func main() {

	// Consume mailer_queue
	consumer.ConsumeMailerQueue()
}
