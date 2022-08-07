package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kiyutink/sowhenthen/poll"
	"github.com/kiyutink/sowhenthen/vote"
)

const timeout = time.Second * 30

func main() {
	mongoClient, err := newMongoClient("mongodb://localhost:27017")
	defer mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	srv := NewServer(poll.NewMongoStorage(mongoClient), vote.NewMongoStorage(mongoClient))
	srv.routes()
	fmt.Println("listening on localhost:8001")
	err = http.ListenAndServe("localhost:8001", srv)
	fmt.Println(err)
}
