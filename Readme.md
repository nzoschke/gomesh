# Oh My gRPC

Building a reliable service oriented architecture is easier than ever, once you learn the gRPC framework and ecosystem of tools that interoperate around Protocol Buffer service definitions.

This is an example SOA with all the gRPC services configured correctly and explained in depth. See [the docs folder](docs/) for detailed guides about service definitions, proxies, remote procedure calls, API gateways, data stores, observability, versioning and more.

With this foundation you can skip over all the setup, and focus entirely on your business logic code.

## Motivation

It demonstrates:

| Component           | Via                    | Config, Code              |
|---------------------|------------------------|:-------------------------:|
| Service Definitions | Protocol Buffers       | 🛠                        |
| Services[2]         | gRPC, Go               | [💾](cmd/users-v1/main.go)|
| Service Discovery   | Envoy, Consul          | 🛠                        |
| RPC                 | Envoy, gRPC            | 🛠                        |
| Datastores          | Envoy, Mongo, Redis    | 🛠                        |
| Rest API Gateway    | Envoy, Swagger         | 🛠                        |
| GraphQL API Gateway | Rejoiner               | 🛠                        |
| Observability       | Envoy, gRPC middleware | 🛠                        |

[2]: docs/grpc-service.md

## Quick Start

This project uses:

- [Go 1.11](https://golang.org/)
- [grpc-go](https://github.com/grpc/grpc-go)
- [Prototool](https://github.com/uber/prototool)


Install the CLI tools:

```console
$ brew install go prototool
```

### Get the project

We start by getting and testing the `github.com/nzoschke/omgrpc`.

```shell
$ git clone https://github.com/nzoschke/omgrpc.git ~/dev/omgrpc
$ cd ~/dev/omgrpc

$ go run cmd/users-v1/main.go
listening on :8080

$ prototool grpc                    \
--address 0.0.0.0:8080              \
--method omgrpc.users.v1.Users/Get  \
--data '{"name": "foo"}'

rpc error: code = NotFound desc =  not found
```

This gives us confidence in our gRPC environment.

## Docs

Check out [the docs folder](docs/) where each component is explained in more detail.

## Contributing

Find a bug or see a way to improve the project? [Open an issue](https://github.com/nzoschke/omgrpc/issues).

## License

This work is copyright Noah Zoschke and licensed under a [Creative Commons Attribution 4.0 Unported License](https://creativecommons.org/licenses/by/4.0/).