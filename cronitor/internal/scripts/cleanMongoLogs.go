package scripts

import (
	"context"
	"cronitor/internal/clients/elastic"
	"cronitor/internal/config"
	"cronitor/internal/data"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// CleanMongoLogs is a function that cleans mongo logs older than 30 days
func CleanMongoLogs() {

	esLog := make(map[string]interface{})
	esLog["status"] = "success"
	esLog["tag"] = "[CleanMongoLogs]"
	esLog["message"] = "Mongo logs cleaned !"
	esLog["timestamp"] = time.Now()

	client := data.InitMongoDB()
	defer data.CloseMongoDB(client)

	db := client.Database(config.EnvConfigs.MongoDb)

	// Before 30 days
	oneMonthAgo := time.Now().Add(-30 * 24 * time.Hour)

	// Delete logs from MongoDB
	collection := db.Collection("userservice")
	filter := bson.M{"timestamp": bson.M{"$lt": oneMonthAgo}}
	_, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Default().Println("Error while deleting logs from MongoDB")
		esLog["error"] = err.Error()
		esLog["status"] = "failed"
	}

	// Delete logs from MongoDB
	collection = db.Collection("mailerservice")
	filter = bson.M{"timestamp": bson.M{"$lt": oneMonthAgo}}
	_, err = collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Default().Println("Error while deleting logs from MongoDB")
		esLog["error"] = err.Error()
		esLog["status"] = "failed"
	}

	// send log to elasticsearch
	elastic.SendLog(esLog)

}
