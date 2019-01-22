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
	users "github.com/nzoschke/gomesh-interface/gen/go/users/v2"
	widgets "github.com/nzoschke/gomesh-interface/gen/go/widgets/v2"
	susers "github.com/nzoschke/gomesh/server/users/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func server(t *testing.T, srv *susers.Server) (net.Listener, *grpc.Server, *grpc.ClientConn) {
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

func TestGetWithWidgets(t *testing.T) {
	mw := widgets.MockWidgetsClient{}
	srv := &susers.Server{
		WidgetsClient: &mw,
	}

	l, s, conn := server(t, srv)
	go s.Serve(l)
	defer s.GracefulStop()
	c := users.NewUsersClient(conn)

	p := "orgs/foo/users/bar"
	mw.On(
		"List",
		mock.Anything,
		&widgets.ListRequest{
			Parent: p,
		},
	).Return(
		&widgets.ListResponse{
			Widgets: []*widgets.Widget{
				&widgets.Widget{
					Parent:      p,
					Name:        p + "/widgets/red",
					DisplayName: "Red",
				},
			},
		},
		nil,
	)

	u, err := c.Get(context.Background(), &users.GetRequest{
		Name: "orgs/foo/users/bar",
	})
	assert.NoError(t, err)
	assert.EqualValues(t, &users.User{
		Name: "orgs/foo/users/bar",
		Widgets: []*widgets.Widget{
			&widgets.Widget{
				Parent:      "orgs/foo/users/bar",
				Name:        "orgs/foo/users/bar/widgets/red",
				DisplayName: "Red",
			},
		},
	}, u)
}

func TestValidate(t *testing.T) {
	mw := widgets.MockWidgetsClient{}
	srv := &susers.Server{
		WidgetsClient: &mw,
	}

	l, s, conn := server(t, srv)
	go s.Serve(l)
	defer s.GracefulStop()
	c := users.NewUsersClient(conn)

	_, err := c.Get(context.Background(), &users.GetRequest{
		Name: "invalid",
	})
	assert.Error(t, err)
	serr := status.Convert(err)
	assert.Equal(t, codes.InvalidArgument, serr.Code())
	assert.Equal(t, `invalid GetRequest.Name: value does not match regex pattern "^orgs/[a-z0-9._-]+/users/[a-z0-9._-]+$"`, serr.Message())
}
