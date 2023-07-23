package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"loggerservice/internal/config"
	"loggerservice/internal/utils"
	"time"
)

var DSN *string

type LogMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Source         string             `bson:"source"`
	Request        interface{}        `bson:"request"`
	RequestHeader  interface{}        `bson:"requestHeader"`
	Response       interface{}        `bson:"response"`
	ResponseHeader interface{}        `bson:"responseHeader"`
	Timestamp      time.Time          `bson:"timestamp"`
}

// InitDSN is a function that initializes the data source name
func InitDSN() {
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.EnvConfigs.MongoUser,
		config.EnvConfigs.MongoPass,
		config.EnvConfigs.MongoHost,
		config.EnvConfigs.MongoPort)

	dsn += "/?connect=direct"

	DSN = &dsn
}

// InitDB is a function that initializes the MongoDB client
func InitDB() *mongo.Client {

	log.Default().Println("Connecting to MongoDB...")

	dbURI := *DSN

	clientOptions := options.Client().ApplyURI(dbURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.FatalErr("Error while connecting to MongoDB", err)

	err = client.Ping(context.TODO(), nil)
	utils.FatalErr("Error while connecting to MongoDB", err)

	log.Default().Println("Connected to MongoDB.")

	return client
}

// CloseDB is a function that closes the MongoDB connection
func CloseDB(client *mongo.Client) {

	log.Default().Println("Closing MongoDB connection...")

	err := client.Disconnect(context.TODO())
	utils.FatalErr("Error while closing the MongoDB connection", err)

	log.Default().Println("MongoDB connection closed.")
}
