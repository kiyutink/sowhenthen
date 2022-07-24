package main

import (
	"net/http"
)

func main() {
	srv := NewServer(NewMemoryPollStorer())
	srv.Routes()
	http.ListenAndServe(":80", srv)
}
