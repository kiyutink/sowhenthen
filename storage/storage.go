package storage

import (
	"context"

	"github.com/kiyutink/sowhenthen/entities"
)

type Poll interface {
	Create(ctx context.Context, p entities.Poll) (entities.Poll, error)
	GetOne(ctx context.Context, id string) (entities.Poll, error)
}

type Vote interface {
	Create(ctx context.Context, vote entities.Vote) (entities.Vote, error)
	GetMany(ctx context.Context, pollId string) ([]entities.Vote, error)
}

type Storage struct {
	Poll Poll
	Vote Vote
}
