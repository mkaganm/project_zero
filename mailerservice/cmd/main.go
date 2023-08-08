package main

import (
	"mailerservice/internal/consumer"
)

func main() {

	// Consume mailer_queue
	consumer.ConsumeMailerQueue()
}
