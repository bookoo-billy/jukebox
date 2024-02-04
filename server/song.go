package server

import (
	"context"

	"github.com/bookoo-billy/jukebox/db"
	v1 "github.com/bookoo-billy/jukebox/proto/api/v1"
)

type SongServer struct {
	v1.UnimplementedSongServiceServer
	dao db.SongDAO
}

func NewSongServer(dao db.SongDAO) v1.SongServiceServer {
	return &SongServer{dao: dao}
}

func (s *SongServer) List(ctx context.Context, query *v1.SongQuery) (*v1.SongList, error) {
	return s.dao.List(ctx, query)
}

func (s *SongServer) Create(ctx context.Context, request *v1.SongCreateRequest) (*v1.Song, error) {
	return s.dao.Create(ctx, request.Song)
}

func (s *SongServer) Get(ctx context.Context, query *v1.SongQuery) (*v1.Song, error) {
	return s.dao.Get(ctx, query)
}

func (s *SongServer) Update(ctx context.Context, request *v1.SongUpdateRequest) (*v1.Song, error) {
	return s.dao.Update(ctx, request.Song)
}

func (s *SongServer) Delete(ctx context.Context, query *v1.SongQuery) (*v1.Song, error) {
	return s.dao.Delete(ctx, query)
}
