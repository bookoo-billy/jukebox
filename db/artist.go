package db

import (
	"context"

	v1 "github.com/bookoo-billy/jukebox/api/v1"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArtistDAO struct {
	collection *mongo.Collection
}

func NewArtistDAO(mDb *mongo.Database) *ArtistDAO {
	collection := mDb.Collection("artists")
	return &ArtistDAO{collection: collection}
}

func (a *ArtistDAO) Get(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error) {
	res := a.collection.FindOne(ctx, query)

	artist := &v1.Artist{}
	res.Decode(artist)
	return artist, nil
}

func (a *ArtistDAO) Create(ctx context.Context, artist *v1.Artist) (*v1.Artist, error) {
	artist.Id = uuid.New().String()
	_, err := a.collection.InsertOne(ctx, artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
