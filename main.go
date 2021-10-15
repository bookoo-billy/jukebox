package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	v1 "github.com/bookoo-billy/jukebox/api/v1"
	"github.com/bookoo-billy/jukebox/db"
	"github.com/bookoo-billy/jukebox/server"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9834", "gRPC server endpoint")
	httpServerEndpoint = flag.String("http-server-endpoint", "localhost:9835", "HTTP server endpoint")
	mongoDBURI         = flag.String("mongo-db-uri", "mongodb://localhost:27017/jukebox", "mongo DB URI")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	var g errgroup.Group

	glog.Info("Starting gRPC + HTTP servers...")
	g.Go(func() error {
		return Run(ctx, "tcp", *grpcServerEndpoint)
	})

	g.Go(func() error {
		return RunInProcessGateway(ctx, *httpServerEndpoint)
	})

	g.Wait()
}

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()

	jDB := db.NewJukeboxDB(*mongoDBURI)
	defer jDB.Close()

	v1.RegisterAlbumServiceServer(s, server.NewAlbumServer(jDB.Albums()))
	v1.RegisterArtistServiceServer(s, server.NewArtistServer(jDB.Artists()))
	v1.RegisterPlaylistServiceServer(s, server.NewPlaylistServer(jDB.Playlists()))
	v1.RegisterSongServiceServer(s, server.NewSongServer(jDB.Songs()))

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
		glog.Info("Shutting down gRPC server...")
	}()

	glog.Infof("gRPC server started on %s", address)
	return s.Serve(l)
}

// RunInProcessGateway starts the invoke in process http gateway.
func RunInProcessGateway(ctx context.Context, addr string) error {
	opts := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			EmitUnpopulated: false,
		},
	})
	mux := runtime.NewServeMux(opts)
	grpcOpts := []grpc.DialOption{grpc.WithInsecure()}

	mustRegisterHandler(v1.RegisterAlbumServiceHandlerFromEndpoint, ctx, mux, *grpcServerEndpoint, grpcOpts)
	mustRegisterHandler(v1.RegisterArtistServiceHandlerFromEndpoint, ctx, mux, *grpcServerEndpoint, grpcOpts)
	mustRegisterHandler(v1.RegisterPlaylistServiceHandlerFromEndpoint, ctx, mux, *grpcServerEndpoint, grpcOpts)
	mustRegisterHandler(v1.RegisterSongServiceHandlerFromEndpoint, ctx, mux, *grpcServerEndpoint, grpcOpts)

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	glog.Infof("http server started on %s", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}

func mustRegisterHandler(register func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error, ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	err := register(ctx, mux, endpoint, opts)
	if err != nil {
		panic("Unable to register service handler from endpoint: " + err.Error())
	}
}
