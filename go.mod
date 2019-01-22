module github.com/nzoschke/gomesh

require (
	github.com/avast/retry-go v2.1.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/envoyproxy/go-control-plane v0.6.6
	github.com/gogo/googleapis v1.1.0
	github.com/gogo/protobuf v1.2.0
	github.com/golang/protobuf v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/grpc-gateway v1.6.4
	github.com/lyft/protoc-gen-validate v0.0.12 // indirect
	github.com/nzoschke/gomesh-interface v0.0.0-20190105205516-47a6d4f9ad23
	github.com/ory/hydra v0.0.0-20181218121201-bdb6634e3d87
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/segmentio/conf v1.0.0
	github.com/segmentio/go-snakecase v1.1.0 // indirect
	github.com/segmentio/objconv v1.0.1 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/net v0.0.0-20181029044818-c44066c5c816
	golang.org/x/oauth2 v0.0.0-20181003184128-c57b0facaced
	google.golang.org/grpc v1.18.0
	gopkg.in/go-playground/mold.v2 v2.2.0 // indirect
	gopkg.in/validator.v2 v2.0.0-20180514200540-135c24b11c19 // indirect
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190110200230-915654e7eabc
	golang.org/x/text => github.com/golang/text v0.3.0
)
