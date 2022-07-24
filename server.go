package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	pollController *PollController
	router         *chi.Mux
}

func (s *Server) Routes() {
	s.router.Get("/polls", s.pollController.GetMany)
	s.router.Get("/polls/{id}", s.pollController.GetOne)
	s.router.Post("/polls", s.pollController.Post)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(storer PollStorer) *Server {
	return &Server{
		pollController: NewPollController(storer),
		router:         chi.NewRouter(),
	}
}
