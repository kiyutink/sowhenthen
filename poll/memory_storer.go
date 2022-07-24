package poll

type MemoryStorer struct {
	storage map[int]Poll
	last    int
}

func NewMemeoryStorer() *MemoryStorer {
	return &MemoryStorer{
		storage: map[int]Poll{},
	}
}

func (mpc *MemoryStorer) GetOne(id int) Poll {
	poll := mpc.storage[id]
	return poll
}

func (mpc *MemoryStorer) Create(p Poll) {
	mpc.last++
	p.Id = mpc.last
	mpc.storage[p.Id] = p
}

func (mpc *MemoryStorer) Delete(id int) {
	delete(mpc.storage, id)
}

func (mpc *MemoryStorer) GetMany() []Poll {
	polls := []Poll{}
	for _, poll := range mpc.storage {
		polls = append(polls, poll)
	}
	return polls
}

func (mpc *MemoryStorer) Dump() interface{} {
	return mpc.storage
}
