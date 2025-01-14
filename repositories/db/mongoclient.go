package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	client     *mongo.Client
	clientOnce sync.Once
)

// GetMongoClient returns a singleton MongoDB client
func GetMongoClient(uri string) *mongo.Client {
	clientOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var err error
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		// Verify connection
		if err := client.Ping(ctx, nil); err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}
		fmt.Println("Connected to MongoDB")
	})
	return client
}

func Disconnect() {
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
			//ch <- err.Error()
		}
	}
	
	fmt.Println("Disconnected MongoDB")
}
