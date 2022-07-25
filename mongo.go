package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongoClient(url string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // TODO: Check if context is right
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	return client, err
}
