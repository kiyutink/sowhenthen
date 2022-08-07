package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/storage"
)

type Server struct {
	controller *Controller
	router     *chi.Mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(s storage.Storage) *Server {
	return &Server{
		controller: newController(s),
		router:     chi.NewRouter(),
	}
}
