package main

import "net/http"

func (c *Controller) handleViewsCreatePoll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/create.html")
	}
}

func (c *Controller) handleViewsVote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/vote.html")
	}
}
