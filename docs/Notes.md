# V1

Proto
- [x] users/v1/users.proto
- [x] proto -> interface workflow

gRPC
- [x] cmd/server/users-v1/main.go
- [x] standard error codes
- [x] cmd/client/users-v1/main.go
- [x] server/users/v1/users.go

Proto Standards
- [x] users/v2/users.proto
- [x] api design guide, standard methods
- [x] make lint
- [x] widgets/v2/widgets.proto
- [x] annotations

gRPC Go Middleware / Transcoders
- [x] logging middleware
- [x] validation middleware
- [x] grpc-gateway server

Envoy Sidecar
- [ ] make dev
- [ ] docker-compose.yml
- [ ] config/envoy/sidecar.yaml

Envoy API Gateway
- [ ] make dev
- [ ] docker-compose.yml ??
- [ ] config/envoy/proxy.yaml
- [ ] service discovery

--

Stretch
- [ ] generating docs
- [ ] rate limiting
- [ ] fault tolerance
- [ ] database ?
- [ ] testing w/ mocks

Overall design

- [ ] name / parent / id / slug
- [ ] make db-up
- [ ] make servers-up