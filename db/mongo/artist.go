package mongo

import (
	"context"
	"errors"

	"github.com/bookoo-billy/jukebox/db"
	"github.com/sirupsen/logrus"

	v1 "github.com/bookoo-billy/jukebox/proto/api/v1"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArtistDAO struct {
	collection *mongo.Collection
}

func NewArtistDAO(mDb *mongo.Database) db.ArtistDAO {
	collection := mDb.Collection("Artists")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}, {Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		logrus.WithError(err).Panic("Failed in create artist index")
	}

	return &ArtistDAO{collection: collection}
}

func (a *ArtistDAO) Create(ctx context.Context, artist *v1.Artist) (*v1.Artist, error) {
	err := a.validate(artist)
	if err != nil {
		return nil, err
	}

	artist.Id = uuid.New().String()
	_, err = a.collection.InsertOne(ctx, artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (a *ArtistDAO) Delete(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error) {
	res := a.collection.FindOneAndDelete(ctx, query)

	Artist := &v1.Artist{}
	err := res.Decode(Artist)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	return Artist, nil
}

func (a *ArtistDAO) Get(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error) {
	res := a.collection.FindOne(ctx, query)

	Artist := &v1.Artist{}
	err := res.Decode(Artist)
	if err != nil {
		return nil, err
	}

	return Artist, nil
}

func (a *ArtistDAO) List(ctx context.Context, query *v1.ArtistQuery) (*v1.ArtistList, error) {
	cursor, err := a.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	Artists := &v1.ArtistList{Items: []*v1.Artist{}}
	err = cursor.All(ctx, &Artists.Items)
	if err != nil {
		return nil, err
	}

	return Artists, nil
}

func (a *ArtistDAO) Update(ctx context.Context, artist *v1.Artist) (*v1.Artist, error) {
	if artist.Id == "" {
		return nil, errors.New("artist.Id is a required field")
	}

	res := a.collection.FindOneAndUpdate(ctx, bson.M{"id": artist.Id}, bson.M{"$set": artist}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err := res.Decode(artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (a *ArtistDAO) validate(artist *v1.Artist) error {
	if artist.Name == "" {
		return errors.New("artist.Name is a required field")
	}

	return nil
}
