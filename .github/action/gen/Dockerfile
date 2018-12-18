FROM golang:1.11

ARG PROTOC_VERSION=3.6.1
ARG PROTOTOOL_VERSION=1.3.0

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

RUN \
  curl -sSL https://github.com/uber/prototool/releases/download/v$PROTOTOOL_VERSION/prototool-Linux-x86_64 -o /bin/prototool && \
  chmod +x /bin/prototool

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