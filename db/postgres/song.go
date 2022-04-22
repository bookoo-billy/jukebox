package postgres

import (
	"context"
	"errors"
	"github.com/bookoo-billy/jukebox/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	v1 "github.com/bookoo-billy/jukebox/gen/api/v1"
	"github.com/google/uuid"
)

type SongDAO struct {
	db *gorm.DB
}

func NewSongDAO(db *gorm.DB) db.SongDAO {
	return &SongDAO{db: db}
}

func (a *SongDAO) Create(ctx context.Context, song *v1.Song) (*v1.Song, error) {
	err := a.validate(song)
	if err != nil {
		return nil, err
	}

	song.Id = uuid.New().String()
	tx := a.db.WithContext(ctx).Create(song)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return song, nil
}

func (a *SongDAO) Delete(ctx context.Context, query *v1.SongQuery) (*v1.Song, error) {
	deleted := &v1.Song{}
	tx := a.db.Model(deleted).WithContext(ctx).Clauses(clause.Returning{}).Delete(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return deleted, nil
}

func (a *SongDAO) Get(ctx context.Context, query *v1.SongQuery) (*v1.Song, error) {
	song := &v1.Song{}
	tx := a.db.Model(song).WithContext(ctx).First(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return song, nil
}

func (a *SongDAO) List(ctx context.Context, query *v1.SongQuery) (*v1.SongList, error) {
	var items []*v1.Song
	tx := a.db.Model(items).WithContext(ctx).Find(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v1.SongList{Items: items}, nil
}

func (a *SongDAO) Update(ctx context.Context, song *v1.Song) (*v1.Song, error) {
	if song.Id == "" {
		return nil, errors.New("song.Id is a required field")
	}

	updated := &v1.Song{}
	tx := a.db.Model(updated).WithContext(ctx).Clauses(clause.Returning{}).Updates(song)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return updated, nil
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
