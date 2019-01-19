package main

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	widgets "github.com/nzoschke/gomesh-interface/gen/go/widgets/v2"
	swidgets "github.com/nzoschke/gomesh/server/widgets/v2"
	"github.com/segmentio/conf"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	ErrorRate int `conf:"error-rate" help:"Percent of requests to fail"`
	Port      int `conf:"p"          help:"Port to listen"`
}

func main() {
	config := config{
		ErrorRate: 0,
		Port:      8002,
	}
	conf.Load(&config)

	if err := serve(config); err != nil {
		panic(err)
	}
}

func serve(config config) error {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logEntry := log.NewEntry(logger)

	logger.Debugf("Config: %+v", config)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(logEntry),
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)

	widgets.RegisterWidgetsServer(s, &swidgets.Server{
		ErrorRate: config.ErrorRate,
	})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}
