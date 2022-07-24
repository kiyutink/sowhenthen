package main

import (
	"net/http"

	"github.com/kiyutink/sowhenthen/poll"
)

func main() {
	srv := NewServer(poll.NewMemeoryStorer())
	srv.Routes()
	http.ListenAndServe(":80", srv)
}
