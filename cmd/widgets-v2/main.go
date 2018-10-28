package main

import (
	"context"
	"fmt"
	"net"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	widgets "github.com/nzoschke/gomesh/gen/go/proto/widgets/v2"
	"github.com/segmentio/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type config struct {
	Port      int    `conf:"p" help:"Port to listen"`
	MongoAddr string `conf:"m" help:"Mongo address"`
}

func main() {
	config := config{
		Port:      8000,
		MongoAddr: "mongo:29781",
	}
	conf.Load(&config)

	if err := serve(config); err != nil {
		panic(err)
	}
}

func serve(config config) error {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)
	widgets.RegisterWidgetsServer(s, &Server{})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}

// Server implements the widgets/v2 interface
type Server struct {
	// WidgetsClient widgets.WidgetsClient
}

// BatchGet Widgets by names
func (s *Server) BatchGet(ctx context.Context, r *widgets.BatchGetRequest) (*widgets.BatchGetResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Create Widget
func (s *Server) Create(ctx context.Context, r *widgets.CreateRequest) (*widgets.Widget, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Delete Widget
func (s *Server) Delete(ctx context.Context, r *widgets.DeleteRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Get Widgets
func (s *Server) Get(ctx context.Context, r *widgets.GetRequest) (*widgets.Widget, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// List Widgets
func (s *Server) List(ctx context.Context, r *widgets.ListRequest) (*widgets.ListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Update Widget
func (s *Server) Update(ctx context.Context, r *widgets.UpdateRequest) (*widgets.Widget, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
