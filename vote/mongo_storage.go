package vote

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "sowhenthen"
	collectionName = "votes"
)

type mongoModel struct {
	PollId    primitive.ObjectID `bson:"pollId"`
	Options   []string           `bson:"options"`
	VoterName string             `bson:"voterName"`
}

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStorage(client *mongo.Client) *MongoStorage {
	return &MongoStorage{client, client.Database(databaseName).Collection(collectionName)}
}

func (ms *MongoStorage) Create(ctx context.Context, vote Vote) (Vote, error) {
	pollIdObjId, err := primitive.ObjectIDFromHex(vote.PollId)
	if err != nil {
		return Vote{}, fmt.Errorf("error converting pollId to objectId: %w", err)
	}
	model := mongoModel{
		PollId:    pollIdObjId,
		Options:   vote.Options,
		VoterName: vote.VoterName,
	}
	_, err = ms.collection.InsertOne(ctx, model)
	if err != nil {
		return Vote{}, fmt.Errorf("error saving vote: %w", err)
	}

	return vote, nil
}

func (ms *MongoStorage) GetMany(ctx context.Context, pollId string) ([]Vote, error) {
	pollIdObjId, err := primitive.ObjectIDFromHex(pollId)
	if err != nil {
		return nil, fmt.Errorf("error creating objectId: %w", err)
	}
	res, err := ms.collection.Find(ctx, bson.M{"pollId": pollIdObjId})
	if err != nil {
		return nil, fmt.Errorf("error getting votes: %w", err)
	}

	vm := []mongoModel{{}}
	err = res.All(ctx, &vm)
	if err != nil {
		return nil, fmt.Errorf("error getting votes: %w", err)
	}

	votes := make([]Vote, len(vm))

	for i, vote := range vm {
		v := Vote{
			PollId:    vote.PollId.Hex(),
			Options:   vote.Options,
			VoterName: vote.VoterName,
		}
		votes[i] = v
	}

	return votes, nil
}
