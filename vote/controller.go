package vote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Storer interface {
	Create(ctx context.Context, vote Vote) (Vote, error)
	GetMany(ctx context.Context, pollId string) ([]Vote, error)
}

type Controller struct {
	storer Storer
}

func NewController(storer Storer) *Controller {
	return &Controller{storer}
}

func (c *Controller) HandlePost() http.HandlerFunc {
	type request struct {
		Option    string `json:"option"`
		VoterName string `json:"voter_name"`
	}

	type response struct {
		PollId    string `json:"poll_id"`
		Option    string `json:"option"`
		VoterName string `json:"voter_name"`
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

		vote := Vote{}
		vote.PollId = pollId
		vote.Option = req.Option
		vote.VoterName = req.VoterName
		_, err = c.storer.Create(context.Background(), vote)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error creating vote: %v", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response(vote))
	}
}

func (c *Controller) HandleGetMany() http.HandlerFunc {
	type response []struct {
		PollId    string `json:"poll_id"`
		Option    string `json:"option"`
		VoterName string `json:"voter_name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "pollId")
		votes, err := c.storer.GetMany(context.TODO(), pollId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		res := make(response, len(votes))
		for i, vote := range votes {
			res[i].Option = vote.Option
			res[i].PollId = vote.PollId
			res[i].VoterName = vote.VoterName
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
