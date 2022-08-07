package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kiyutink/sowhenthen/poll"
	"github.com/kiyutink/sowhenthen/vote"
)

const timeout = time.Second * 30

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
	srv := NewServer(poll.NewMongoStorage(mongoClient), vote.NewMongoStorage(mongoClient))
	srv.routes()
	socket := fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Printf("server listening on socket %v\n", socket)
	err = http.ListenAndServe(socket, srv)
	fmt.Println(err)
}
