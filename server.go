package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/poll"
)

type Server struct {
	pollController *poll.Controller
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

func NewServer(storer poll.Storer) *Server {
	return &Server{
		pollController: poll.NewController(storer),
		router:         chi.NewRouter(),
	}
}
