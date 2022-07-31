package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/poll"
	"github.com/kiyutink/sowhenthen/vote"
)

type Server struct {
	pollController *poll.Controller
	voteController *vote.Controller
	router         *chi.Mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(pollStorer poll.Storer, voteStorer vote.Storer) *Server {
	return &Server{
		pollController: poll.NewController(pollStorer),
		voteController: vote.NewController(voteStorer),
		router:         chi.NewRouter(),
	}
}
