package main

import (
	"fmt"
	"net"

	widgets "github.com/nzoschke/gomesh-interface/gen/go/widgets/v2"
	swidgets "github.com/nzoschke/gomesh/server/widgets/v2"
	"github.com/segmentio/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port int `conf:"p" help:"Port to listen"`
}

func main() {
	config := config{
		Port: 8001,
	}
	conf.Load(&config)

	if err := serve(config); err != nil {
		panic(err)
	}
}

func serve(config config) error {
	s := grpc.NewServer()
	widgets.RegisterWidgetsServer(s, &swidgets.Server{})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}
