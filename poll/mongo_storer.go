package poll

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "sowhenthen"
	collectionName = "polls"
)

type mongoModel struct {
	Id      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"title"`
	Options []string           `bson:"options"`
}

type MongoStorer struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStorer(client *mongo.Client) *MongoStorer {
	return &MongoStorer{client, client.Database(databaseName).Collection(collectionName)}
}

func (ms *MongoStorer) GetOne(ctx context.Context, id string) (Poll, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Poll{}, fmt.Errorf("invalid id: %w", err)
	}
	res := ms.collection.FindOne(ctx, bson.M{"_id": objId})
	err = res.Err()
	if err != nil {
		return Poll{}, fmt.Errorf("error looking up poll: %w", err)
	}

	pm := mongoModel{}
	err = res.Decode(&pm)
	if err != nil {
		return Poll{}, fmt.Errorf("error decoding poll: %w", err)
	}

	poll := Poll{}
	poll.Id = pm.Id.Hex()
	poll.Options = pm.Options
	poll.Title = pm.Title

	return poll, nil
}

func (ms *MongoStorer) Create(ctx context.Context, p Poll) (Poll, error) {
	pm := mongoModel{
		Id:      primitive.NewObjectID(),
		Title:   p.Title,
		Options: p.Options,
	}

	insertRes, err := ms.collection.InsertOne(ctx, pm)
	if err != nil {
		return p, fmt.Errorf("error creating poll: %w", err)
	}
	p.Id = insertRes.InsertedID.(primitive.ObjectID).Hex()
	return p, nil
}

func (ms *MongoStorer) Dump() interface{} {
	return nil
}
