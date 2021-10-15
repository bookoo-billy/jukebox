package db

import (
	v1 "github.com/bookoo-billy/jukebox/api/v1"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaylistDAO struct {
	collection *mongo.Collection
}

func NewPlaylistDAO(mDb *mongo.Database) *PlaylistDAO {
	collection := mDb.Collection("playlists")
	return &PlaylistDAO{collection: collection}
}

func (a *PlaylistDAO) Get(query *v1.PlaylistQuery) (*v1.Playlist, error) {
	return &v1.Playlist{}, nil
}
