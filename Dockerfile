FROM golang:1.11

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go install ./...
