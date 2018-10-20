FROM golang:1.11

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY cmd cmd
COPY gen gen
RUN go install ./...
