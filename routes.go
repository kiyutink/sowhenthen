package main

func (s *Server) routes() {
	s.router.Get("/polls", s.pollController.GetMany)
	s.router.Get("/polls/{id}", s.pollController.GetOne)
	s.router.Post("/polls", s.pollController.Post)
}
