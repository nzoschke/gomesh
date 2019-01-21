FROM golang:1.10.4 AS build
WORKDIR /go/src/github.com/lyft/ratelimit
ARG RATELIMIT_VERSION=459ed655

RUN git clone https://github.com/lyft/ratelimit.git /go/src/github.com/lyft/ratelimit && git reset --hard $RATELIMIT_VERSION

RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN script/install-glide
RUN glide install

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/ratelimit -ldflags="-w -s" -v github.com/lyft/ratelimit/src/service_cmd

FROM alpine:3.8 AS final
RUN apk --no-cache add ca-certificates
COPY --from=build /usr/local/bin/ratelimit /bin/ratelimit

RUN mkdir -p /srv/runtime_data/current/ratelimit
ENTRYPOINT /bin/ratelimit