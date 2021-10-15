package server

import (
	"context"

	v1 "github.com/bookoo-billy/jukebox/api/v1"
	"github.com/bookoo-billy/jukebox/db"
)

type SongServer struct {
	v1.UnimplementedSongServiceServer
	dao *db.SongDAO
}

func NewSongServer(dao *db.SongDAO) v1.SongServiceServer {
	return &SongServer{dao: dao}
}

func (s *SongServer) List(context.Context, *v1.SongQuery) (*v1.SongList, error) {
	panic("implement me")
}

func (s *SongServer) Create(context.Context, *v1.SongCreateRequest) (*v1.Song, error) {
	panic("implement me")
}

func (s *SongServer) Get(context.Context, *v1.SongQuery) (*v1.Song, error) {
	panic("implement me")
}

func (s *SongServer) Update(context.Context, *v1.SongUpdateRequest) (*v1.Song, error) {
	panic("implement me")
}

func (s *SongServer) Delete(context.Context, *v1.SongQuery) (*v1.Song, error) {
	panic("implement me")
}

func (s *SongServer) Play(context.Context, *v1.SongPlayRequest) (*v1.Song, error) {
	panic("implement me")
}

func (s *SongServer) Stop(context.Context, *v1.SongStopRequest) (*v1.Song, error) {
	panic("implement me")
}

func (s *SongServer) Receiver(context.Context, *v1.ReceiverResponse) (*v1.ReceiverCommand, error) {
	panic("implement me")
}
