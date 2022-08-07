package mongo

import (
	"github.com/kiyutink/sowhenthen/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName        = "sowhenthen"
	pollsCollectionName = "polls"
	votesCollectionName = "votes"
)

func NewStorage(client *mongo.Client) storage.Storage {
	ps := newPolLStorage(client)
	vs := newVoteStorage(client)

	return storage.Storage{Poll: ps, Vote: vs}
}
