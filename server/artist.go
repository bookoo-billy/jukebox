package server

import (
	"context"

	v1 "github.com/bookoo-billy/jukebox/api/v1"
	"github.com/bookoo-billy/jukebox/db"
)

type ArtistServer struct {
	v1.UnimplementedArtistServiceServer
	dao *db.ArtistDAO
}

func NewArtistServer(dao *db.ArtistDAO) v1.ArtistServiceServer {
	return &ArtistServer{dao: dao}
}

func (p *ArtistServer) List(context.Context, *v1.ArtistQuery) (*v1.ArtistList, error) {
	panic("implement me")
}

func (p *ArtistServer) Create(context.Context, *v1.ArtistCreateRequest) (*v1.Artist, error) {
	panic("implement me")
}

func (p *ArtistServer) Get(context.Context, *v1.ArtistQuery) (*v1.Artist, error) {
	panic("implement me")
}

func (p *ArtistServer) Update(context.Context, *v1.ArtistUpdateRequest) (*v1.Artist, error) {
	panic("implement me")
}

func (p *ArtistServer) Delete(context.Context, *v1.ArtistQuery) (*v1.Artist, error) {
	panic("implement me")
}
