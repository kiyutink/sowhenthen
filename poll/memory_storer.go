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

func (ms *MemoryStorer) GetOne(ctx context.Context, id string) (Poll, error) {
	poll := ms.storage[id]
	return poll, nil
}

func (ms *MemoryStorer) Create(ctx context.Context, p Poll) (Poll, error) {
	ms.last++
	p.Id = strconv.Itoa(ms.last)
	ms.storage[p.Id] = p
	return p, nil
}

func (ms *MemoryStorer) Delete(ctx context.Context, id string) error {
	delete(ms.storage, id)
	return nil
}

func (ms *MemoryStorer) GetMany(ctx context.Context) ([]Poll, error) {
	polls := []Poll{}
	for _, poll := range ms.storage {
		polls = append(polls, poll)
	}
	return polls, nil
}

func (ms *MemoryStorer) Dump() interface{} {
	return ms.storage
}
