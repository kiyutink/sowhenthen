package main

func (s *Server) routes() {
	s.router.Post("/polls", s.pollController.HandlePost())
	s.router.Get("/polls/{id}", s.pollController.HandleGetOne())
	s.router.Get("/polls/{pollId}/votes", s.voteController.HandleGetMany())
	s.router.Post("/polls/{pollId}/votes", s.voteController.HandlePost())
}
