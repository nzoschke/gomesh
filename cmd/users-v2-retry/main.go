package main

import (
	"context"
	"fmt"
	"net"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	users "github.com/nzoschke/gomesh/gen/go/users/v2"
	widgets "github.com/nzoschke/gomesh/gen/go/widgets/v1"
	"github.com/segmentio/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type config struct {
	Port        int    `conf:"p" help:"Port to listen"`
	WidgetsAddr string `conf:"w" help:"Widgets service address to dial"`
}

func main() {
	config := config{
		Port:        8000,
		WidgetsAddr: "0.0.0.0:8001",
	}
	conf.Load(&config)

	if err := serve(config); err != nil {
		panic(err)
	}
}

func serve(config config) error {
	conn, err := grpc.Dial(
		config.WidgetsAddr,
		grpc.WithAuthority("widgets-v1"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(3),
		)),
	)
	if err != nil {
		return err
	}
	defer conn.Close()
	c := widgets.NewWidgetsClient(conn)

	s := grpc.NewServer()
	users.RegisterUsersServer(s, &Server{
		WidgetsClient: c,
	})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}

// Server implements the widgets/v1 interface
type Server struct {
	WidgetsClient widgets.WidgetsClient
}

// Create creates a User
func (s *Server) Create(ctx context.Context, u *users.CreateRequest) (*users.User, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Get returns a User with their Widgets
func (s *Server) Get(ctx context.Context, u *users.GetRequest) (*users.User, error) {
	r, err := s.WidgetsClient.List(ctx, &widgets.ListRequest{
		Parent: u.Name,
	})
	if err != nil {
		return nil, err
	}

	return &users.User{
		Name:    u.Name,
		Widgets: r.Widgets,
	}, nil
}
