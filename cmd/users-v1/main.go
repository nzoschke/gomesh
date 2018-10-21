package main

import (
	"context"
	"fmt"
	"net"

	"github.com/satori/go.uuid"

	"github.com/golang/protobuf/ptypes"
	users "github.com/nzoschke/omgrpc/gen/go/protos/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	if err := serve(); err != nil {
		panic(err)
	}
}

func serve() error {
	s := grpc.NewServer()
	users.RegisterUsersServer(s, &Server{})
	reflection.Register(s)

	l, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		return err
	}

	fmt.Println("listening on :8000!")
	return s.Serve(l)
}

// Server implements the users/v1 interface
type Server struct{}

// Create creates a User
func (s *Server) Create(ctx context.Context, u *users.CreateRequest) (*users.User, error) {
	return &users.User{
		CreateTime:  ptypes.TimestampNow(),
		DisplayName: u.User.DisplayName,
		Id:          uuid.NewV4().String(),
		Name:        u.User.Name,
		Parent:      u.Parent,
	}, nil
}

// Get returns a User or NotFound error
func (s *Server) Get(ctx context.Context, u *users.GetRequest) (*users.User, error) {
	return nil, status.Errorf(codes.NotFound, "%s not found", u.Name)
}
