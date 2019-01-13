package main

import (
	"fmt"
	"net"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2alpha"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	sauth "github.com/nzoschke/gomesh/server/auth/v2alpha"
	"github.com/segmentio/conf"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port           int    `conf:"p"  help:"Port to listen"`
	HydraHost      string `conf:"hh" help:"Hydra service host to dial"`
	HydraAuthority string `conf:"ha" help:"Hydra service authority header"`
}

func main() {
	config := config{
		Port:           8000,
		HydraHost:      "0.0.0.0:4444",
		HydraAuthority: "hydra",
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

	auth.RegisterAuthorizationServer(s, &sauth.Server{})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}
