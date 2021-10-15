// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AlbumServiceClient is the client API for AlbumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlbumServiceClient interface {
	List(ctx context.Context, in *AlbumQuery, opts ...grpc.CallOption) (*AlbumList, error)
	Create(ctx context.Context, in *AlbumCreateRequest, opts ...grpc.CallOption) (*Album, error)
	Get(ctx context.Context, in *AlbumQuery, opts ...grpc.CallOption) (*Album, error)
	Update(ctx context.Context, in *AlbumUpdateRequest, opts ...grpc.CallOption) (*Album, error)
	Delete(ctx context.Context, in *AlbumQuery, opts ...grpc.CallOption) (*Album, error)
}

type albumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlbumServiceClient(cc grpc.ClientConnInterface) AlbumServiceClient {
	return &albumServiceClient{cc}
}

func (c *albumServiceClient) List(ctx context.Context, in *AlbumQuery, opts ...grpc.CallOption) (*AlbumList, error) {
	out := new(AlbumList)
	err := c.cc.Invoke(ctx, "/jukebox.AlbumService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albumServiceClient) Create(ctx context.Context, in *AlbumCreateRequest, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/jukebox.AlbumService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albumServiceClient) Get(ctx context.Context, in *AlbumQuery, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/jukebox.AlbumService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albumServiceClient) Update(ctx context.Context, in *AlbumUpdateRequest, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/jukebox.AlbumService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albumServiceClient) Delete(ctx context.Context, in *AlbumQuery, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/jukebox.AlbumService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlbumServiceServer is the server API for AlbumService service.
// All implementations must embed UnimplementedAlbumServiceServer
// for forward compatibility
type AlbumServiceServer interface {
	List(context.Context, *AlbumQuery) (*AlbumList, error)
	Create(context.Context, *AlbumCreateRequest) (*Album, error)
	Get(context.Context, *AlbumQuery) (*Album, error)
	Update(context.Context, *AlbumUpdateRequest) (*Album, error)
	Delete(context.Context, *AlbumQuery) (*Album, error)
	mustEmbedUnimplementedAlbumServiceServer()
}

// UnimplementedAlbumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAlbumServiceServer struct {
}

func (UnimplementedAlbumServiceServer) List(context.Context, *AlbumQuery) (*AlbumList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAlbumServiceServer) Create(context.Context, *AlbumCreateRequest) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAlbumServiceServer) Get(context.Context, *AlbumQuery) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAlbumServiceServer) Update(context.Context, *AlbumUpdateRequest) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAlbumServiceServer) Delete(context.Context, *AlbumQuery) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAlbumServiceServer) mustEmbedUnimplementedAlbumServiceServer() {}

// UnsafeAlbumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlbumServiceServer will
// result in compilation errors.
type UnsafeAlbumServiceServer interface {
	mustEmbedUnimplementedAlbumServiceServer()
}

func RegisterAlbumServiceServer(s grpc.ServiceRegistrar, srv AlbumServiceServer) {
	s.RegisterService(&AlbumService_ServiceDesc, srv)
}

func _AlbumService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.AlbumService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).List(ctx, req.(*AlbumQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlbumService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.AlbumService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).Create(ctx, req.(*AlbumCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlbumService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.AlbumService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).Get(ctx, req.(*AlbumQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlbumService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.AlbumService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).Update(ctx, req.(*AlbumUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlbumService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.AlbumService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).Delete(ctx, req.(*AlbumQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// AlbumService_ServiceDesc is the grpc.ServiceDesc for AlbumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlbumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jukebox.AlbumService",
	HandlerType: (*AlbumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AlbumService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AlbumService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AlbumService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AlbumService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AlbumService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jukebox.proto",
}

// ArtistServiceClient is the client API for ArtistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtistServiceClient interface {
	List(ctx context.Context, in *ArtistQuery, opts ...grpc.CallOption) (*ArtistList, error)
	Create(ctx context.Context, in *ArtistCreateRequest, opts ...grpc.CallOption) (*Artist, error)
	Get(ctx context.Context, in *ArtistQuery, opts ...grpc.CallOption) (*Artist, error)
	Update(ctx context.Context, in *ArtistUpdateRequest, opts ...grpc.CallOption) (*Artist, error)
	Delete(ctx context.Context, in *ArtistQuery, opts ...grpc.CallOption) (*Artist, error)
}

type artistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtistServiceClient(cc grpc.ClientConnInterface) ArtistServiceClient {
	return &artistServiceClient{cc}
}

func (c *artistServiceClient) List(ctx context.Context, in *ArtistQuery, opts ...grpc.CallOption) (*ArtistList, error) {
	out := new(ArtistList)
	err := c.cc.Invoke(ctx, "/jukebox.ArtistService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) Create(ctx context.Context, in *ArtistCreateRequest, opts ...grpc.CallOption) (*Artist, error) {
	out := new(Artist)
	err := c.cc.Invoke(ctx, "/jukebox.ArtistService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) Get(ctx context.Context, in *ArtistQuery, opts ...grpc.CallOption) (*Artist, error) {
	out := new(Artist)
	err := c.cc.Invoke(ctx, "/jukebox.ArtistService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) Update(ctx context.Context, in *ArtistUpdateRequest, opts ...grpc.CallOption) (*Artist, error) {
	out := new(Artist)
	err := c.cc.Invoke(ctx, "/jukebox.ArtistService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) Delete(ctx context.Context, in *ArtistQuery, opts ...grpc.CallOption) (*Artist, error) {
	out := new(Artist)
	err := c.cc.Invoke(ctx, "/jukebox.ArtistService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtistServiceServer is the server API for ArtistService service.
// All implementations must embed UnimplementedArtistServiceServer
// for forward compatibility
type ArtistServiceServer interface {
	List(context.Context, *ArtistQuery) (*ArtistList, error)
	Create(context.Context, *ArtistCreateRequest) (*Artist, error)
	Get(context.Context, *ArtistQuery) (*Artist, error)
	Update(context.Context, *ArtistUpdateRequest) (*Artist, error)
	Delete(context.Context, *ArtistQuery) (*Artist, error)
	mustEmbedUnimplementedArtistServiceServer()
}

// UnimplementedArtistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArtistServiceServer struct {
}

func (UnimplementedArtistServiceServer) List(context.Context, *ArtistQuery) (*ArtistList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedArtistServiceServer) Create(context.Context, *ArtistCreateRequest) (*Artist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArtistServiceServer) Get(context.Context, *ArtistQuery) (*Artist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedArtistServiceServer) Update(context.Context, *ArtistUpdateRequest) (*Artist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArtistServiceServer) Delete(context.Context, *ArtistQuery) (*Artist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArtistServiceServer) mustEmbedUnimplementedArtistServiceServer() {}

// UnsafeArtistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtistServiceServer will
// result in compilation errors.
type UnsafeArtistServiceServer interface {
	mustEmbedUnimplementedArtistServiceServer()
}

func RegisterArtistServiceServer(s grpc.ServiceRegistrar, srv ArtistServiceServer) {
	s.RegisterService(&ArtistService_ServiceDesc, srv)
}

func _ArtistService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.ArtistService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).List(ctx, req.(*ArtistQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.ArtistService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).Create(ctx, req.(*ArtistCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.ArtistService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).Get(ctx, req.(*ArtistQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.ArtistService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).Update(ctx, req.(*ArtistUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.ArtistService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).Delete(ctx, req.(*ArtistQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtistService_ServiceDesc is the grpc.ServiceDesc for ArtistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jukebox.ArtistService",
	HandlerType: (*ArtistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ArtistService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ArtistService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ArtistService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ArtistService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ArtistService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jukebox.proto",
}

// PlaylistServiceClient is the client API for PlaylistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaylistServiceClient interface {
	List(ctx context.Context, in *PlaylistQuery, opts ...grpc.CallOption) (*PlaylistList, error)
	Create(ctx context.Context, in *PlaylistCreateRequest, opts ...grpc.CallOption) (*Playlist, error)
	Get(ctx context.Context, in *PlaylistQuery, opts ...grpc.CallOption) (*Playlist, error)
	Update(ctx context.Context, in *PlaylistUpdateRequest, opts ...grpc.CallOption) (*Playlist, error)
	Delete(ctx context.Context, in *PlaylistQuery, opts ...grpc.CallOption) (*Playlist, error)
}

type playlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistServiceClient(cc grpc.ClientConnInterface) PlaylistServiceClient {
	return &playlistServiceClient{cc}
}

func (c *playlistServiceClient) List(ctx context.Context, in *PlaylistQuery, opts ...grpc.CallOption) (*PlaylistList, error) {
	out := new(PlaylistList)
	err := c.cc.Invoke(ctx, "/jukebox.PlaylistService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) Create(ctx context.Context, in *PlaylistCreateRequest, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/jukebox.PlaylistService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) Get(ctx context.Context, in *PlaylistQuery, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/jukebox.PlaylistService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) Update(ctx context.Context, in *PlaylistUpdateRequest, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/jukebox.PlaylistService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) Delete(ctx context.Context, in *PlaylistQuery, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/jukebox.PlaylistService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServiceServer is the server API for PlaylistService service.
// All implementations must embed UnimplementedPlaylistServiceServer
// for forward compatibility
type PlaylistServiceServer interface {
	List(context.Context, *PlaylistQuery) (*PlaylistList, error)
	Create(context.Context, *PlaylistCreateRequest) (*Playlist, error)
	Get(context.Context, *PlaylistQuery) (*Playlist, error)
	Update(context.Context, *PlaylistUpdateRequest) (*Playlist, error)
	Delete(context.Context, *PlaylistQuery) (*Playlist, error)
	mustEmbedUnimplementedPlaylistServiceServer()
}

// UnimplementedPlaylistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlaylistServiceServer struct {
}

func (UnimplementedPlaylistServiceServer) List(context.Context, *PlaylistQuery) (*PlaylistList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPlaylistServiceServer) Create(context.Context, *PlaylistCreateRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPlaylistServiceServer) Get(context.Context, *PlaylistQuery) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPlaylistServiceServer) Update(context.Context, *PlaylistUpdateRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPlaylistServiceServer) Delete(context.Context, *PlaylistQuery) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPlaylistServiceServer) mustEmbedUnimplementedPlaylistServiceServer() {}

// UnsafePlaylistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaylistServiceServer will
// result in compilation errors.
type UnsafePlaylistServiceServer interface {
	mustEmbedUnimplementedPlaylistServiceServer()
}

func RegisterPlaylistServiceServer(s grpc.ServiceRegistrar, srv PlaylistServiceServer) {
	s.RegisterService(&PlaylistService_ServiceDesc, srv)
}

func _PlaylistService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.PlaylistService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).List(ctx, req.(*PlaylistQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.PlaylistService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).Create(ctx, req.(*PlaylistCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.PlaylistService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).Get(ctx, req.(*PlaylistQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.PlaylistService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).Update(ctx, req.(*PlaylistUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.PlaylistService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).Delete(ctx, req.(*PlaylistQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaylistService_ServiceDesc is the grpc.ServiceDesc for PlaylistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaylistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jukebox.PlaylistService",
	HandlerType: (*PlaylistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PlaylistService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PlaylistService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PlaylistService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PlaylistService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PlaylistService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jukebox.proto",
}

// SongServiceClient is the client API for SongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SongServiceClient interface {
	List(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (*SongList, error)
	Create(ctx context.Context, in *SongCreateRequest, opts ...grpc.CallOption) (*Song, error)
	Get(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (*Song, error)
	Update(ctx context.Context, in *SongUpdateRequest, opts ...grpc.CallOption) (*Song, error)
	Delete(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (*Song, error)
	Play(ctx context.Context, in *SongPlayRequest, opts ...grpc.CallOption) (*Song, error)
	Stop(ctx context.Context, in *SongStopRequest, opts ...grpc.CallOption) (*Song, error)
	Receiver(ctx context.Context, in *ReceiverResponse, opts ...grpc.CallOption) (*ReceiverCommand, error)
}

type songServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSongServiceClient(cc grpc.ClientConnInterface) SongServiceClient {
	return &songServiceClient{cc}
}

func (c *songServiceClient) List(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (*SongList, error) {
	out := new(SongList)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Create(ctx context.Context, in *SongCreateRequest, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Get(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Update(ctx context.Context, in *SongUpdateRequest, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Delete(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Play(ctx context.Context, in *SongPlayRequest, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Stop(ctx context.Context, in *SongStopRequest, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) Receiver(ctx context.Context, in *ReceiverResponse, opts ...grpc.CallOption) (*ReceiverCommand, error) {
	out := new(ReceiverCommand)
	err := c.cc.Invoke(ctx, "/jukebox.SongService/Receiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SongServiceServer is the server API for SongService service.
// All implementations must embed UnimplementedSongServiceServer
// for forward compatibility
type SongServiceServer interface {
	List(context.Context, *SongQuery) (*SongList, error)
	Create(context.Context, *SongCreateRequest) (*Song, error)
	Get(context.Context, *SongQuery) (*Song, error)
	Update(context.Context, *SongUpdateRequest) (*Song, error)
	Delete(context.Context, *SongQuery) (*Song, error)
	Play(context.Context, *SongPlayRequest) (*Song, error)
	Stop(context.Context, *SongStopRequest) (*Song, error)
	Receiver(context.Context, *ReceiverResponse) (*ReceiverCommand, error)
	mustEmbedUnimplementedSongServiceServer()
}

// UnimplementedSongServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSongServiceServer struct {
}

func (UnimplementedSongServiceServer) List(context.Context, *SongQuery) (*SongList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSongServiceServer) Create(context.Context, *SongCreateRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSongServiceServer) Get(context.Context, *SongQuery) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSongServiceServer) Update(context.Context, *SongUpdateRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSongServiceServer) Delete(context.Context, *SongQuery) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSongServiceServer) Play(context.Context, *SongPlayRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedSongServiceServer) Stop(context.Context, *SongStopRequest) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedSongServiceServer) Receiver(context.Context, *ReceiverResponse) (*ReceiverCommand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receiver not implemented")
}
func (UnimplementedSongServiceServer) mustEmbedUnimplementedSongServiceServer() {}

// UnsafeSongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SongServiceServer will
// result in compilation errors.
type UnsafeSongServiceServer interface {
	mustEmbedUnimplementedSongServiceServer()
}

func RegisterSongServiceServer(s grpc.ServiceRegistrar, srv SongServiceServer) {
	s.RegisterService(&SongService_ServiceDesc, srv)
}

func _SongService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).List(ctx, req.(*SongQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Create(ctx, req.(*SongCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Get(ctx, req.(*SongQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Update(ctx, req.(*SongUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Delete(ctx, req.(*SongQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongPlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Play(ctx, req.(*SongPlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Stop(ctx, req.(*SongStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_Receiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiverResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).Receiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jukebox.SongService/Receiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).Receiver(ctx, req.(*ReceiverResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// SongService_ServiceDesc is the grpc.ServiceDesc for SongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jukebox.SongService",
	HandlerType: (*SongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _SongService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SongService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SongService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SongService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SongService_Delete_Handler,
		},
		{
			MethodName: "Play",
			Handler:    _SongService_Play_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _SongService_Stop_Handler,
		},
		{
			MethodName: "Receiver",
			Handler:    _SongService_Receiver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jukebox.proto",
}
