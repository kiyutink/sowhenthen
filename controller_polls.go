package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/entities"
	"github.com/kiyutink/sowhenthen/storage"
)

func (c *Controller) handlePollsGetOne() http.HandlerFunc {
	type response struct {
		Id      string   `json:"id"`
		Title   string   `json:"title"`
		Options []string `json:"options"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		poll, err := c.storage.Poll.GetOne(ctx, id)
		if err != nil {
			switch err.(type) {
			case *storage.NotFoundError:
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response(poll))
	}
}

func (c *Controller) handlePollsCreateOne() http.HandlerFunc {
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

		p := entities.Poll{}
		p.Options = req.Options
		p.Title = req.Title

		poll, err := c.storage.Poll.Create(ctx, p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(response(poll))
	}
}
