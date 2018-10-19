package main

import (
	"context"
	"fmt"
	"net"

	widgets "github.com/nzoschke/omgrpc/gen/go/protos/widgets/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := serve(); err != nil {
		panic(err)
	}
}

func serve() error {
	s := grpc.NewServer()
	widgets.RegisterWidgetsServer(s, &Server{})

	l, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		return err
	}

	fmt.Println("listening on :8001!")
	return s.Serve(l)
}

// Server implements the widgets/v1 interface
type Server struct{}

// List lists widgets
func (s *Server) List(ctx context.Context, r *widgets.ListRequest) (*widgets.ListResponse, error) {
	return &widgets.ListResponse{
		Widgets: []*widgets.Widget{
			&widgets.Widget{
				Name: "widgets/red",
			},
			&widgets.Widget{
				Name: "widgets/blue",
			},
		},
	}, nil
}
