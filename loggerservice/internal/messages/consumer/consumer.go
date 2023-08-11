package consumer

import (
	"context"
	"encoding/json"
	"loggerservice/internal/config"
	"loggerservice/internal/data/elastic"
	"loggerservice/internal/data/mongo"
	"loggerservice/internal/messages"
	"loggerservice/internal/utils"
	"time"
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

// ConsumeMessages consumes logs from rabbitmq and sends to mongo and elastic
func ConsumeMessages() {

	// Connect to RabbitMQ
	conn := messages.Connect()
	defer messages.Close(conn)
	// Create a channel
	ch := messages.CreateChannel(conn)
	defer messages.CloseChannel(ch)

	// Declare a queue
	msqs, err := ch.Consume(
		"logger_mongo_queue", // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	utils.LogErr("Failed to register a consumer", err)

	// Declare a queue
	msgsEslog, err := ch.Consume(
		"logger_elastic_queue", // queue
		"",                     // consumer
		true,                   // auto-ack
		false,                  // exclusive
		false,                  // no-local
		false,                  // no-wait
		nil,                    // args
	)
	utils.LogErr("Failed to register a consumer", err)

	// Create a channel to receive forever
	forever := make(chan bool)

	// Consume messages
	go func() {
		for d := range msqs {
			println("Received a message: ", string(d.Body))

			var logMsg MongoLogMessage
			err := json.Unmarshal(d.Body, &logMsg)
			utils.LogErr("Error while unmarshalling message: ", err)

			client := mongo.InitDB()

			db := client.Database(config.EnvConfigs.MongoDb)
			collection := db.Collection(logMsg.Collection)

			_, err = collection.InsertOne(context.TODO(), mongo.LogMongo{
				Source:         logMsg.Source,
				Method:         logMsg.Method,
				Request:        logMsg.Request,
				RequestHeader:  logMsg.RequestHeader,
				Response:       logMsg.Response,
				ResponseHeader: logMsg.ResponseHeader,
				Duration:       logMsg.Duration,
				Status:         logMsg.Status,
				Timestamp:      time.Now(),
			})
			utils.LogErr("Error inserting log to mongo:", err)
			utils.LogInfo("Inserted log to mongo:")
			utils.LogErr("Error when sending mail in consumer", err)

			mongo.CloseDB(client)
		}

	}()

	go func() {
		for d := range msgsEslog {
			println("Received a message: ", string(d.Body))

			var logMsg ElasticLogMessage
			err := json.Unmarshal(d.Body, &logMsg)
			utils.LogErr("Error while unmarshalling message: ", err)

			err = elastic.SendLog2Elastic(logMsg.Data, logMsg.Index)
			utils.LogErr("Error when sending log to elastic in consumer", err)

		}
	}()

	// Block the channel
	<-forever
}
