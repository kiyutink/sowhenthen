package mongo

import (
	"context"
	"fmt"

	"github.com/kiyutink/sowhenthen/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type pollModel struct {
	Id      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"title"`
	Options []string           `bson:"options"`
}

type pollStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func newPolLStorage(client *mongo.Client) *pollStorage {
	return &pollStorage{client, client.Database(databaseName).Collection(pollsCollectionName)}
}

func (ps *pollStorage) GetOne(ctx context.Context, id string) (entities.Poll, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Poll{}, fmt.Errorf("invalid id: %w", err)
	}
	res := ps.collection.FindOne(ctx, bson.M{"_id": objId})
	err = res.Err()
	if err != nil {
		return entities.Poll{}, fmt.Errorf("error looking up poll: %w", err)
	}

	pm := pollModel{}
	err = res.Decode(&pm)
	if err != nil {
		return entities.Poll{}, fmt.Errorf("error decoding poll: %w", err)
	}

	poll := entities.Poll{}
	poll.Id = pm.Id.Hex()
	poll.Options = pm.Options
	poll.Title = pm.Title

	return poll, nil
}

func (ps *pollStorage) Create(ctx context.Context, p entities.Poll) (entities.Poll, error) {
	pm := pollModel{
		Id:      primitive.NewObjectID(),
		Title:   p.Title,
		Options: p.Options,
	}

	insertRes, err := ps.collection.InsertOne(ctx, pm)
	if err != nil {
		return p, fmt.Errorf("error creating poll: %w", err)
	}
	p.Id = insertRes.InsertedID.(primitive.ObjectID).Hex()
	return p, nil
}

func (ps *pollStorage) Dump() interface{} {
	return nil
}
