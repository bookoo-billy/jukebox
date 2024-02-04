package postgres

import (
	"context"
	"errors"

	"github.com/bookoo-billy/jukebox/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	v1 "github.com/bookoo-billy/jukebox/proto/api/v1"
	"github.com/google/uuid"
)

type PlaylistDAO struct {
	db *gorm.DB
}

func NewPlaylistDAO(db *gorm.DB) db.PlaylistDAO {
	return &PlaylistDAO{db: db}
}

func (a *PlaylistDAO) Create(ctx context.Context, playlist *v1.Playlist) (*v1.Playlist, error) {
	err := a.validate(playlist)
	if err != nil {
		return nil, err
	}

	playlist.Id = uuid.New().String()
	tx := a.db.WithContext(ctx).Create(playlist)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return playlist, nil
}

func (a *PlaylistDAO) Delete(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error) {
	playlist := &v1.Playlist{}
	tx := a.db.Model(playlist).
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Delete(v1.Playlist{Id: query.Id})

	if tx.Error != nil {
		return nil, tx.Error
	}

	return playlist, nil
}

func (a *PlaylistDAO) Get(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error) {
	playlist := &v1.Playlist{}
	tx := a.db.Model(playlist).WithContext(ctx).First(query)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return playlist, nil
}

func (a *PlaylistDAO) List(ctx context.Context, query *v1.PlaylistQuery) (*v1.PlaylistList, error) {
	var items []*v1.Playlist
	tx := a.db.Model(items).WithContext(ctx).Find(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v1.PlaylistList{Items: items}, nil
}

func (a *PlaylistDAO) Update(ctx context.Context, playlist *v1.Playlist) (*v1.Playlist, error) {
	if playlist.Id == "" {
		return nil, errors.New("Playlist.Id is a required field")
	}

	updated := &v1.Playlist{}
	tx := a.db.Model(updated).WithContext(ctx).Clauses(clause.Returning{}).Updates(playlist)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return updated, nil
}

func (a *PlaylistDAO) validate(playlist *v1.Playlist) error {
	if playlist.Name == "" {
		return errors.New("playlist.Name is a required field")
	}

	return nil
}
