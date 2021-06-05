package mongoclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is singleton of db client/pool
var ClientMongo *mongo.Client = nil

// NewConnectedMongoDriver define a connection to DB
func NewConnectedMongoDriver() *mongo.Client {
	if ClientMongo == nil {
		config := NewDBConfig()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		clientOptions := options.Client().ApplyURI(fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			config.Username,
			config.Password,
			config.Addr,
			config.Port,
			config.DBName,
			// "admin",
		))
		defer cancel()
		database, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = database.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")
		ClientMongo = database
		return database
	}
	return ClientMongo
}
