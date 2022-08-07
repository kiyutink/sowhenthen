package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kiyutink/sowhenthen/mongo"
	mongoDB "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const timeout = time.Second * 30

func newMongoClient(url string) (*mongoDB.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	client, err := mongoDB.Connect(ctx, options.Client().ApplyURI(url))
	return client, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mongoClient, err := newMongoClient(os.Getenv("MONGO_URL"))
	defer mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	srv := NewServer(mongo.NewStorage(mongoClient))
	srv.routes()
	socket := fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Printf("server listening on socket %v\n", socket)
	err = http.ListenAndServe(socket, srv)
	fmt.Println(err)
}
