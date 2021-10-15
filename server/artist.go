package server

import (
	"context"

	"github.com/bookoo-billy/jukebox/db"
	v1 "github.com/bookoo-billy/jukebox/gen/api/v1"
)

type ArtistServer struct {
	v1.UnimplementedArtistServiceServer
	dao *db.ArtistDAO
}

func NewArtistServer(dao *db.ArtistDAO) v1.ArtistServiceServer {
	return &ArtistServer{dao: dao}
}

func (s *ArtistServer) List(ctx context.Context, query *v1.ArtistQuery) (*v1.ArtistList, error) {
	return s.dao.List(ctx, query)
}

func (s *ArtistServer) Create(ctx context.Context, request *v1.ArtistCreateRequest) (*v1.Artist, error) {
	return s.dao.Create(ctx, request.Artist)
}

func (s *ArtistServer) Get(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error) {
	return s.dao.Get(ctx, query)
}

func (s *ArtistServer) Update(ctx context.Context, request *v1.ArtistUpdateRequest) (*v1.Artist, error) {
	return s.dao.Update(ctx, request.Artist)
}

func (s *ArtistServer) Delete(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error) {
	return s.dao.Delete(ctx, query)
}
