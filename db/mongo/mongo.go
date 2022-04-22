package mongo

import (
	"context"
	"github.com/bookoo-billy/jukebox/db"
	"github.com/sirupsen/logrus"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JukeboxDb struct {
	client      *mongo.Client
	albumDao    db.AlbumDAO
	artistDao   db.ArtistDAO
	playlistDao db.PlaylistDAO
	receiverDao db.ReceiverDAO
	songDao     db.SongDAO
}

func NewJukeboxDB(uri string) db.JukeboxDb {
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

	return &JukeboxDb{
		client:      client,
		albumDao:    NewAlbumDAO(mDb),
		artistDao:   NewArtistDAO(mDb),
		playlistDao: NewPlaylistDAO(mDb),
		receiverDao: NewReceiverDAO(mDb),
		songDao:     NewSongDAO(mDb),
	}
}

func (db *JukeboxDb) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.client.Disconnect(ctx)
}

func (db *JukeboxDb) Albums() db.AlbumDAO {
	return db.albumDao
}

func (db *JukeboxDb) Artists() db.ArtistDAO {
	return db.artistDao
}

func (db *JukeboxDb) Playlists() db.PlaylistDAO {
	return db.playlistDao
}

func (db *JukeboxDb) Receivers() db.ReceiverDAO {
	return db.receiverDao
}

func (db *JukeboxDb) Songs() db.SongDAO {
	return db.songDao
}
