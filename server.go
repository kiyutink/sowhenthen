package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	controller *Controller
	router     *chi.Mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(pollStorage PollStorage, voteStorage VoteStorage) *Server {
	return &Server{
		controller: NewController(pollStorage, voteStorage),
		router:     chi.NewRouter(),
	}
}
