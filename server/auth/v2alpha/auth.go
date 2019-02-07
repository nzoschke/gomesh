package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2alpha"
	"github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/gogo/protobuf/types"
)

// Interface assertion
var _ auth.AuthorizationServer = (*Server)(nil)

// Server implements the auth/v2alpha interface
type Server struct {
	HydraConfig *HydraConfiguration
}

// Check does an auth header check and returns a 200 on auth, non-200 otherwise
func (s *Server) Check(ctx context.Context, in *auth.CheckRequest) (*auth.CheckResponse, error) {
	h := in.Attributes.Request.Http.Headers["authorization"]
	if h == "" {
		return responseDenied("No auth header present", int(envoy_type.StatusCode_Unauthorized))
	}

	auth := strings.SplitN(h, " ", 2)
	if len(auth) != 2 {
		return responseDenied("Invalid auth header", int(envoy_type.StatusCode_Unauthorized))
	}

	switch auth[0] {
	case "Basic":
		return s.basicCheck(auth[1])
	case "Bearer":
		return s.bearerCheck(ctx, auth[1])
	default:
		return responseDenied("Invalid auth header", int(envoy_type.StatusCode_Unauthorized))
	}
}

func (s *Server) basicCheck(token string) (*auth.CheckResponse, error) {
	payload, _ := base64.StdEncoding.DecodeString(token)
	parts := strings.SplitN(string(payload), ":", 2)

	if len(parts) != 2 || !basicValidate(parts[0], parts[1]) {
		return responseDenied("Invalid basic credentials", int(envoy_type.StatusCode_Unauthorized))
	}

	return responseOk(fmt.Sprintf("users/%s", parts[0]))
}

// TODO: implement real password checking
func basicValidate(username, password string) bool {
	return true
}

func (s *Server) bearerCheck(ctx context.Context, token string) (*auth.CheckResponse, error) {
	h, err := NewHydraSDK(ctx, s.HydraConfig)
	if err != nil {
		return responseDenied(err.Error(), int(envoy_type.StatusCode_Unauthorized))
	}

	i, _, err := h.IntrospectOAuth2Token(token, "")
	if err != nil {
		return responseDenied(err.Error(), int(envoy_type.StatusCode_Unauthorized))
	}

	if !i.Active {
		return responseDenied("Inactive bearer token", int(envoy_type.StatusCode_Unauthorized))
	}

	return responseOk(fmt.Sprintf("clients/%s", i.ClientId))
}

func responseOk(subject string) (*auth.CheckResponse, error) {
	return &auth.CheckResponse{
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: &auth.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					&core.HeaderValueOption{
						Append: &types.BoolValue{
							Value: false,
						},
						Header: &core.HeaderValue{
							Key:   "x-subject-id",
							Value: subject,
						},
					},
				},
			},
		},
		Status: &rpc.Status{
			Code: int32(rpc.OK),
		},
	}, nil
}

func responseDenied(body string, status int) (*auth.CheckResponse, error) {
	return &auth.CheckResponse{
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Body: body,
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
			},
		},
		Status: &rpc.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
	}, nil
}
