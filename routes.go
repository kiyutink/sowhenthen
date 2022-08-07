package main

func (s *Server) routes() {
	s.router.Post("/api/polls", s.controller.handlePollsCreateOne())
	s.router.Get("/api/polls/{id}", s.controller.handlePollsGetOne())
	s.router.Get("/api/polls/{pollId}/votes", s.controller.handleVotesGetMany())
	s.router.Post("/api/polls/{pollId}/votes", s.controller.handleVotesCreateOne())
	s.router.Get("/", s.controller.handleViewsCreatePoll())
	s.router.Get("/{id}", s.controller.handleViewsVote())
}
