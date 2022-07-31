package main

func (s *Server) routes() {
	s.router.Get("/polls/{id}", s.pollController.HandleGetOne())
	s.router.Post("/polls", s.pollController.HandlePost())
	s.router.Get("/votes", s.voteController.HandleGetOne())
	s.router.Post("/votes", s.voteController.HandlePost())
}
