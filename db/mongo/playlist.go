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

type PlaylistDAO struct {
	collection *mongo.Collection
}

func NewPlaylistDAO(mDb *mongo.Database) db.PlaylistDAO {
	collection := mDb.Collection("Playlists")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}, {Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		logrus.WithError(err).Panic("Failed in create playlist index")
	}

	return &PlaylistDAO{collection: collection}
}

func (a *PlaylistDAO) Create(ctx context.Context, playlist *v1.Playlist) (*v1.Playlist, error) {
	err := a.validate(playlist)
	if err != nil {
		return nil, err
	}

	playlist.Id = uuid.New().String()
	_, err = a.collection.InsertOne(ctx, playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (a *PlaylistDAO) Delete(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error) {
	res := a.collection.FindOneAndDelete(ctx, query)

	Playlist := &v1.Playlist{}
	err := res.Decode(Playlist)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	return Playlist, nil
}

func (a *PlaylistDAO) Get(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error) {
	res := a.collection.FindOne(ctx, query)

	Playlist := &v1.Playlist{}
	err := res.Decode(Playlist)
	if err != nil {
		return nil, err
	}

	return Playlist, nil
}

func (a *PlaylistDAO) List(ctx context.Context, query *v1.PlaylistQuery) (*v1.PlaylistList, error) {
	cursor, err := a.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	Playlists := &v1.PlaylistList{Items: []*v1.Playlist{}}
	err = cursor.All(ctx, &Playlists.Items)
	if err != nil {
		return nil, err
	}

	return Playlists, nil
}

func (a *PlaylistDAO) Update(ctx context.Context, playlist *v1.Playlist) (*v1.Playlist, error) {
	if playlist.Id == "" {
		return nil, errors.New("Playlist.Id is a required field")
	}

	res := a.collection.FindOneAndUpdate(ctx, bson.M{"id": playlist.Id}, bson.M{"$set": playlist}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err := res.Decode(playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (a *PlaylistDAO) validate(playlist *v1.Playlist) error {
	if playlist.Name == "" {
		return errors.New("playlist.Name is a required field")
	}

	return nil
}
