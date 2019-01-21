package metadata

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Get gets a metadata value from incoming context
func Get(ctx context.Context, key string) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ok
	}

	vs := md.Get(key)
	if len(vs) == 0 {
		return "", false
	}

	return vs[0], true
}

// Set appends a metadata value to the incoming context
func Set(ctx context.Context, key, value string) (context.Context, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md := metadata.New(map[string]string{
			key: value,
		})
		return metadata.NewIncomingContext(ctx, md), true
	}

	md.Append(key, value)
	return metadata.NewIncomingContext(ctx, md), true
}

// TraceIDForwarder forwards `uber-trace-id` value from incoming to outgoing context metadata
func TraceIDForwarder() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for _, v := range md.Get("uber-trace-id") {
				ctx = metadata.AppendToOutgoingContext(ctx, "uber-trace-id", v)
			}
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
