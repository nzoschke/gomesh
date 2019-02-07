package auth

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/avast/retry-go"
	"github.com/nzoschke/gomesh/internal/metadata"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// HydraConfiguration is config for Hydra
type HydraConfiguration struct {
	AdminURL     string
	Authority    string
	PublicURL    string
	ClientID     string
	ClientSecret string
	Scopes       []string
}

// Transport adds host and trace headers to requests
type Transport struct {
	Authority string
	RequestID string
	TraceID   string
}

// RoundTrip implements the RoundTripper interface
// It rewrites the host for Envoy forwarding and propogates a trace ID
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("uber-trace-id", t.TraceID)
	req.Header.Add("x-request-id", t.RequestID)
	req.Host = t.Authority
	return http.DefaultTransport.RoundTrip(req)
}

// NewHydraSDK returns a Hydra SDK
func NewHydraSDK(ctx context.Context, config *HydraConfiguration) (*hydra.CodeGenSDK, error) {
	if config.AdminURL == "" {
		return nil, errors.New("Please specify the ORY Hydra Admin URL")
	}

	config.AdminURL = strings.TrimLeft(config.AdminURL, "/")
	o := swagger.NewAdminApiWithBasePath(config.AdminURL)
	sdk := &hydra.CodeGenSDK{
		AdminApi: o,
		Configuration: &hydra.Configuration{
			AdminURL:     config.AdminURL,
			PublicURL:    config.PublicURL,
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			Scopes:       config.Scopes,
		},
	}

	// custom transport for hydra SDK requests
	rid, _ := metadata.Get(ctx, "x-request-id")
	tid, _ := metadata.Get(ctx, "uber-trace-id")
	sdk.AdminApi.Configuration.Transport = &Transport{
		Authority: config.Authority,
		RequestID: rid,
		TraceID:   tid,
	}

	return sdk, nil
}

// CreateToken creates an OAuth client and client credentials token, retrying while Hydra initializes
func CreateToken(config *HydraConfiguration) (token string, err error) {
	// Inject a trace ID header to correlate Hydra errors and requests
	// FIXME: use Jager client to emit root span, addressing Jaeger UI `trace-without-root-span`
	tr := rand.Uint64()
	ctx, _ := metadata.Set(context.Background(), "uber-trace-id", fmt.Sprintf("%x:%x:0:1", tr, tr))

	retry.Do(func() error {
		token, err = createToken(ctx, config)
		if err != nil {
			fmt.Printf("CreateToken error: %+v\n", err)
		} else {
			fmt.Printf("CreateToken TOKEN=%s\n", token)
		}
		return err
	})

	return
}

// createToken creates an OAuth client and client credentials token
func createToken(ctx context.Context, config *HydraConfiguration) (string, error) {
	h, err := NewHydraSDK(ctx, config)
	if err != nil {
		return "", err
	}

	c, _, err := h.CreateOAuth2Client(swagger.OAuth2Client{
		ClientId:     config.ClientID,
		ClientSecret: config.ClientSecret,
		GrantTypes:   []string{"client_credentials"},
	})
	if err != nil {
		return "", err
	}

	// custom transport for oauth2 requests
	ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
		Transport: h.AdminApi.Configuration.Transport,
	})

	oauthConfig := clientcredentials.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		TokenURL:     fmt.Sprintf("%s/oauth2/token", config.PublicURL),
	}

	t, err := oauthConfig.Token(ctx)
	if err != nil {
		return "", err
	}

	return t.AccessToken, nil
}
