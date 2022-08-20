package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/entities"
	"github.com/kiyutink/sowhenthen/storage"
)

type testPollStorage struct {
	storage.Poll
}

func (tps *testPollStorage) Create(ctx context.Context, p entities.Poll) (entities.Poll, error) {
	newPoll := p
	newPoll.Id = "test-id"
	return newPoll, nil
}

func (tps *testPollStorage) GetOne(ctx context.Context, id string) (entities.Poll, error) {
	if id == "non-existent-poll" {
		return entities.Poll{}, &storage.NotFoundError{Identifier: id, Err: errors.New("poll doesn't exist")}
	}
	return entities.Poll{
		Id:      "test-id",
		Options: []string{"test-option-1", "test-option-2"},
		Title:   "test-title",
	}, nil
}

func newTestController() *Controller {
	testStorage := storage.Storage{Poll: &testPollStorage{}, Vote: &testVoteStorage{}}
	return &Controller{testStorage}
}

func TestHandlePollsGetOne(t *testing.T) {
	c := newTestController()

	tests := []struct {
		id             string
		expectedStatus int
	}{
		{"test-poll", http.StatusOK},
		{"non-existent-poll", http.StatusNotFound},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/polls/%v", tt.id), nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", tt.id)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		w := httptest.NewRecorder()

		c.handlePollsGetOne()(w, r)
		if tt.expectedStatus != w.Result().StatusCode {
			t.Errorf("expected response status code to be %v, instead got %v", tt.expectedStatus, w.Code)
		}
	}
}

func TestHandlePollsCreateOne(t *testing.T) {
	type request struct {
		Title   string   `json:"title"`
		Options []string `json:"options"`
	}

	c := newTestController()

	tests := []struct {
		requestBody    string
		expectedStatus int
	}{
		{
			`{ "title": "test", "options": ["option1"] }`,
			http.StatusCreated,
		},
		{
			`{ "title": "test", "options": [] }`,
			http.StatusBadRequest,
		},
		{
			`{ "title": "test" }`,
			http.StatusBadRequest,
		},
		{
			`{ "options": "test" }`,
			http.StatusBadRequest,
		},
		{
			`{ "options": ["test"] }`,
			http.StatusBadRequest,
		},
		{
			"",
			http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("POST", "/api/polls", strings.NewReader(tt.requestBody))
		w := httptest.NewRecorder()
		c.handlePollsCreateOne()(w, r)
		if w.Result().StatusCode != tt.expectedStatus {
			t.Errorf("expected status to be %v, instead got %v", tt.expectedStatus, w.Result().StatusCode)
		}
	}
}
