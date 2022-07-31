package vote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Storer interface {
	Create(ctx context.Context, vote Vote) (Vote, error)
}

type Controller struct {
	storer Storer
}

func NewController(storer Storer) *Controller {
	return &Controller{storer}
}

func (c *Controller) HandlePost() http.HandlerFunc {
	type request struct {
		PollId    string `json:"poll_id"`
		Option    string `json:"option"`
		VoterName string `json:"voter_name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error decoding json: %v", err)))
			return
		}

		vote := Vote(req)
		_, err = c.storer.Create(context.Background(), vote)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error creating vote: %v", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(vote)
	}
}

func (c *Controller) HandleGetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("here you'll see votes"))
	}
}
