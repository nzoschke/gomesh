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
	ClientID        string `conf:"ci" help:"Hydra introspection client id"`
	ClientSecret    string `conf:"cs" help:"Hydra introspection client secret"`
	Port            int    `conf:"p"  help:"Port to listen"`
	HydraAuthority  string `conf:"ha" help:"Hydra service authority header"`
	HydraHostAdmin  string `conf:"hha" help:"Hydra service admin host to dial"`
	HydraHostPublic string `conf:"hhp" help:"Hydra service public host to dial"`
}

func main() {
	config := config{
		ClientID:        "my-client",
		ClientSecret:    "secret",
		Port:            8002,
		HydraHostAdmin:  "0.0.0.0:4445",
		HydraHostPublic: "0.0.0.0:4444",
		HydraAuthority:  "hydra",
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

	hc := &sauth.HydraConfiguration{
		AdminURL:     fmt.Sprintf("http://%s", config.HydraHostAdmin),
		Authority:    config.HydraAuthority,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		PublicURL:    fmt.Sprintf("http://%s", config.HydraHostPublic),
	}

	// For demo purposes automatically create a client and token
	go func() {
		sauth.CreateToken(hc)
	}()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(logEntry),
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)

	auth.RegisterAuthorizationServer(s, &sauth.Server{
		HydraConfig: hc,
	})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}
