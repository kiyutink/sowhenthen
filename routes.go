package main

func (s *Server) routes() {
	s.router.Get("/polls/{id}", s.pollController.HandleGetOne())
	s.router.Post("/polls", s.pollController.HandlePost())
	s.router.Delete("/polls", s.pollController.HandleDelete())
}
