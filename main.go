package main

import (
	"net/http"

	"github.com/kiyutink/sowhenthen/poll"
)

func main() {
	srv := NewServer(poll.NewMemeoryStorer())
	srv.routes()
	http.ListenAndServe(":80", srv)
}
