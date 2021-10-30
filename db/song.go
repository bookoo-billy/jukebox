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

type SongDAO struct {
	collection *mongo.Collection
}

func NewSongDAO(mDb *mongo.Database) *SongDAO {
	collection := mDb.Collection("Songs")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}, {Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		logrus.WithError(err).Panic("Failed in create song index")
	}

	return &SongDAO{collection: collection}
}

func (a *SongDAO) Create(ctx context.Context, song *v1.Song) (*v1.Song, error) {
	err := a.validate(song)
	if err != nil {
		return nil, err
	}

	song.Id = uuid.New().String()
	_, err = a.collection.InsertOne(ctx, song)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (a *SongDAO) Delete(ctx context.Context, query *v1.SongQuery) (*v1.Song, error) {
	res := a.collection.FindOneAndDelete(ctx, query)

	Song := &v1.Song{}
	err := res.Decode(Song)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	return Song, nil
}

func (a *SongDAO) Get(ctx context.Context, query *v1.SongQuery) (*v1.Song, error) {
	res := a.collection.FindOne(ctx, query)

	Song := &v1.Song{}
	err := res.Decode(Song)
	if err != nil {
		return nil, err
	}

	return Song, nil
}

func (a *SongDAO) List(ctx context.Context, query *v1.SongQuery) (*v1.SongList, error) {
	cursor, err := a.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	Songs := &v1.SongList{Items: []*v1.Song{}}
	err = cursor.All(ctx, &Songs.Items)
	if err != nil {
		return nil, err
	}

	return Songs, nil
}

func (a *SongDAO) Update(ctx context.Context, song *v1.Song) (*v1.Song, error) {
	if song.Id == "" {
		return nil, errors.New("song.Id is a required field")
	}

	res := a.collection.FindOneAndUpdate(ctx, bson.M{"id": song.Id}, bson.M{"$set": song}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err := res.Decode(song)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (a *SongDAO) validate(song *v1.Song) error {
	if song.Artist.Id == "" {
		return errors.New("song.Artist.Id is a required field")
	}

	if song.Album.Id == "" {
		return errors.New("song.Album.Id is a required field")
	}

	if song.Name == "" {
		return errors.New("song.Name is a required field")
	}

	return nil
}
