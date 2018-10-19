package main

import (
	"context"
	"fmt"
	"net"

	users "github.com/nzoschke/omgrpc/gen/go/protos/users/v2"
	widgets "github.com/nzoschke/omgrpc/gen/go/protos/widgets/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	if err := serve(); err != nil {
		panic(err)
	}
}

func serve() error {
	conn, err := grpc.Dial("0.0.0.0:8001", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := widgets.NewWidgetsClient(conn)

	s := grpc.NewServer()
	users.RegisterUsersServer(s, &Server{
		WidgetsClient: c,
	})

	l, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		return err
	}

	fmt.Println("listening on :8000!")
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
