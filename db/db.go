package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Database, *mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, err
	}

	// Ping the database to ensure connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, nil, err
	}

	log.Println("✅ Connected to MongoDB")

	return client.Database("grpc-echo-mongo-app"), client, nil
}
