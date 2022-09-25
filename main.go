package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kiyutink/sowhenthen/mongo"
	mongoDB "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const timeout = time.Second * 30

func newMongoClient(url string) (*mongoDB.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	client, err := mongoDB.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return client, err
	}
	err = client.Ping(ctx, readpref.Primary())
	return client, err
}

func main() {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	mongoClient, err := newMongoClient(os.Getenv("MONGO_URL"))
	defer mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	srv := NewServer(mongo.NewStorage(mongoClient))
	srv.routes()
	fmt.Println("server listening on port", os.Getenv("PORT"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), srv)
	fmt.Println(err)
}
