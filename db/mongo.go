package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JukeboxDB struct {
	client      *mongo.Client
	albumDao    *AlbumDAO
	artistDao   *ArtistDAO
	playlistDao *PlaylistDAO
	receiverDao *ReceiverDAO
	songDao     *SongDAO
}

func NewJukeboxDB(uri string) *JukeboxDB {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logrus.WithError(err).Panic("Failed in create mongo client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.WithError(err).Panic("Failed to connect to mongo")
	}

	mDb := client.Database("jukebox")

	return &JukeboxDB{
		client:      client,
		albumDao:    NewAlbumDAO(mDb),
		artistDao:   NewArtistDAO(mDb),
		playlistDao: NewPlaylistDAO(mDb),
		receiverDao: NewReceiverDAO(mDb),
		songDao:     NewSongDAO(mDb),
	}
}

func (db *JukeboxDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.client.Disconnect(ctx)
}

func (db *JukeboxDB) Albums() *AlbumDAO {
	return db.albumDao
}

func (db *JukeboxDB) Artists() *ArtistDAO {
	return db.artistDao
}

func (db *JukeboxDB) Playlists() *PlaylistDAO {
	return db.playlistDao
}

func (db *JukeboxDB) Receivers() *ReceiverDAO {
	return db.receiverDao
}

func (db *JukeboxDB) Songs() *SongDAO {
	return db.songDao
}
