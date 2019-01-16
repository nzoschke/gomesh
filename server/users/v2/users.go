package users

import (
	"context"
	"fmt"

	empty "github.com/golang/protobuf/ptypes/empty"
	users "github.com/nzoschke/gomesh-interface/gen/go/users/v2"
	widgets "github.com/nzoschke/gomesh-interface/gen/go/widgets/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Interface assertion
var _ users.UsersServer = (*Server)(nil)

// Server implements the users/v3 interface
type Server struct {
	WidgetsClient widgets.WidgetsClient
}

// BatchGet returns a batch of Users by name
func (s *Server) BatchGet(ctx context.Context, in *users.BatchGetRequest) (*users.BatchGetResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Create creates a User by parent and id
func (s *Server) Create(ctx context.Context, in *users.CreateRequest) (*users.User, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Delete deletes a User by name
func (s *Server) Delete(ctx context.Context, in *users.DeleteRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Get returns a User with their Widgets
func (s *Server) Get(ctx context.Context, u *users.GetRequest) (*users.User, error) {
	subj, _ := mdGet(ctx, "x-subject-id")
	fmt.Printf("users.Get x-subject-id=%s\n", subj)

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

// List returns a page of Users
func (s *Server) List(ctx context.Context, in *users.ListRequest) (*users.ListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Update updates a User
func (s *Server) Update(ctx context.Context, in *users.UpdateRequest) (*users.User, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func mdGet(ctx context.Context, key string) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ok
	}

	vs := md.Get(key)
	if len(vs) == 0 {
		return "", false
	}

	return vs[0], true
}
