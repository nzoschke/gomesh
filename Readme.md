# gomesh

Building a reliable service oriented architecture is easier than ever, once you learn the gRPC framework and ecosystem of tools that interoperate around Protocol Buffer service definitions.

This is an example SOA with all the gRPC services configured correctly and explained in depth. See [the docs folder](docs/) for detailed guides about service definitions, proxies, remote procedure calls, API gateways, data stores, observability, versioning and more.

With this foundation you can skip over all the setup, and focus entirely on your business logic code.

## Motivation

It demonstrates:

| Component                | Via                                | Config, Code                      |
|--------------------------|------------------------------------|:---------------------------------:|
| [Service Definitions][1] | Protocol Buffers                   | [⚙️](protos/users/v1/users.proto) |
| [Services][2]            | gRPC, Go                           | [💾](cmd/users-v1/main.go)        |
| [Clients][3]             | gRPC, Go                           | [💾](cmd/users-v2/main.go)        |
| [Service Proxy][4]       | Envoy, gRPC                        | [⚙️](config/envoy/sidecar.yaml)   |
| [Observability][5]       | Envoy, gRPC middleware, Prometheus | [⚙️](configs/prometheus.yml)      |
| [API Gateway][4]         | Envoy                              | [⚙️](config/envoy/gateway.yaml)   |
| Rate Limiting            | Envoy, Redis                       | 🛠                                |
| Service Discovery        | Envoy, Consul                      | 🛠                                |
| [Fault tolerance][6]     | Envoy, gRPC middleware             | 🛠                                |
| Datastores               | Envoy, Mongo, Redis                | 🛠                                |
| REST API Gateway         | Envoy, Swagger                     | 🛠                                |
| GraphQL API Gateway      | Rejoiner                           | 🛠                                |

[1]: docs/protocol-buffers.md
[2]: docs/grpc-service.md
[3]: docs/grpc-client.md
[4]: docs/envoy-service-proxy.md
[5]: docs/observability-prometheus.md
[6]: docs/fault-tolerance.md

## Quick Start

This project uses:

- [Docker CE](https://www.docker.com/community-edition)
- [Go 1.11](https://golang.org/)
- [fullstorydev/grpcurl](https://github.com/fullstorydev/grpcurl)
- [uber/prototool](https://github.com/uber/prototool)

### Install the CLI tools:

```console
$ brew install go prototool
$ curl -L https://github.com/fullstorydev/grpcurl/releases/download/v1.1.0/grpcurl_1.1.0_osx_x86_64.tar.gz | tar -xvz -C /usr/local/bin/ grpcurl
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

We start by getting and testing the `github.com/nzoschke/gomesh`.

```shell
$ git clone https://github.com/nzoschke/gomesh.git ~/dev/gomesh
$ cd ~/dev/gomesh

$ go run cmd/server/users-v1/main.go
listening on :8000

$ grpcurl -plaintext localhost:8000 list
gomesh.users.v1.Users
grpc.reflection.v1alpha.ServerReflection

$ grpcurl -plaintext localhost:8000 describe gomesh.users.v1.Users
gomesh.users.v1.Users is a service:
service Users {
  rpc Create ( .gomesh.users.v1.CreateRequest ) returns ( .gomesh.users.v1.User );
  rpc Get ( .gomesh.users.v1.GetRequest ) returns ( .gomesh.users.v1.User );
}

$ go run cmd/client/users-v1/main.go
USER: id:"5581e658-08ea-40bc-afde-cd4864623259" parent:"orgs/myorg" name:"users/myusername" display_name:"My Full Name" create_time:<seconds:1546543774 nanos:304714000 > 
2019/01/03 11:29:34 rpc error: code = NotFound desc = users/foo not found
```

This gives us confidence in our gRPC and Go environment.

### Develop the project

We can pull in newer client/server interfaces with `go get`:

```shell
$ go get github.com/nzoschke/gomesh-interface@7b002f2c
```

We can start all the development services:

```shell
$ make dev
```

```shell
```

## Docs

Check out [the docs folder](docs/) where each component is explained in more detail.

## Contributing

Find a bug or see a way to improve the project? [Open an issue](https://github.com/nzoschke/omgrpc/issues).

## License

This work is copyright Noah Zoschke and licensed under a [Creative Commons Attribution 4.0 Unported License](https://creativecommons.org/licenses/by/4.0/).
