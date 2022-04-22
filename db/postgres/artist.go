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

type ArtistDAO struct {
	db *gorm.DB
}

func NewArtistDAO(db *gorm.DB) db.ArtistDAO {
	return &ArtistDAO{db: db}
}

func (a *ArtistDAO) Create(ctx context.Context, artist *v1.Artist) (*v1.Artist, error) {
	err := a.validate(artist)
	if err != nil {
		return nil, err
	}

	artist.Id = uuid.New().String()
	tx := a.db.Create(artist)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return artist, nil
}

func (a *ArtistDAO) Delete(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error) {
	artist := &v1.Artist{}
	tx := a.db.Model(artist).
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Delete(v1.Artist{Id: query.Id})

	if tx.Error != nil {
		return nil, tx.Error
	}

	return artist, nil
}

func (a *ArtistDAO) Get(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error) {
	artist := &v1.Artist{
		Id:     query.Id,
		Name:   query.Name,
	}
	tx := a.db.WithContext(ctx).First(&artist)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return artist, nil
}

func (a *ArtistDAO) List(ctx context.Context, query *v1.ArtistQuery) (*v1.ArtistList, error) {
	var items []*v1.Artist
	tx := a.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where(&v1.Artist{
			Id:     query.Id,
			Name:   query.Name,
		}).
		Find(items)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v1.ArtistList{Items: items}, nil
}

func (a *ArtistDAO) Update(ctx context.Context, artist *v1.Artist) (*v1.Artist, error) {
	if artist.Id == "" {
		return nil, errors.New("artist.Id is a required field")
	}

	updated := &v1.Artist{}
	tx := a.db.Model(updated).WithContext(ctx).Clauses(clause.Returning{}).Updates(artist)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return artist, nil
}

func (a *ArtistDAO) validate(artist *v1.Artist) error {
	if artist.Name == "" {
		return errors.New("artist.Name is a required field")
	}

	return nil
}
