package server

import (
	"context"

	v1 "github.com/bookoo-billy/jukebox/api/v1"
	"github.com/bookoo-billy/jukebox/db"
)

type PlaylistServer struct {
	v1.UnimplementedPlaylistServiceServer
	dao *db.PlaylistDAO
}

func NewPlaylistServer(dao *db.PlaylistDAO) v1.PlaylistServiceServer {
	return &PlaylistServer{dao: dao}
}

func (p *PlaylistServer) List(context.Context, *v1.PlaylistQuery) (*v1.PlaylistList, error) {
	panic("implement me")
}

func (p *PlaylistServer) Create(context.Context, *v1.PlaylistCreateRequest) (*v1.Playlist, error) {
	panic("implement me")
}

func (p *PlaylistServer) Get(context.Context, *v1.PlaylistQuery) (*v1.Playlist, error) {
	panic("implement me")
}

func (p *PlaylistServer) Update(context.Context, *v1.PlaylistUpdateRequest) (*v1.Playlist, error) {
	panic("implement me")
}

func (p *PlaylistServer) Delete(context.Context, *v1.PlaylistQuery) (*v1.Playlist, error) {
	panic("implement me")
}
