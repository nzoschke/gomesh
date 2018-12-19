package users

import (
	"context"

	users "github.com/nzoschke/gomesh-proto/gen/go/users/v3"
	widgets "github.com/nzoschke/gomesh-proto/gen/go/widgets/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Interface assertion
var _ users.UsersServer = (*Server)(nil)

// Server implements the users/v3 interface
type Server struct {
	WidgetsClient widgets.WidgetsClient
}

func handleError(err error) error {
	return status.Error(codes.Unimplemented, "unimplemented")
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
