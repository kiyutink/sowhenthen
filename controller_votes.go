package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/entities"
)

func (c *Controller) handleVotesCreateOne() http.HandlerFunc {
	type request struct {
		Options   []string `json:"options"`
		VoterName string   `json:"voterName"`
	}

	type response struct {
		PollId    string   `json:"pollId"`
		Options   []string `json:"options"`
		VoterName string   `json:"voterName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "pollId")
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error decoding json: %v", err)))
			return
		}

		vote := entities.Vote{}
		vote.PollId = pollId
		vote.Options = req.Options
		vote.VoterName = req.VoterName
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_, err = c.storage.Vote.Create(ctx, vote)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error creating vote: %v", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response(vote))
	}
}

func (c *Controller) handleVotesGetMany() http.HandlerFunc {
	type response []struct {
		PollId    string   `json:"pollId"`
		Options   []string `json:"options"`
		VoterName string   `json:"voterName"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "pollId")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		votes, err := c.storage.Vote.GetMany(ctx, pollId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		res := make(response, len(votes))
		for i, vote := range votes {
			res[i].Options = vote.Options
			res[i].PollId = vote.PollId
			res[i].VoterName = vote.VoterName
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
