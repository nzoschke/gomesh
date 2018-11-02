package users_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	users "github.com/nzoschke/gomesh/gen/go/users/v3"
	widgets "github.com/nzoschke/gomesh/gen/go/widgets/v2"
	usersServer "github.com/nzoschke/gomesh/server/users/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func server(t *testing.T, srv *usersServer.Server) (net.Listener, *grpc.Server, *grpc.ClientConn) {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)
	users.RegisterUsersServer(s, srv)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 0))
	assert.NoError(t, err)

	c, err := grpc.Dial(l.Addr().String(), grpc.WithInsecure())
	assert.NoError(t, err)

	return l, s, c
}
func TestGet(t *testing.T) {
	mw := widgets.MockWidgetsClient{}
	srv := &usersServer.Server{
		WidgetsClient: &mw,
	}

	l, s, conn := server(t, srv)
	go s.Serve(l)
	defer s.GracefulStop()
	c := users.NewUsersClient(conn)

	mw.On(
		"List",
		mock.Anything,
		&widgets.ListRequest{
			Parent: "users/foo",
		},
	).Return(
		&widgets.ListResponse{
			Widgets: []*widgets.Widget{
				&widgets.Widget{
					Parent:      "users/foo",
					Name:        "users/foo/widgets/red",
					DisplayName: "Red",
				},
			},
		},
		nil,
	)

	_, err := c.Get(context.Background(), &users.GetRequest{
		Name: "invalid",
	})
	assert.Error(t, err)
	serr := status.Convert(err)
	assert.Equal(t, codes.InvalidArgument, serr.Code())
	assert.Equal(t, `invalid GetRequest.Name: value does not match regex pattern "^users/[a-z0-9._-]+$"`, serr.Message())

	u, err := c.Get(context.Background(), &users.GetRequest{
		Name: "users/foo",
	})
	assert.NoError(t, err)
	assert.EqualValues(t, &users.User{
		Name: "users/foo",
		Widgets: []*widgets.Widget{
			&widgets.Widget{
				Parent:      "users/foo",
				Name:        "users/foo/widgets/red",
				DisplayName: "Red",
			},
		},
	}, u)
}
