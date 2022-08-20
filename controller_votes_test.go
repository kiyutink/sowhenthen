package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/kiyutink/sowhenthen/entities"
	"github.com/kiyutink/sowhenthen/storage"
)

type testVoteStorage struct {
	storage.Vote
}

func (tvs *testVoteStorage) GetMany(ctx context.Context, pollId string) ([]entities.Vote, error) {
	if pollId == "test-id" {
		return []entities.Vote{{PollId: "test-id", Options: []string{"test-option-1"}}}, nil
	}

	return nil, &storage.NotFoundError{Identifier: pollId, Err: errors.New("poll doesn't exist")}
}

func (tvs *testVoteStorage) CreateOne(ctx context.Context, v entities.Vote) (entities.Vote, error) {
	if v.PollId == "test-id" {
		return v, nil
	}
	return entities.Vote{}, &storage.InvalidRequestError{Message: "poll doesn't exist"}
}

func TestHandleVotesGetMany(t *testing.T) {
	tests := []struct {
		pollId         string
		expectedStatus int
	}{
		{"test-id", http.StatusOK},
		{"invalid-id", http.StatusNotFound},
	}

	c := newTestController()

	for _, tt := range tests {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/polls/%v/votes", tt.pollId), nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("pollId", tt.pollId)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		w := httptest.NewRecorder()

		c.handleVotesGetMany()(w, r)

		if w.Result().StatusCode != tt.expectedStatus {
			t.Errorf("expected status code to be %v, instead got %v", tt.expectedStatus, w.Result().StatusCode)
		}
	}
}

func TestHandleVotesCreateOne(t *testing.T) {
}
