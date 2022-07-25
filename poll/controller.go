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
	Delete(ctx context.Context, id string) error
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
	type response struct {
		Id      string   `json:"id"`
		Title   string   `json:"title"`
		Options []string `json:"options"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		poll, err := c.storer.Create(ctx, Poll{
			Title:   "Just another poll",
			Options: []string{"Option 1", "Option 2"},
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(response(poll))
	}
}

func (c *Controller) HandleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		c.storer.Delete(context.Background(), id)
		w.WriteHeader(http.StatusOK)
		fmt.Println(c.storer.Dump())
	}
}

func NewController(storer Storer) *Controller {
	return &Controller{storer}
}
