package postgres

import (
	"github.com/bookoo-billy/jukebox/db"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type JukeboxDb struct {
	client      *gorm.DB
	albumDao    db.AlbumDAO
	artistDao   db.ArtistDAO
	playlistDao db.PlaylistDAO
	receiverDao db.ReceiverDAO
	songDao     db.SongDAO
}

func NewJukeboxDB(uri string) db.JukeboxDb {
	client, err := gorm.Open(postgres.New(postgres.Config{
		DSN: uri,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Panic("Failed in create postgres client")
	}

	return &JukeboxDb{
		albumDao:    NewAlbumDAO(client),
		artistDao:   NewArtistDAO(client),
		playlistDao: NewPlaylistDAO(client),
		receiverDao: NewReceiverDAO(client),
		songDao:     NewSongDAO(client),
	}
}

func (db *JukeboxDb) Close() error {
	return nil
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
