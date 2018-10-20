# Go gRPC Clients

An obvious application of gRPC is to fetch data from a remote network service in Go. To accomplish this, we can use the generated Go client interface and the `google.golang.org/grpc` package to connect to a remote gRPC service. The promise of gRPC is that we can reliably get data from a network server.

## Defining Related Services in .proto Files

First we create `.proto` files for two services -- the service we will call directly, and the service it calls out to for more data. 
For this example we will make a new version of our User service that includes a related `Widgets` collection, and a new `Widgets` service with a `List` method that that returns widget data.

```proto
syntax = "proto3";

package omgrpc.users.v2;

option go_package = "v2pb";

import "google/protobuf/timestamp.proto";
import "protos/widgets/v1/widgets.proto";

service Users {
  rpc Get(GetRequest) returns (User);
  rpc Create(CreateRequest) returns (User);
}

message User {
  ...

  repeated omgrpc.widgets.v1.Widget widgets = 6;
}

message GetRequest {
  string name = 1;
}
```
> From protos/users/v2/users.proto

```proto
syntax = "proto3";

package omgrpc.widgets.v1;

option go_package = "v1pb";

import "google/protobuf/timestamp.proto";

service Widgets {
  rpc List(ListRequest) returns (ListResponse);
}

message Widget {
  string id = 1;
  string parent = 2;
  string name = 3;
  string display_name = 4;
  google.protobuf.Timestamp create_time = 5;
}

message ListRequest {
  string parent = 1;
  int32 page_size = 2;
  string page_token = 3;
}

message ListResponse {
  repeated Widget widgets = 1;
  string next_page_token = 2;
}
```
> From protos/widgets/v1/widgets.proto

## Generating Client and Server Interfaces

Next we create a `.go` files with a gRPC client interface. The `prototool generate` command invokes the `protoc` compiler to create this for us.

```shell
$ prototool generate
```

```go
package v1pb

type Widget struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string               `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

type ListRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

type ListResponse struct {
	Widgets              []*Widget `protobuf:"bytes,1,rep,name=widgets,proto3" json:"widgets,omitempty"`
	NextPageToken        string    `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

// WidgetsClient is the client API for Widgets service.
type WidgetsClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}
```
> From gen/go/protos/widgets/v1/widgets.pb.go

## Implementing the Server

We also have to implement the Widgets server. For this example we can return static data:

```go
// Server implements the widgets/v1 interface
type Server struct{}

// List lists widgets
func (s *Server) List(ctx context.Context, r *widgets.ListRequest) (*widgets.ListResponse, error) {
	return &widgets.ListResponse{
		Widgets: []*widgets.Widget{
			&widgets.Widget{
				Name: "widgets/red",
			},
			&widgets.Widget{
				Name: "widgets/blue",
			},
		},
	}, nil
}
```
> from cmd/widgets-v1/main.go

## Using the Client

Now we can use the client. We import the generated Go package, and use the `google.golang.org/grpc` to dial the Widgets service.

```go
package main

import (
	"context"
	"fmt"
	"net"

	users "github.com/nzoschke/omgrpc/gen/go/protos/users/v2"
	widgets "github.com/nzoschke/omgrpc/gen/go/protos/widgets/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func serve() error {
	conn, err := grpc.Dial("0.0.0.0:8001", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := widgets.NewWidgetsClient(conn)

	s := grpc.NewServer()
	users.RegisterUsersServer(s, &Server{
		WidgetsClient: c,
	})

    ...
}

// Server implements the widgets/v1 interface
type Server struct {
	WidgetsClient widgets.WidgetsClient
}

// Get returns a User with their Widgets
func (s *Server) Get(ctx context.Context, u *users.GetRequest) (*users.User, error) {
	r, err := s.WidgetsClient.List(ctx, &widgets.ListRequest{
		Parent: u.Name,
	})
	if err != nil {
		return nil, err
	}

	return &users.User{
		Name:    u.Name,
		Widgets: r.Widgets,
	}, nil
}
```

## Testing

Finally we can test our service. We run the server programs, then use the `prototool grpc` command to translate JSON ⟷ gRPC requests and responses:

```shell
$ go run cmd/users-v2/main.go
listening on :8000

$ go run cmd/widgets-v1/main.go -p 8001
listening on :8001

$ prototool grpc                      \
--address 0.0.0.0:8000                \
--method omgrpc.users.v2.Users/Get    \
--data '{
    "name": "users/fred"
  }'
```

```json
{
  "name": "users/fred",
  "widgets": [
    {
      "name": "widgets/red"
    },
    {
      "name": "widgets/blue"
    }
  ]
}
```

## Summary

Using gRPC clients in Go is easy. We just have to:

- Define relations in .proto files
- Generate a Go client interfaces
- Make gRPC requests over the network

Go, gRPC, and `prototool` make building network services easy.