# gomesh

Building a reliable service oriented architecture (SOA) is easier than ever, once you learn Protocol Buffer service definitions, the gRPC service framework, the Envoy service proxy and ecosystem of related tools.

This is an example SOA with all the gRPC services and Envoy proxies configured correctly and explained in depth. See [the docs folder](docs/) for detailed guides about service definitions, proxies, remote procedure calls, API gateways, data stores, observability, versioning and more.

With this foundation you can skip over all the setup, and focus entirely on your business logic code.

## Motivation

Mesh networking with Envoy is one of the latest advances in building a SOA. The gRPC service framework is well-suited for mesh networking due to its design, performance, error-handling and first-class support in Envoy. Go is well-suited for building gRPC services due to its performance and type safety. Check out the [Intro to Go Service Mesh with Envoy and gRPC](docs/intro-gomesh.md) for more explanation.

Due to its flexibility, Envoy can be difficult to configure. This project demonstrates a solid foundation for Envoy, gRPC and Go. You can clone it and run different configurations locally with a few commands to get a feel for the architecture and its performance and observability.

Then you can copy the configs into sidecars for your production services. The configs, with tweaks, have been tested in production at [Segment](https://segment.com/).

It demonstrates:

| Component                     | With                                      | Docs, Code, Cfg |
|-------------------------------|-------------------------------------------|:---------------:|
| Service Definitions           | Protocol Buffers                          | [⚙️][2]          |
| Client and Server Interfaces  | Docker, GitHub Actions, prototool         | [⚙️][4][💾][5]   |
| Services                      | gRPC, Go                                  | [💾][7]          |
| Clients                       | gRPC, Go                                  | [💾][9]          |
| Proto Design and Conventions  | Google API Design Guide, prototool        | [📖][11]         |
| gRPC Middleware               | go-grpc-middleware, protoc-gen-validate   | [⚙️][13][💾][14] |
| Service Mesh                  | Docker, Envoy sidecar/discovery, gRPC     | [⚙️][16]         |
| Service Proxy                 | Docker, Envoy discovery/transcoding, gRPC | [⚙️][18]         |
| Observability                 | Envoy logging/stats, Prometheus discovery | [⚙️][20]         |
| Fault tolerance               | Envoy retries/circuit breaker             | 🛠               |
| REST API Gateway              | Envoy filters, ory/hydra, lyft/ratelimit  | 🛠               |
| Running on AWS                | ALB, CloudFormation, ECS                  | 🛠               |
| Datastores                    | Envoy, Mongo, Redis                       | 🛠               |

[2]: https://github.com/nzoschke/gomesh-proto/blob/master/proto/users/v1/users.proto
[4]: https://github.com/nzoschke/gomesh-proto/blob/master/proto/prototool.yaml
[5]: https://github.com/nzoschke/gomesh-proto/tree/master/.github/action/gen
[7]: cmd/server/users-v1/main.go
[9]: cmd/client/users-v1/main.go
[11]: https://cloud.google.com/apis/design/
[13]: https://github.com/nzoschke/gomesh-proto/blob/master/proto/users/v2/users.proto
[14]: cmd/server/users-v2/main.go
[16]: config/envoy/sidecar.yaml
[18]: config/envoy/proxy-xds.yaml
[19]: docs/observability-envoy-prometheus.md
[20]: config/prometheus/dns-sd.yml

## Quick Start

This project spans two repositories to isolate generated code from service definitions and implementations.

1. ([gomesh-interface](https://github.com/nzoschke/gomesh-interface)) for generated Go client and server interfaces
2. ([gomesh](https://github.com/nzoschke/gomesh)) for .proto definitions and gRPC service implementations

This project uses:

- [bojand/ghz](https://github.com/bojand/ghz)
- [Docker CE](https://www.docker.com/community-edition)
- [fullstorydev/grpcurl](https://github.com/fullstorydev/grpcurl)
- [Go 1.11](https://golang.org/)
- [rakyll/hey](https://github.com/rakyll/hey)
- [uber/prototool](https://github.com/uber/prototool)

### Install the CLI tools:

```console
$ brew install go prototool
$ curl -L https://github.com/bojand/ghz/releases/download/v0.22.0/ghz_0.22.0_Darwin_x86_64.tar.gz | tar -xvz -C /usr/local/bin/ ghz
$ curl -L https://github.com/fullstorydev/grpcurl/releases/download/v1.1.0/grpcurl_1.1.0_osx_x86_64.tar.gz | tar -xvz -C /usr/local/bin/ grpcurl
$ go install github.com/rakyll/hey
$ open https://store.docker.com/search?type=edition&offering=community
```

<details>
<summary>We may want to upgrade existing tools...</summary>
&nbsp;

```console
$ brew upgrade go prototool
```
</details>

<details>
<summary>We may want to double check the installed versions...</summary>
&nbsp;

```console
$ docker version
Client: Docker Engine - Community
 Version:           18.09.0
 API version:       1.39
 Go version:        go1.10.4
 Git commit:        4d60db4
 Built:             Wed Nov  7 00:47:43 2018
 OS/Arch:           darwin/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          18.09.0
  API version:      1.39 (minimum version 1.12)
  Go version:       go1.10.4
  Git commit:       4d60db4
  Built:            Wed Nov  7 00:55:00 2018
  OS/Arch:          linux/amd64
  Experimental:     false

$ ghz -v
0.22.0

$ go version
go version go1.11.4 darwin/amd64

$ grpcurl -version
grpcurl v1.1.0

$ prototool version
Version:                 1.3.0
Default protoc version:  3.6.1
Go version:              go1.11
Built:                   Mon Sep 17 17:46:54 UTC 2018
OS/Arch:                 darwin/amd64
```
</details>

### Get the project

We start by getting and testing `github.com/nzoschke/gomesh`.

```shell
$ git clone https://github.com/nzoschke/gomesh.git ~/dev/gomesh
$ cd ~/dev/gomesh
```

### Run Go servers

We can run a Go gRPC server and client, and use `grpcurl` to interact with it:

```shell
$ go run cmd/server/users-v1/main.go
listening on :8002

$ grpcurl -plaintext localhost:8002 list
gomesh.users.v1.Users
grpc.reflection.v1alpha.ServerReflection

$ grpcurl -plaintext localhost:8002 describe gomesh.users.v1.Users
gomesh.users.v1.Users is a service:
service Users {
  rpc Create ( .gomesh.users.v1.CreateRequest ) returns ( .gomesh.users.v1.User );
  rpc Get ( .gomesh.users.v1.GetRequest ) returns ( .gomesh.users.v1.User );
}

$ go run cmd/client/users-v1/main.go
USER: id:"5581e658-08ea-40bc-afde-cd4864623259" parent:"orgs/myorg" name:"users/myusername" display_name:"My Full Name" create_time:<seconds:1546543774 nanos:304714000 > 
2019/01/03 11:29:34 rpc error: code = NotFound desc = users/foo not found
```

We can run Go gRPC servers and REST gateway, and use `curl` to interact with it:

```shell
$ go run cmd/server/users-v2/main.go
$ go run cmd/server/widgets-v2/main.go -p 9002
$ go run cmd/server/users-v2-gateway/main.go
$ curl -s localhost:9000/v2/orgs/foo/users/bar | jq
```

```json
{
  "name": "orgs/foo/users/bar",
  "widgets": [
    {
      "parent": "orgs/foo/users/bar",
      "name": "orgs/foo/users/bar/widgets/bar",
      "display_name": "A fine widget",
      "color": "WIDGET_COLOR_BLUE"
    }
  ]
}
```

We can pull in newer client/server interfaces with `go get`:

```shell
$ go get github.com/nzoschke/gomesh-interface@master
```

This gives us confidence in our gRPC and Go environment.

### Run Envoy proxies with Docker

We can run an Envoy mesh as many containers configured with Docker Compose.

We can bulid all the Docker images with:

```shell
$ make dc-build
```

We can start a service mesh with:

```shell
$ make dc-up-mesh
$ grpcurl -d '{"name": "orgs/foo/users/bar"}' -plaintext localhost:10000 gomesh.users.v2.Users/Get
```

We can start all the API Gateway services / sidecars:

```shell
$ make dc-up-gateway
GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/server/widgets-v2 cmd/server/widgets-v2/main.go
...
Creating network "gomesh_default" with the default driver
Creating gomesh_prometheus_1 ... done
Creating gomesh_widgets-v2_1 ... done
...

$ curl localhost/v2/orgs/myorg/users/myuser
No auth header present
```

You can auth with any username:password, or you can find an OAuth token in the logs.

```
$ docker logs --since 5m gomesh_auth-v2alpha_1 2>&1 | grep TOKEN
CreateToken TOKEN=sxhmAdIuo_sVeNk0UQhkCeZeCp7u-FiFJelg40eEnF4.5vBxqCQegt9Eg-ITQsrnla_KZf1GvMrOAOS_a-qm_qg
```

```shell
$ curl -u foo:bar localhost/v2/orgs/myorg/users/myuser
$ curl -H "Authorization: Bearer $TOKEN" localhost/v2/orgs/myorg/users/myuser
```

Expected output:

```json
{
 "parent": "",
 "name": "orgs/myorg/users/myuser",
 "display_name": "",
 "widgets": [
  {
   "parent": "orgs/myorg/users/myuser",
   "name": "orgs/myorg/users/myuser/widgets/bar",
   "display_name": "A fine widget",
   "color": "WIDGET_COLOR_BLUE"
  }
 ]
}
```

We can shut everything down:

```shell
$ make dc-down
```

This gives us confidence in our Docker and Envoy configuration.

## Docs

Check out [the docs folder](docs/) where components are explained in more detail.

## Contributing

Find a bug or see a way to improve the project? [Open an issue](https://github.com/nzoschke/gomesh/issues).

## License

This work is copyright Noah Zoschke and licensed under a [Creative Commons Attribution 4.0 Unported License](https://creativecommons.org/licenses/by/4.0/).
