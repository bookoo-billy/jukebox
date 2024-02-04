package postgres

import (
	"context"
	"errors"

	"github.com/bookoo-billy/jukebox/db"
	v1 "github.com/bookoo-billy/jukebox/proto/api/v1"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AlbumDAO struct {
	db *gorm.DB
}

func NewAlbumDAO(db *gorm.DB) db.AlbumDAO {
	return &AlbumDAO{
		db: db,
	}
}

func (a *AlbumDAO) Create(ctx context.Context, album *v1.Album) (*v1.Album, error) {
	err := a.validate(album)
	if err != nil {
		return nil, err
	}

	album.Id = uuid.New().String()
	tx := a.db.WithContext(ctx).Create(album)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return album, nil
}

func (a *AlbumDAO) Delete(ctx context.Context, query *v1.AlbumQuery) (*v1.Album, error) {
	album := &v1.Album{}
	tx := a.db.Model(album).WithContext(ctx).Clauses(clause.Returning{}).Delete(v1.Album{Id: query.Id})
	if tx.Error != nil {
		return nil, tx.Error
	}

	return album, nil
}

func (a *AlbumDAO) Get(ctx context.Context, query *v1.AlbumQuery) (*v1.Album, error) {
	album := &v1.Album{}
	tx := a.db.Model(album).WithContext(ctx).First(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return album, nil
}

func (a *AlbumDAO) List(ctx context.Context, query *v1.AlbumQuery) (*v1.AlbumList, error) {
	var items []*v1.Album
	tx := a.db.WithContext(ctx).
		Where(v1.Album{
			Id:   query.Id,
			Name: query.Name,
			Artist: &v1.Artist{
				Id:   query.ArtistId,
				Name: query.ArtistName,
			},
		}).
		Find(items)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v1.AlbumList{Items: items}, nil
}

func (a *AlbumDAO) Update(ctx context.Context, album *v1.Album) (*v1.Album, error) {
	if album.Id == "" {
		return nil, errors.New("album.Id is a required field")
	}

	updated := &v1.Album{}
	tx := a.db.Model(updated).WithContext(ctx).Clauses(clause.Returning{}).Updates(album)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return updated, nil
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
