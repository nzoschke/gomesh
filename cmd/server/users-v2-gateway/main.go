package main

import (
	"context"
	"fmt"
	"net/http"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	users "github.com/nzoschke/gomesh-interface/gen/go/users/v2"
	"github.com/segmentio/conf"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type config struct {
	Port      int    `conf:"p" help:"Port to listen"`
	UsersAddr string `conf:"w" help:"Users service address to dial"`
}

func main() {
	config := config{
		Port:      9000,
		UsersAddr: "0.0.0.0:8002",
	}
	conf.Load(&config)

	if err := serve(config); err != nil {
		panic(err)
	}
}

func serve(config config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logEntry := log.NewEntry(logger)

	mux := runtime.NewServeMux()

	err := users.RegisterUsersHandlerFromEndpoint(ctx, mux, config.UsersAddr, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_logrus.UnaryClientInterceptor(logEntry),
		),
	})
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Port), mux)
}
