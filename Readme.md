# gomesh

Building a reliable service oriented architecture is easier than ever, once you learn the gRPC framework and ecosystem of tools that interoperate around Protocol Buffer service definitions.

This is an example SOA with all the gRPC services configured correctly and explained in depth. See [the docs folder](docs/) for detailed guides about service definitions, proxies, remote procedure calls, API gateways, data stores, observability, versioning and more.

With this foundation you can skip over all the setup, and focus entirely on your business logic code.

## Motivation

It demonstrates:

| Component                           | Via                                       | Config, Code    |
|-------------------------------------|-------------------------------------------|:---------------:|
| [Service Definitions][1]            | Protocol Buffers                          | [⚙️][2]          |
| [Client and Server Interfaces][3]   | Docker, GitHub Actions, prototool         | [⚙️][4][💾][5]   |
| [Services][6]                       | gRPC, Go                                  | [💾][7]          |
| [Clients][8]                        | gRPC, Go                                  | [💾][9]          |
| [Proto Design and Conventions][10]  | Google API Design Guide, prototool        | [📖][11]         |
| [gRPC Middleware][12]               | go-grpc-middleware, protoc-gen-validate   | [⚙️][13][💾][14] |
| [Service Mesh][15]                  | Docker, Envoy sidecar/discovery, gRPC     | [⚙️][16]         |
| [Service Proxy][17]                 | Docker, Envoy discovery/transcoding       | [⚙️][18]         |
| [Observability][19]                 | Envoy logging/stats, Prometheus discovery | [⚙️][20]         |
| Fault tolerance                     | Envoy retries/circuit breaker             | 🛠               |
| Datastores                          | Envoy, Mongo, Redis                       | 🛠               |
| REST API Gateway                    | Envoy filters, Hydra, lyft/ratelimit      | 🛠               |
| GraphQL API Gateway                 | Rejoiner                                  | 🛠               |

[1]: docs/protocol-buffers.md
[2]: https://github.com/nzoschke/gomesh-proto/blob/master/proto/users/v1/users.proto
[3]: docs/generating-clients-server-interfaces.md
[4]: https://github.com/nzoschke/gomesh-proto/blob/master/proto/prototool.yaml
[5]: https://github.com/nzoschke/gomesh-proto/tree/master/.github/action/gen
[6]: docs/grpc-services.md
[7]: cmd/server/users-v1/main.go
[8]: docs/grpc-clients.md
[9]: cmd/client/users-v1/main.go
[10]: docs/proto-standards.md
[11]: https://cloud.google.com/apis/design/
[12]: docs/grpc-middleware.md
[13]: https://github.com/nzoschke/gomesh-proto/blob/master/proto/users/v2/users.proto
[14]: cmd/server/users-v2/main.go
[15]: docs/envoy-service-mesh.md
[16]: config/envoy/sidecar.yaml
[17]: docs/envoy-service-proxy.md
[18]: config/envoy/proxy-xds.yaml
[19]: docs/observability-envoy-prometheus.md
[20]: config/prometheus/dns-sd.yml

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
$ go get github.com/nzoschke/gomesh-interface@branch
```

We can start all the mesh services / sidecars:

```shell
$ make dc-up-mesh
GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/server/widgets-v2 cmd/server/widgets-v2/main.go
...
Creating network "gomesh_default" with the default driver
Creating gomesh_prometheus_1 ... done
Creating gomesh_widgets-v2_1 ... done
...

$ grpcurl -d '{"name": "orgs/myorg/users/myuser"}' -plaintext localhost:10000 gomesh.users.v2.Users/Get
```

```json
{
  "name": "orgs/myorg/users/myuser",
  "widgets": [
    {
      "parent": "orgs/myorg/users/myuser",
      "name": "orgs/myorg/users/myuser/widgets/bar",
      "displayName": "A fine widget",
      "color": "WIDGET_COLOR_BLUE"
    }
  ]
}
```

We can start all the service proxy and mesh:

```shell
$ make dc-up-proxy
$ curl localhost/v2/orgs/myorg/users/myuser
```

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

## Docs

Check out [the docs folder](docs/) where each component is explained in more detail.

## Contributing

Find a bug or see a way to improve the project? [Open an issue](https://github.com/nzoschke/gomesh/issues).

## License

This work is copyright Noah Zoschke and licensed under a [Creative Commons Attribution 4.0 Unported License](https://creativecommons.org/licenses/by/4.0/).
