package poll

import (
	"context"
	"strconv"
)

type MemoryStorer struct {
	storage map[string]Poll
	last    int
}

func NewMemeoryStorer() *MemoryStorer {
	return &MemoryStorer{
		storage: map[string]Poll{},
	}
}

func (mpc *MemoryStorer) GetOne(ctx context.Context, id string) (Poll, error) {
	poll := mpc.storage[id]
	return poll, nil
}

func (mpc *MemoryStorer) Create(ctx context.Context, p Poll) (Poll, error) {
	mpc.last++
	p.Id = strconv.Itoa(mpc.last)
	mpc.storage[p.Id] = p
	return p, nil
}

func (mpc *MemoryStorer) Delete(ctx context.Context, id string) error {
	delete(mpc.storage, id)
	return nil
}

func (mpc *MemoryStorer) GetMany(ctx context.Context) ([]Poll, error) {
	polls := []Poll{}
	for _, poll := range mpc.storage {
		polls = append(polls, poll)
	}
	return polls, nil
}

func (mpc *MemoryStorer) Dump() interface{} {
	return mpc.storage
}
