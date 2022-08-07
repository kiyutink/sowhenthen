package main

import (
	"context"

	"github.com/kiyutink/sowhenthen/poll"
	"github.com/kiyutink/sowhenthen/vote"
)

type PollStorage interface {
	Create(ctx context.Context, p poll.Poll) (poll.Poll, error)
	GetOne(ctx context.Context, id string) (poll.Poll, error)
}

type VoteStorage interface {
	Create(ctx context.Context, vote vote.Vote) (vote.Vote, error)
	GetMany(ctx context.Context, pollId string) ([]vote.Vote, error)
}

type Storage struct {
	poll PollStorage
	vote VoteStorage
}

type Controller struct {
	storage Storage
}

func NewController(ps PollStorage, vs VoteStorage) *Controller {
	return &Controller{
		storage: Storage{poll: ps, vote: vs},
	}
}
