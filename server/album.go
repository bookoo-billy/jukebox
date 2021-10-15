package server

import (
	"context"

	"github.com/bookoo-billy/jukebox/db"
	v1 "github.com/bookoo-billy/jukebox/gen/api/v1"
)

type AlbumServer struct {
	v1.UnimplementedAlbumServiceServer
	dao *db.AlbumDAO
}

func NewAlbumServer(dao *db.AlbumDAO) v1.AlbumServiceServer {
	return &AlbumServer{dao: dao}
}

func (p *AlbumServer) List(ctx context.Context, query *v1.AlbumQuery) (*v1.AlbumList, error) {
	return p.dao.List(ctx, query)
}

func (p *AlbumServer) Create(ctx context.Context, create *v1.AlbumCreateRequest) (*v1.Album, error) {
	return p.dao.Create(ctx, create.Album)
}

func (p *AlbumServer) Get(ctx context.Context, query *v1.AlbumQuery) (*v1.Album, error) {
	return p.dao.Get(ctx, query)
}

func (p *AlbumServer) Update(ctx context.Context, update *v1.AlbumUpdateRequest) (*v1.Album, error) {
	return p.dao.Update(ctx, update.Album)
}

func (p *AlbumServer) Delete(ctx context.Context, album *v1.AlbumQuery) (*v1.Album, error) {
	return p.dao.Delete(ctx, album)
}
