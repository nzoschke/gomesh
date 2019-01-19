package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/avast/retry-go"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2alpha"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	sauth "github.com/nzoschke/gomesh/server/auth/v2alpha"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"github.com/segmentio/conf"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/clientcredentials"
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

	t := &sauth.Transport{
		Authority: config.HydraAuthority,
		Transport: http.DefaultTransport.(*http.Transport),
	}
	http.DefaultTransport = t

	h, err := hydra.NewSDK(&hydra.Configuration{
		AdminURL:     fmt.Sprintf("http://%s", config.HydraHostAdmin),
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		PublicURL:    fmt.Sprintf("http://%s", config.HydraHostPublic),
	})
	if err != nil {
		return err
	}

	// For demo purposes try to create a client id and token
	retry.Do(func() error {
		token, err := createDemoToken(config)
		if err != nil {
			logger.Errorf("createDemoToken error: %+v", err)
		} else {
			logger.Infof("createDemoToken TOKEN=%s", token)
		}
		return err
	})

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(logEntry),
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)

	auth.RegisterAuthorizationServer(s, &sauth.Server{
		Hydra:     h,
		Transport: t,
	})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}

func createDemoToken(config config) (string, error) {
	h, err := hydra.NewSDK(&hydra.Configuration{
		AdminURL: fmt.Sprintf("http://%s", config.HydraHostAdmin),
	})

	c, _, err := h.CreateOAuth2Client(swagger.OAuth2Client{
		ClientId:     config.ClientID,
		ClientSecret: config.ClientSecret,
		GrantTypes:   []string{"client_credentials"},
	})
	if err != nil {
		return "", err
	}

	oauthConfig := clientcredentials.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		TokenURL:     fmt.Sprintf("http://%s/oauth2/token", config.HydraHostPublic),
	}
	t, err := oauthConfig.Token(context.Background())
	if err != nil {
		return "", err
	}

	return t.AccessToken, nil
}
