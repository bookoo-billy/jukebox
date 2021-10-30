package db

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"

	v1 "github.com/bookoo-billy/jukebox/gen/api/v1"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReceiverDAO struct {
	collection *mongo.Collection
}

func NewReceiverDAO(mDb *mongo.Database) *ReceiverDAO {
	collection := mDb.Collection("Receivers")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}, {Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		logrus.WithError(err).Panic("Failed in create receiver index")
	}

	return &ReceiverDAO{collection: collection}
}

func (a *ReceiverDAO) Create(ctx context.Context, receiver *v1.Receiver) (*v1.Receiver, error) {
	err := a.validate(receiver)
	if err != nil {
		return nil, err
	}

	receiver.Id = uuid.New().String()
	_, err = a.collection.InsertOne(ctx, receiver)
	if err != nil {
		return nil, err
	}

	return receiver, nil
}

func (a *ReceiverDAO) Delete(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error) {
	res := a.collection.FindOneAndDelete(ctx, query)

	Receiver := &v1.Receiver{}
	err := res.Decode(Receiver)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	return Receiver, nil
}

func (a *ReceiverDAO) Get(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error) {
	res := a.collection.FindOne(ctx, query)

	Receiver := &v1.Receiver{}
	err := res.Decode(Receiver)
	if err != nil {
		return nil, err
	}

	return Receiver, nil
}

func (a *ReceiverDAO) List(ctx context.Context, query *v1.ReceiverQuery) (*v1.ReceiverList, error) {
	cursor, err := a.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	Receivers := &v1.ReceiverList{Items: []*v1.Receiver{}}
	err = cursor.All(ctx, &Receivers.Items)
	if err != nil {
		return nil, err
	}

	return Receivers, nil
}

func (a *ReceiverDAO) Update(ctx context.Context, receiver *v1.Receiver) (*v1.Receiver, error) {
	if receiver.Id == "" {
		return nil, errors.New("receiver.Id is a required field")
	}

	res := a.collection.FindOneAndUpdate(ctx, bson.M{"id": receiver.Id}, bson.M{"$set": receiver}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err := res.Decode(receiver)
	if err != nil {
		return nil, err
	}

	return receiver, nil
}

func (a *ReceiverDAO) validate(receiver *v1.Receiver) error {
	if receiver.Name == "" {
		return errors.New("receiver.Name is a required field")
	}

	return nil
}
