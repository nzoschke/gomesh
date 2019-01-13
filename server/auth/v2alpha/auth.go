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
}

// Check does an auth header check and returns a 200 on auth, non-200 otherwise
func (s *Server) Check(ctx context.Context, in *auth.CheckRequest) (*auth.CheckResponse, error) {
	h := in.Attributes.Request.Http.Headers["authorization"]
	if h == "" {
		return deniedResponse("No auth header present", int(envoy_type.StatusCode_Unauthorized))
	}

	auth := strings.SplitN(h, " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return deniedResponse("Invalid auth header", int(envoy_type.StatusCode_Unauthorized))
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	parts := strings.SplitN(string(payload), ":", 2)

	if len(parts) != 2 || !validate(parts[0], parts[1]) {
		return deniedResponse("Invalid basic credentials", int(envoy_type.StatusCode_Unauthorized))
	}

	return okResponse(fmt.Sprintf("user/%s", parts[0]))
}

func okResponse(subject string) (*auth.CheckResponse, error) {
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

func deniedResponse(body string, status int) (*auth.CheckResponse, error) {
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

func validate(username, password string) bool {
	return true
}
