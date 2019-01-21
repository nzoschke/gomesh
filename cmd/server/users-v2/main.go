package main

import (
	"context"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	users "github.com/nzoschke/gomesh-interface/gen/go/users/v2"
	widgets "github.com/nzoschke/gomesh-interface/gen/go/widgets/v2"
	"github.com/nzoschke/gomesh/internal/metadata"
	susers "github.com/nzoschke/gomesh/server/users/v2"
	"github.com/segmentio/conf"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port             int    `conf:"p"  help:"Port to listen"`
	WidgetsHost      string `conf:"wh" help:"Widgets service host to dial"`
	WidgetsAuthority string `conf:"wa" help:"Widgets service authority header"`
}

func main() {
	config := config{
		Port:             8002,
		WidgetsHost:      "0.0.0.0:9002",
		WidgetsAuthority: "widgets-v2",
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

	conn, err := grpc.Dial(
		config.WidgetsHost,
		grpc.WithAuthority(config.WidgetsAuthority),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				metadata.TraceIDForwarder(),
				grpc_logrus.UnaryClientInterceptor(logEntry),
				grpc_logrus.PayloadUnaryClientInterceptor(
					logEntry,
					func(ctx context.Context, methodFullName string) bool {
						return true
					},
				),
			),
		),
	)
	if err != nil {
		return err
	}
	defer conn.Close()
	c := widgets.NewWidgetsClient(conn)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(logEntry),
				grpc_logrus.PayloadUnaryServerInterceptor(
					logEntry,
					func(ctx context.Context, methodFullName string, servingObject interface{}) bool {
						return true
					},
				),
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)

	users.RegisterUsersServer(s, &susers.Server{
		WidgetsClient: c,
	})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}
