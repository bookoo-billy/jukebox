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

type AlbumDAO struct {
	collection *mongo.Collection
}

func NewAlbumDAO(mDb *mongo.Database) db.AlbumDAO {
	collection := mDb.Collection("albums")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}, {Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		logrus.WithError(err).Panic("Failed in create album index")
	}

	return &AlbumDAO{collection: collection}
}

func (a *AlbumDAO) Create(ctx context.Context, album *v1.Album) (*v1.Album, error) {
	err := a.validate(album)
	if err != nil {
		return nil, err
	}

	album.Id = uuid.New().String()
	_, err = a.collection.InsertOne(ctx, album)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (a *AlbumDAO) Delete(ctx context.Context, query *v1.AlbumQuery) (*v1.Album, error) {
	res := a.collection.FindOneAndDelete(ctx, query)

	album := &v1.Album{}
	err := res.Decode(album)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	return album, nil
}

func (a *AlbumDAO) Get(ctx context.Context, query *v1.AlbumQuery) (*v1.Album, error) {
	res := a.collection.FindOne(ctx, query)

	album := &v1.Album{}
	err := res.Decode(album)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (a *AlbumDAO) List(ctx context.Context, query *v1.AlbumQuery) (*v1.AlbumList, error) {
	cursor, err := a.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	albums := &v1.AlbumList{Items: []*v1.Album{}}
	err = cursor.All(ctx, &albums.Items)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (a *AlbumDAO) Update(ctx context.Context, album *v1.Album) (*v1.Album, error) {
	if album.Id == "" {
		return nil, errors.New("album.Id is a required field")
	}

	res := a.collection.FindOneAndUpdate(ctx, bson.M{"id": album.Id}, bson.M{"$set": album}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err := res.Decode(album)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (a *AlbumDAO) validate(album *v1.Album) error {
	if album.Artist == nil {
		return errors.New("album.Artist is a required field")
	}

	if album.Artist.Id == "" {
		return errors.New("album.Artist.Id is a required field")
	}

	return nil
}
