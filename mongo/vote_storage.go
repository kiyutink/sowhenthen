package mongo

import (
	"context"
	"fmt"

	"github.com/kiyutink/sowhenthen/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type voteModel struct {
	PollId    primitive.ObjectID `bson:"pollId"`
	Options   []string           `bson:"options"`
	VoterName string             `bson:"voterName"`
}

type voteStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func newVoteStorage(client *mongo.Client) *voteStorage {
	return &voteStorage{client, client.Database(databaseName).Collection(votesCollectionName)}
}

func (ms *voteStorage) Create(ctx context.Context, vote entities.Vote) (entities.Vote, error) {
	pollIdObjId, err := primitive.ObjectIDFromHex(vote.PollId)
	if err != nil {
		return entities.Vote{}, fmt.Errorf("error converting pollId to objectId: %w", err)
	}
	model := voteModel{
		PollId:    pollIdObjId,
		Options:   vote.Options,
		VoterName: vote.VoterName,
	}
	_, err = ms.collection.InsertOne(ctx, model)
	if err != nil {
		return entities.Vote{}, fmt.Errorf("error saving vote: %w", err)
	}

	return vote, nil
}

func (ms *voteStorage) GetMany(ctx context.Context, pollId string) ([]entities.Vote, error) {
	pollIdObjId, err := primitive.ObjectIDFromHex(pollId)
	if err != nil {
		return nil, fmt.Errorf("error creating objectId: %w", err)
	}
	res, err := ms.collection.Find(ctx, bson.M{"pollId": pollIdObjId})
	if err != nil {
		return nil, fmt.Errorf("error getting votes: %w", err)
	}

	vm := []voteModel{{}}
	err = res.All(ctx, &vm)
	if err != nil {
		return nil, fmt.Errorf("error getting votes: %w", err)
	}

	votes := make([]entities.Vote, len(vm))

	for i, vote := range vm {
		v := entities.Vote{
			PollId:    vote.PollId.Hex(),
			Options:   vote.Options,
			VoterName: vote.VoterName,
		}
		votes[i] = v
	}

	return votes, nil
}
