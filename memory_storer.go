package main

type MemoryPollStorer struct {
	storage map[int]Poll
	last    int
}

func NewMemoryPollStorer() *MemoryPollStorer {
	return &MemoryPollStorer{
		storage: map[int]Poll{},
	}
}

func (mpc *MemoryPollStorer) GetOne(id int) Poll {
	poll := mpc.storage[id]
	return poll
}

func (mpc *MemoryPollStorer) Create(p Poll) {
	mpc.last++
	p.Id = mpc.last
	mpc.storage[p.Id] = p
}

func (mpc *MemoryPollStorer) Delete(id int) {
	delete(mpc.storage, id)
}

func (mpc *MemoryPollStorer) GetMany() []Poll {
	polls := []Poll{}
	for _, poll := range mpc.storage {
		polls = append(polls, poll)
	}
	return polls
}

func (mpc *MemoryPollStorer) Dump() interface{} {
	return mpc.storage
}
