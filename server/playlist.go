package server

import (
	"context"
	"github.com/bookoo-billy/jukebox/db"
	v1 "github.com/bookoo-billy/jukebox/gen/api/v1"
)

type PlaylistServer struct {
	v1.UnimplementedPlaylistServiceServer
	dao db.PlaylistDAO
}

func NewPlaylistServer(dao db.PlaylistDAO) v1.PlaylistServiceServer {
	return &PlaylistServer{dao: dao}
}

func (s *PlaylistServer) List(ctx context.Context, query *v1.PlaylistQuery) (*v1.PlaylistList, error) {
	return s.dao.List(ctx, query)
}

func (s *PlaylistServer) Create(ctx context.Context, request *v1.PlaylistCreateRequest) (*v1.Playlist, error) {
	return s.dao.Create(ctx, request.Playlist)
}

func (s *PlaylistServer) Get(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error) {
	return s.dao.Get(ctx, query)
}

func (s *PlaylistServer) Update(ctx context.Context, request *v1.PlaylistUpdateRequest) (*v1.Playlist, error) {
	return s.dao.Update(ctx, request.Playlist)
}

func (s *PlaylistServer) Delete(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error) {
	return s.dao.Delete(ctx, query)
}
