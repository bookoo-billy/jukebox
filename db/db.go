package db

import (
	v1 "bookoo-billy/jukebox/gen/proto/v1"
	"context"
)

type JukeboxDb interface {
	Close() error
	Albums() AlbumDAO
	Artists() ArtistDAO
	Playlists() PlaylistDAO
	Receivers() ReceiverDAO
	Songs() SongDAO
}

type AlbumDAO interface {
	Create(ctx context.Context, album *v1.Album) (*v1.Album, error)
	Delete(ctx context.Context, query *v1.AlbumQuery) (*v1.Album, error)
	Get(ctx context.Context, query *v1.AlbumQuery) (*v1.Album, error)
	List(ctx context.Context, query *v1.AlbumQuery) (*v1.AlbumList, error)
	Update(ctx context.Context, album *v1.Album) (*v1.Album, error)
}

type ArtistDAO interface {
	Create(ctx context.Context, artist *v1.Artist) (*v1.Artist, error)
	Delete(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error)
	Get(ctx context.Context, query *v1.ArtistQuery) (*v1.Artist, error)
	List(ctx context.Context, query *v1.ArtistQuery) (*v1.ArtistList, error)
	Update(ctx context.Context, artist *v1.Artist) (*v1.Artist, error)
}

type PlaylistDAO interface {
	Create(ctx context.Context, playlist *v1.Playlist) (*v1.Playlist, error)
	Delete(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error)
	Get(ctx context.Context, query *v1.PlaylistQuery) (*v1.Playlist, error)
	List(ctx context.Context, query *v1.PlaylistQuery) (*v1.PlaylistList, error)
	Update(ctx context.Context, playlist *v1.Playlist) (*v1.Playlist, error)
}

type ReceiverDAO interface {
	Create(ctx context.Context, receiver *v1.Receiver) (*v1.Receiver, error)
	Delete(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error)
	Get(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error)
	List(ctx context.Context, query *v1.ReceiverQuery) (*v1.ReceiverList, error)
	Update(ctx context.Context, receiver *v1.Receiver) (*v1.Receiver, error)
}

type SongDAO interface {
	Create(ctx context.Context, song *v1.Song) (*v1.Song, error)
	Delete(ctx context.Context, query *v1.SongQuery) (*v1.Song, error)
	Get(ctx context.Context, query *v1.SongQuery) (*v1.Song, error)
	List(ctx context.Context, query *v1.SongQuery) (*v1.SongList, error)
	Update(ctx context.Context, song *v1.Song) (*v1.Song, error)
}
