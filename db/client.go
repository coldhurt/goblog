package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	if client != nil && ctx != nil && cancel != nil {
		return client, ctx, cancel
	}

	connectionURI := viper.GetString("MONGODB_URI")
	connectTimeout := viper.GetInt("MONGODB_TIMTOUT")

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(connectTimeout)*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

func GetCollection(collectionName string) (*mongo.Client, *mongo.Collection, context.Context, context.CancelFunc) {
	client, ctx, cancel := GetConnection()
	collection := client.Database(viper.GetString("MONGODB_DATABASE")).Collection(collectionName)
	return client, collection, ctx, cancel
}
