package poll

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Storer interface {
	Create(ctx context.Context, p Poll) (Poll, error)
	GetOne(ctx context.Context, id string) (Poll, error)
	Dump() interface{}
}

type Controller struct {
	storer Storer
}

func (c *Controller) HandleGetOne() http.HandlerFunc {
	type response struct {
		Id      string   `json:"id"`
		Title   string   `json:"title"`
		Options []string `json:"options"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		poll, err := c.storer.GetOne(context.Background(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response(poll))
		fmt.Println(c.storer.Dump())
	}
}

func (c *Controller) HandlePost() http.HandlerFunc {
	type request struct {
		Title   string   `json:"title"`
		Options []string `json:"options"`
	}

	type response struct {
		Id      string   `json:"id"`
		Title   string   `json:"title"`
		Options []string `json:"options"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error decoding json: %v", err)))
		}

		p := Poll{}
		p.Options = req.Options
		p.Title = req.Title

		poll, err := c.storer.Create(ctx, p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(response(poll))
	}
}

func NewController(storer Storer) *Controller {
	return &Controller{storer}
}
