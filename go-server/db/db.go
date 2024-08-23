package db

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/emmanuelayinde/coody-restaurant/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	config.LoadConfig()
	mongoDB := config.AppConfig.DB_URL
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Connected")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("CoodyRestaurant").Collection(collectionName)
	return collection
}
