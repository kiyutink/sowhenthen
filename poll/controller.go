package poll

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Storer interface {
	Create(ctx context.Context, p Poll) (Poll, error)
	Delete(ctx context.Context, id string) error
	GetOne(ctx context.Context, id string) (Poll, error)
	GetMany(ctx context.Context) ([]Poll, error)
	Dump() interface{}
}

type Controller struct {
	storer Storer
}

func (pc *Controller) HandleGetMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		polls, _ := pc.storer.GetMany(context.Background()) // TODO: Unignore error
		names := []string{}
		for _, poll := range polls {
			names = append(names, poll.Id)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strings.Join(names, ", ")))
		fmt.Println(pc.storer.Dump())
	}
}

func (pc *Controller) HandleGetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		poll, _ := pc.storer.GetOne(context.Background(), id) // TODO: Unignore error
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(poll.Title))
		fmt.Println(pc.storer.Dump())
	}
}

func (pc *Controller) HandlePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pc.storer.Create(context.Background(), Poll{ // TODO: Unignore error, use returned Poll
			Title: "Just another poll",
		})
		w.WriteHeader(http.StatusCreated)
		fmt.Println(pc.storer.Dump())
	}
}

func (pc *Controller) HandleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		pc.storer.Delete(context.Background(), id)
		w.WriteHeader(http.StatusOK)
		fmt.Println(pc.storer.Dump())
	}
}

func NewController(storer Storer) *Controller {
	return &Controller{storer}
}
