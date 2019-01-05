# Protocol Buffer Service Definition File
## With prototool

```shell
$ prototool create proto/users/v2/users.proto
```

```proto
syntax = "proto3";

package gomesh.users.v2;

option go_package = "v2pb";
option java_multiple_files = true;
option java_outer_classname = "UsersProto";
option java_package = "com.gomesh.users.v2";
```

```proto
syntax = "proto3";

package gomesh.users.v1;

option go_package = "v1pb";
option java_multiple_files = true;
option java_outer_classname = "UsersProto";
option java_package = "com.gomesh.users.v1";

import "google/protobuf/timestamp.proto";

service Users {
  rpc Get(GetRequest) returns (User);
  rpc Create(CreateRequest) returns (User);
}

message User {
  string id = 1;
  string parent = 2;
  string name = 3;
  string display_name = 4;
  google.protobuf.Timestamp create_time = 5;
}

message GetRequest {
  string name = 1;
}

message CreateRequest {
  string parent = 1;
  string user_id = 2;
  User user = 3;
}
```
> From gomesh-proto/proto/users/v1/users.proto

```shell
$ prototool format -w proto
$ prototool lint proto
```

# Generating Client / Server Interfaces
## With prototool, GitHub Actions

```Dockerfile
FROM golang:1.11

ARG PROTOC_VERSION=3.6.1
ARG PROTOTOOL_VERSION=dev

ARG MOCKERY_VERSION=ea265755d541b124de6bc248f7744eab9005fd33
ARG PROTOC_GEN_GO_VERSION=1.2.0
ARG PROTOC_GEN_SWAGGER_VERSION=1.5.1
ARG PROTOC_GEN_VALIDATE_VERSION=0.0.10
ARG TS_PROTOC_GEN_VERSION=0.8.0

RUN \
  curl -sL https://deb.nodesource.com/setup_10.x | bash - && \
  apt-get update && \
  apt-get install -y curl git nodejs && \
  rm -rf /var/lib/apt/lists/*

RUN GO111MODULE=off go get -u github.com/myitcv/gobin
RUN gobin github.com/uber/prototool/cmd/prototool@$PROTOTOOL_VERSION

RUN \
  mkdir /tmp/prototool-bootstrap && \
  echo 'protoc:\n  version:' $PROTOC_VERSION > /tmp/prototool-bootstrap/prototool.yaml && \
  echo 'syntax = "proto3";' > /tmp/prototool-bootstrap/tmp.proto && \
  prototool compile /tmp/prototool-bootstrap && \
  rm -rf /tmp/prototool-bootstrap

RUN go get github.com/vektra/mockery/... && \
  cd /go/src/github.com/vektra/mockery && \
  git checkout $MOCKERY_VERSION && \
  go install ./cmd/mockery

RUN go get github.com/golang/protobuf/... && \
  cd /go/src/github.com/golang/protobuf && \
  git checkout v$PROTOC_GEN_GO_VERSION && \
  go install ./protoc-gen-go

RUN go get github.com/lyft/protoc-gen-validate && \
  cd /go/src/github.com/lyft/protoc-gen-validate && \
  git checkout v$PROTOC_GEN_VALIDATE_VERSION && \
  go install .

RUN go get github.com/grpc-ecosystem/grpc-gateway/... && \
  cd /go/src/github.com/grpc-ecosystem/grpc-gateway && \
  git checkout v$PROTOC_GEN_SWAGGER_VERSION && \
  go install ./protoc-gen-swagger

RUN npm install -g ts-protoc-gen@$TS_PROTOC_GEN_VERSION

WORKDIR /github/workspace
ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
```

```bash
#!/bin/bash
set -ex

prototool lint      proto
prototool format -l proto
prototool generate  proto
prototool generate  proto_ext

find gen -name 'mock_*.go' -delete
mockery -all -dir gen -inpkg
```

```yaml
---
create:
  packages:
    - directory: .
      name: gomesh

generate:
  go_options:
    extra_modifiers:
      google/api/annotations.proto: google.golang.org/genproto/googleapis/api/annotations
      google/api/http.proto: google.golang.org/genproto/googleapis/api/annotations
    import_path: github.com/nzoschke/gomesh-proto/proto

  plugins:
    - file_suffix: pb
      include_imports: true
      include_source_info: true
      name: descriptor_set
      output: ../gen/pb

    - flags: plugins=grpc
      name: go
      output: ../gen/go
      type: go

    - flags: binary,import_style=commonjs
      name: js
      output: ../gen/js

    - name: swagger
      output: ../gen/swagger
      type: go

    - flags: service=true
      name: ts
      output: ../gen/js

    - flags: lang=go
      name: validate
      output: ../gen/go
      type: go

lint:
  rules:
    add:
      - RPCS_HAVE_COMMENTS
    remove:
      - REQUEST_RESPONSE_TYPES_IN_SAME_FILE
      - REQUEST_RESPONSE_TYPES_UNIQUE

protoc:
  includes:
    - ../proto_ext
    - ../proto_ext/third_party/googleapis
  version: 3.6.1
```
> from proto/prototool.yaml

```shell
$ make gen
```

Don't want to check in gen/ so its .gitignored.

Do want to automatically generate artifacts somewhere so we use a GitHub action.


```c
workflow "publish generated clients and servers" {
  on = "push"
  resolves = ["push-gen"]
}

action "gen" {
  uses = "./.github/action/gen"
}

action "push-gen" {
  needs = ["gen"]
  uses = "./.github/action/gen"
  runs = ".github/push-gen.sh"
  secrets = ["GITHUB_TOKEN", "PUSH_TOKEN"]
}
```

# Building gRPC Clients / Servers
## With Go and go-grpc

```go
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/satori/go.uuid"

	"github.com/golang/protobuf/ptypes"
	users "github.com/nzoschke/gomesh-interface/gen/go/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	if err := serve(); err != nil {
		panic(err)
	}
}

func serve() error {
	s := grpc.NewServer()
	users.RegisterUsersServer(s, &Server{})
	reflection.Register(s)

	l, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		return err
	}

	fmt.Println("listening on :8000")
	return s.Serve(l)
}

// Server implements the users/v1 interface
type Server struct{}

// Create creates a User
func (s *Server) Create(ctx context.Context, in *users.CreateRequest) (*users.User, error) {
	return &users.User{
		CreateTime:  ptypes.TimestampNow(),
		DisplayName: in.User.DisplayName,
		Id:          uuid.NewV4().String(),
		Name:        fmt.Sprintf("users/%s", in.UserId),
		Parent:      in.Parent,
	}, nil
}

// Get returns a User or NotFound error
func (s *Server) Get(ctx context.Context, u *users.GetRequest) (*users.User, error) {
	return nil, status.Errorf(codes.NotFound, "%s not found", u.Name)
}
```
> From gomesh/cmd/server/users-v1/main.go

```go
package main

import (
	"context"
	"fmt"
	"log"

	users "github.com/nzoschke/gomesh-interface/gen/go/users/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	c := users.NewUsersClient(conn)

	ctx := context.Background()
	u, err := c.Create(ctx, &users.CreateRequest{
		Parent: "orgs/myorg",
		UserId: "myusername",
		User: &users.User{
			DisplayName: "My Full Name",
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("USER: %+v\n", u)

	u, err = c.Get(ctx, &users.GetRequest{
		Name: "users/foo",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("USER: %+v\n", u)
}
```
> From gomesh/cmd/client/users-v1/main.go

# Proto Standards
## With Google API Design Guide, lyft/protoc-gen-validate

```
$ prototool create proto/users/v2/users.proto
$ prototool create proto/widgets/v2/widgets.proto
```

```proto
syntax = "proto3";

package gomesh.users.v2;

option go_package = "v2pb";
option java_multiple_files = true;
option java_outer_classname = "UsersProto";
option java_package = "com.gomesh.users.v2";

import "google/api/annotations.proto";
import "google/protobuf/empty.proto";
import "google/protobuf/field_mask.proto";
import "google/protobuf/timestamp.proto";
import "validate/validate.proto";
import "widgets/v2/widgets.proto";

service Users {
  // Get User
  //
  // Takes User name (with parent Org) in path
  rpc Get(GetRequest) returns (User) {
    option (google.api.http) = {
      get: "/v2/{name=orgs/*/users/*}"
    };
  }
  // Create User
  //
  // Takes parent Org in path and User in body
  rpc Create(CreateRequest) returns (User) {
    option (google.api.http) = {
      post: "/v2/{parent=orgs/*}/users"
      body: "*"
    };
  }
  // Update User
  //
  // Takes User name (with parent Org) in path, and User in body
  rpc Update(UpdateRequest) returns (User) {
    option (google.api.http) = {
      patch: "/v2/{user.name=orgs/*/users/*}"
      body: "*"
    };
  }
  // Delete User
  //
  // Takes User name (with parent Org) in path
  rpc Delete(DeleteRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=orgs/*/users/*}"
    };
  }
  // List Users
  //
  // Takes parent Org in path
  rpc List(ListRequest) returns (ListResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=orgs/*}/users"
    };
  }
  // BatchGet Users
  //
  // Takes parent Org in path and User names in query string
  rpc BatchGet(BatchGetRequest) returns (BatchGetResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=orgs/*}/users:batchGet"
    };
  }
}

message User {
  string parent = 1;
  string name = 2; // Output only
  string display_name = 3;
  google.protobuf.Timestamp create_time = 4; // Output only
  repeated gomesh.widgets.v2.Widget widgets = 5;
}

message GetRequest {
  // Name
  //
  // Format "orgs/{org_id}/users/{user_id}"
  string name = 1 [
    (validate.rules).string = {
      pattern: "^orgs/[a-z0-9._-]+/users/[a-z0-9._-]+$"
      max_bytes: 512
    }
  ];
}

message CreateRequest {
  // Parent
  //
  // Format "orgs/{org_id}"
  string parent = 1 [
    (validate.rules).string = {
      pattern: "^orgs/[a-z0-9._-]+$"
      max_bytes: 512
    }
  ];
  // User ID
  string user_id = 2 [
    (validate.rules).string = {
      pattern: "^[a-z0-9._-]+$"
      max_bytes: 512
    }
  ];
  // User
  //
  // Required
  User user = 3 [(validate.rules).message.required = true];
}

message UpdateRequest {
  // User
  //
  // Required
  User user = 1 [(validate.rules).message.required = true];
  // Update Mask
  //
  // Required
  google.protobuf.FieldMask update_mask = 2 [(validate.rules).message.required = true];
}

message DeleteRequest {
  string name = 1;
}

message ListRequest {
  string parent = 1;
  int32 page_size = 2;
  string page_token = 3;
}

message ListResponse {
  repeated User users = 1;
  string next_page_token = 2;
}

message BatchGetRequest {
  string parent = 1;
  repeated string names = 2;
}

message BatchGetResponse {
  repeated User users = 1;
}
```

```go
package main

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	users "github.com/nzoschke/gomesh-interface/gen/go/users/v2"
	widgets "github.com/nzoschke/gomesh-interface/gen/go/widgets/v2"
	susers "github.com/nzoschke/gomesh/server/users/v2"
	"github.com/segmentio/conf"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port        int    `conf:"p" help:"Port to listen"`
	WidgetsAddr string `conf:"w" help:"Widgets service address to dial"`
}

func main() {
	config := config{
		Port:        8000,
		WidgetsAddr: "0.0.0.0:8001",
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

	conn, err := grpc.Dial(
		config.WidgetsAddr,
		grpc.WithAuthority("widgets-v2"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_logrus.UnaryClientInterceptor(logEntry),
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
```

```shell
$ go run cmd/server/widgets-v2/main.go
listening on :8001

$ go run cmd/server/users-v2/main.go
listening on :8000

$ grpcurl -d '{"name": "foo"}' -plaintext localhost:8000 gomesh.users.v2.Users/Get
```

Example output:

```
ERROR:
  Code: InvalidArgument
  Message: invalid GetRequest.Name: value does not match regex pattern "^orgs/[a-z0-9._-]+/users/[a-z0-9._-]+$"
```

```shell
$ grpcurl -d '{"name": "orgs/myorg/users/myusername"}' -plaintext localhost:8000 gomesh.users.v2.Users/Get
```

Example output:

```json
{
  "name": "orgs/myorg/users/myusername",
  "widgets": [
    {
      "parent": "orgs/myorg/users/myusername",
      "name": "orgs/myorg/users/myusername/widgets/bar",
      "displayName": "A fine widget",
      "color": "WIDGET_COLOR_BLUE"
    }
  ]
}
```

Example logs:

```
INFO[0018] finished unary call with code InvalidArgument  error="rpc error: code = InvalidArgument desc = invalid GetRequest.Name: value does not match regex pattern \"^orgs/[a-z0-9._-]+/users/[a-z0-9._-]+$\"" grpc.code=InvalidArgument grpc.method=Get grpc.service=gomesh.users.v2.Users grpc.start_time="2019-01-04T14:51:51-08:00" grpc.time_ms=0.103 span.kind=server system=grpc
DEBU[0034] finished client unary call                    grpc.code=OK grpc.method=List grpc.service=gomesh.widgets.v2.Widgets grpc.time_ms=0.943 span.kind=client system=grpc
INFO[0034] finished unary call with code OK              grpc.code=OK grpc.method=Get grpc.service=gomesh.users.v2.Users grpc.start_time="2019-01-04T14:52:08-08:00" grpc.time_ms=1.053 span.kind=server system=grpc
```

```shell
$ curl -s localhost:9000/v2/orgs/myorg/users/myusername | jq
```

```json
{
  "name": "orgs/myorg/users/myusername",
  "widgets": [
    {
      "parent": "orgs/myorg/users/myusername",
      "name": "orgs/myorg/users/myusername/widgets/bar",
      "display_name": "A fine widget",
      "color": "WIDGET_COLOR_BLUE"
    }
  ]
}
```