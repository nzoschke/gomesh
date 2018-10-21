package main

import (
	"context"
	"fmt"
	"net"

	widgets "github.com/nzoschke/omgrpc/gen/go/protos/widgets/v1"
	"github.com/segmentio/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port int `conf:"p" help:"Port to listen"`
}

func main() {
	config := config{
		Port: 8000,
	}
	conf.Load(&config)

	if err := serve(config); err != nil {
		panic(err)
	}
}

func serve(config config) error {
	s := grpc.NewServer()
	widgets.RegisterWidgetsServer(s, &Server{})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
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
