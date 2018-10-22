FROM golang:1.11

WORKDIR /app

COPY bin/linux_amd64/* /go/bin/
