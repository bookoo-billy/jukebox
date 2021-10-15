package db

import (
	v1 "github.com/bookoo-billy/jukebox/api/v1"
	"go.mongodb.org/mongo-driver/mongo"
)

type SongDAO struct {
	collection *mongo.Collection
}

func NewSongDAO(mDb *mongo.Database) *SongDAO {
	collection := mDb.Collection("songs")
	return &SongDAO{collection: collection}
}

func (a *SongDAO) Get(query *v1.SongQuery) (*v1.Song, error) {
	return &v1.Song{}, nil
}
